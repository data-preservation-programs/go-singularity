package scan

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack/push"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/gotidy/ptr"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

var logger = log.Logger("scan")

// Scan traverses a source attachment, which represents a remote storage location,
// and processes its files. During this process, it identifies new files in the storage
// that are not already in the database and adds them. Existing files that have
// not yet been packaged into jobs are grouped into jobs of a specified size.
//
// The scan starts from the last scanned path, ensuring incremental scanning, and updates
// the `last_scanned_path` field of the SourceAttachment after each file to allow
// resuming interrupted scans.
//
// Parameters:
//   - ctx: Context for timeout and cancellation.
//   - db: A pointer to a gorm.DB object, providing database access.
//   - attachment: The source attachment that points to the storage location to scan.
//
// Returns:
//   - An error if there are issues during the scan or database operations, otherwise nil.
//
// Note: This function assumes the use of a global logger object for logging and leverages
//
//	the storagesystem package to interact with the storage system of the SourceAttachment.
//	The push package is used for processing files and managing file ranges.
func Scan(ctx context.Context, db *gorm.DB, attachment model.SourceAttachment) error {
	db = db.WithContext(ctx)
	directoryCache := make(map[string]model.DirectoryID)
	var remaining = push.NewFileRangeSet()
	var remainingFileRanges []model.FileRange
	err := db.Joins("File").
		Where("attachment_id = ? AND file_ranges.job_id is null", attachment.ID).
		Order("file_ranges.id asc").
		Find(&remainingFileRanges).Error
	if err != nil {
		return errors.WithStack(err)
	}
	logger.With("remaining", len(remainingFileRanges)).Info("remaining file ranges")
	err = addFileRangesAndCreatePackJob(ctx, db, attachment.ID, remaining, attachment.Preparation.MaxSize, remainingFileRanges...)
	if err != nil {
		return errors.WithStack(err)
	}

	sourceScanner, err := storagesystem.NewRCloneHandler(ctx, *attachment.Storage)
	if err != nil {
		return errors.WithStack(err)
	}
	entryChan := sourceScanner.Scan(ctx, "", attachment.LastScannedPath)
	var lastScannedPath *string
	defer func() {
		if lastScannedPath != nil {
			err = database.DoRetry(ctx, func() error {
				return db.Model(&model.SourceAttachment{}).Where("id = ?", attachment.ID).
					Update("last_scanned_path", lastScannedPath).Error
			})
			if err != nil {
				logger.Errorw("failed to update last scanned path", "error", err)
			}
		}
	}()
	for entry := range entryChan {
		if entry.Error != nil {
			logger.Errorw("failed to scan", "error", entry.Error)
			continue
		}

		file, fileRanges, err := push.PushFile(ctx, db, entry.Info, attachment, directoryCache)
		if err != nil {
			return errors.Wrapf(err, "failed to push file %s", entry.Info.Remote())
		}
		if file == nil {
			logger.Infow("file already exists", "path", entry.Info.Remote())
			continue
		}

		lastScannedPath = &file.Path

		err = addFileRangesAndCreatePackJob(ctx, db, attachment.ID, remaining, attachment.Preparation.MaxSize, fileRanges...)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	lastScannedPath = ptr.Of("")

	if len(remaining.FileRanges()) > 0 {
		err = createPackJob(ctx, db, attachment.ID, remaining)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func createPackJob(
	ctx context.Context,
	db *gorm.DB,
	attachmentID model.SourceAttachmentID,
	remaining *push.FileRangeSet,
) error {
	job, err := push.CreatePackJob(ctx, db, attachmentID, remaining.FileRangeIDs())
	if err != nil {
		return errors.WithStack(err)
	}
	logger.Infof("created pack job %d with %d file ranges", job.ID, len(remaining.FileRanges()))
	remaining.Reset()
	return nil
}

func addFileRangesAndCreatePackJob(
	ctx context.Context,
	db *gorm.DB,
	attachmentID model.SourceAttachmentID,
	remaining *push.FileRangeSet,
	maxSize int64,
	fileRanges ...model.FileRange) error {
	for _, fileRange := range fileRanges {
		fit := remaining.AddIfFits(fileRange, maxSize)
		if fit {
			continue
		}
		err := createPackJob(ctx, db, attachmentID, remaining)
		if err != nil {
			return errors.WithStack(err)
		}
		remaining.Add(fileRange)
	}

	return nil
}
