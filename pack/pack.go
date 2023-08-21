package pack

import (
	"context"
	"io"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/pack/daggen"
	"github.com/data-preservation-programs/singularity/pack/encryption"
	util3 "github.com/data-preservation-programs/singularity/pack/util"
	"github.com/data-preservation-programs/singularity/storagesystem"
	util2 "github.com/data-preservation-programs/singularity/util"
	"github.com/google/uuid"
	"github.com/ipfs/boxo/util"
	"github.com/rclone/rclone/fs"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	commcid "github.com/filecoin-project/go-fil-commcid"
	commp "github.com/filecoin-project/go-fil-commp-hashhash"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-log/v2"
	"github.com/multiformats/go-varint"
)

type Result struct {
	FileRangeCIDs map[uint64]cid.Cid
	Objects       map[uint64]fs.Object
	CarResults    []CarResult
}

type CarResult struct {
	StorageID   uint32
	StoragePath string
	CarFileSize int64
	PieceCID    cid.Cid
	PieceSize   int64
	RootCID     cid.Cid
	Header      []byte
	CarBlocks   []model.CarBlock
}

type nopCloser struct {
}

func (n nopCloser) Close() error {
	return nil
}

type WriteCloser struct {
	io.Writer
	io.Closer
}

var EmptyFileCid = cid.NewCidV1(cid.Raw, util.Hash([]byte("")))

var EmptyFileVarint = varint.ToUvarint(uint64(len(EmptyFileCid.Bytes())))

var logger = log.Logger("pack")

// GetCommp calculates the data commitment (CommP) and the piece size based on the
// provided commp.Calc instance and target piece size. It ensures that the
// calculated piece size matches the target piece size specified. If necessary,
// it pads the data to meet the target piece size.
//
// Parameters:
//   - calc: A pointer to a commp.Calc instance, which has been used to write data
//     and will be used to calculate the piece commitment for that data.
//   - targetPieceSize: The desired size of the piece, specified in bytes.
//
// Returns:
//   - cid.Cid: A CID (Content Identifier) representing the data commitment (CommP).
//   - uint64: The size of the piece, in bytes, after potential padding.
//   - error: An error indicating issues during the piece commitment calculation,
//     padding, or CID conversion, or nil if the operation was successful.
func GetCommp(calc *commp.Calc, targetPieceSize uint64) (cid.Cid, uint64, error) {
	rawCommp, rawPieceSize, err := calc.Digest()
	if err != nil {
		return cid.Undef, 0, errors.WithStack(err)
	}

	if rawPieceSize < targetPieceSize {
		rawCommp, err = commp.PadCommP(rawCommp, rawPieceSize, targetPieceSize)
		if err != nil {
			return cid.Undef, 0, errors.WithStack(err)
		}

		rawPieceSize = targetPieceSize
	} else if rawPieceSize > targetPieceSize {
		logger.Warn("piece size is larger than the target piece size")
	}

	commCid, err := commcid.DataCommitmentV1ToCID(rawCommp)
	if err != nil {
		return cid.Undef, 0, errors.WithStack(err)
	}

	return commCid, rawPieceSize, nil
}

func Pack(
	ctx context.Context,
	db *gorm.DB,
	job model.Job,
) ([]model.Car, error) {
	db = db.WithContext(ctx)
	pieceSize := job.Attachment.Preparation.PieceSize
	// storageWriter can be nil for inline preparation
	storageID, storageWriter, err := storagesystem.GetRandomOutputWriter(ctx, job.Attachment.Preparation.OutputStorages)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	storageReader, err := storagesystem.NewRCloneHandler(ctx, *job.Attachment.Storage)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get storage handler for %s", job.Attachment.Storage.Name)
	}

	var encryptor encryption.Encryptor
	if job.Attachment.Preparation.UseEncryption() {
		encryptor, err = encryption.NewAgeEncryptor(job.Attachment.Preparation.EncryptionRecipients)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	assembler := NewAssembler(ctx, storageReader, encryptor, job.FileRanges, job.Attachment.Preparation.MaxSize)
	defer assembler.Close()
	var cars []model.Car
	var numCarBlocks []int
	for assembler.Next() {
		var filename string
		calc := &commp.Calc{}
		var pieceCid cid.Cid
		var finalPieceSize uint64
		var fileSize int64
		if storageWriter != nil {
			reader := io.TeeReader(assembler, calc)
			filename = uuid.NewString() + ".car"
			obj, err := storageWriter.Write(ctx, filename, reader)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			fileSize = obj.Size()

			pieceCid, finalPieceSize, err = GetCommp(calc, uint64(pieceSize))
			if err != nil {
				return nil, errors.WithStack(err)
			}
			_, err = storageWriter.Move(ctx, obj, pieceCid.String()+".car")
			if err != nil && !errors.Is(err, storagesystem.ErrMoveNotSupported) {
				logger.Errorf("failed to move car file from %s to %s: %s", filename, pieceCid.String()+".car", err)
			}
			if err == nil {
				filename = pieceCid.String() + ".car"
			}
		} else {
			_, err = io.Copy(calc, assembler)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			pieceCid, finalPieceSize, err = GetCommp(calc, uint64(pieceSize))
			if err != nil {
				return nil, errors.WithStack(err)
			}
		}
		cars = append(cars, model.Car{
			PieceCID:     model.CID(pieceCid),
			PieceSize:    int64(finalPieceSize),
			RootCID:      model.CID(EmptyFileCid),
			FileSize:     fileSize,
			StorageID:    storageID,
			StoragePath:  filename,
			AttachmentID: job.AttachmentID,
		})
		numCarBlocks = append(numCarBlocks, len(assembler.carBlocks))
	}

	// Update all FileRange and file CID that are not split
	splitFileIDs := make(map[uint64]model.File)
	var updatedFiles []model.File
	splitFileBlks := make(map[uint64][]blocks.Block)
	for _, fileRange := range job.FileRanges {
		err = database.DoRetry(ctx, func() error {
			return db.Model(&model.FileRange{}).Where("id = ?", fileRange.ID).
				Update("cid", fileRange.CID).Error
		})
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if fileRange.Offset == 0 && fileRange.Length == fileRange.File.Size {
			err = database.DoRetry(ctx, func() error {
				return db.Model(&model.File{}).Where("id = ?", fileRange.FileID).
					Update("cid", fileRange.CID).Error
			})
			if err != nil {
				return nil, errors.WithStack(err)
			}
			fileRange.File.CID = fileRange.CID
			updatedFiles = append(updatedFiles, *fileRange.File)
		} else {
			splitFileIDs[fileRange.FileID] = *fileRange.File
		}
	}

	// Now check if all file ranges of a file are ready. If so, update file CID
	for fileID := range splitFileIDs {
		err = database.DoRetry(ctx, func() error {
			return db.Transaction(func(db *gorm.DB) error {
				var allParts []model.FileRange
				err = db.Where("file_id = ?", fileID).Order(clause.OrderByColumn{Column: clause.Column{Name: "carOffset"}}).Find(&allParts).Error
				if err != nil {
					return errors.WithStack(err)
				}
				if underscore.All(allParts, func(p model.FileRange) bool {
					return p.CID != model.CID(cid.Undef)
				}) {
					links := underscore.Map(allParts, func(p model.FileRange) format.Link {
						return format.Link{
							Size: uint64(p.Length),
							Cid:  cid.Cid(p.CID),
						}
					})
					blks, node, err := util3.AssembleFileFromLinks(links)
					if err != nil {
						return errors.Wrap(err, "failed to assemble file from links")
					}
					nodeCid := node.Cid()
					err = db.Model(&model.File{}).Where("id = ?", fileID).Update("cid", model.CID(nodeCid)).Error
					if err != nil {
						return errors.Wrap(err, "failed to update cid of file")
					}
					file := splitFileIDs[fileID]
					file.CID = model.CID(nodeCid)
					updatedFiles = append(updatedFiles, file)
					splitFileBlks[fileID] = blks
				}
				return nil
			})
		})
	}

	err = database.DoRetry(ctx, func() error {
		return db.Transaction(
			func(db *gorm.DB) error {
				j := 0
				for i, n := range numCarBlocks {
					car := cars[i]
					err := db.Create(&car).Error
					if err != nil {
						return errors.WithStack(err)
					}
					for ; j < n; j++ {
						assembler.carBlocks[j].CarID = car.ID
					}
				}
				err = db.CreateInBatches(assembler.carBlocks, util2.BatchSize).Error
				if err != nil {
					return errors.WithStack(err)
				}
				return nil
			},
		)
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = database.DoRetry(ctx, func() error {
		if len(updatedFiles) == 0 {
			return nil
		}
		return db.Transaction(func(db *gorm.DB) error {
			var err error
			tree := daggen.NewDirectoryTree()
			var rootDirID uint64
			for _, file := range updatedFiles {
				dirID := file.DirectoryID
				for {
					// Add the directory to tree if it is not there
					if !tree.Has(*dirID) {
						var dir model.Directory
						err = db.Where("id = ?", dirID).First(&dir).Error
						if err != nil {
							return errors.Wrap(err, "failed to get directory")
						}
						err = tree.Add(ctx, &dir)
						if err != nil {
							return errors.Wrap(err, "failed to add directory to tree")
						}
					}

					dirDetail := tree.Get(*dirID)

					// Update the directory for first iteration
					if dirID == file.DirectoryID {
						err = dirDetail.Data.AddFile(ctx, file.FileName(), cid.Cid(file.CID), uint64(file.Size))
						if err != nil {
							return errors.Wrap(err, "failed to add file to directory")
						}
						if blks, ok := splitFileBlks[file.ID]; ok {
							err = dirDetail.Data.AddBlocks(ctx, blks)
							if err != nil {
								return errors.Wrap(err, "failed to add blocks to directory")
							}
						}
					}

					// Next iteration
					dirID = dirDetail.Dir.ParentID
					if dirID == nil {
						rootDirID = dirDetail.Dir.ID
						break
					}
				}
			}
			// Recursively update all directory internal structure
			_, err = tree.Resolve(ctx, rootDirID)
			if err != nil {
				return errors.Wrap(err, "failed to resolve directory tree")
			}

			// Update all directories in the database
			for dirID, dirDetail := range tree.Cache() {
				bytes, err := dirDetail.Data.MarshalBinary(ctx)
				if err != nil {
					return errors.Wrap(err, "failed to marshall directory data")
				}
				node, _ := dirDetail.Data.Node()
				err = db.Model(&model.Directory{}).Where("id = ?", dirID).Updates(map[string]any{
					"cid":      model.CID(node.Cid()),
					"data":     bytes,
					"exported": false,
				}).Error
				if err != nil {
					return errors.Wrap(err, "failed to update directory")
				}
			}
			return nil
		})
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to update directory CIDs")
	}

	logger.With("jobsID", job.ID).Info("finished packing")
	if job.Attachment.Preparation.DeleteAfterExport && len(job.Attachment.Preparation.OutputStorages) > 0 {
		logger.Info("Deleting original data source")
		for _, file := range updatedFiles {
			object := assembler.objects[file.ID]
			logger.Debugw("removing object", "path", object.Remote())
			err = object.Remove(ctx)
			if err != nil {
				logger.Warnw("failed to remove object", "error", err)
			}
		}
	}

	return cars, nil
}
