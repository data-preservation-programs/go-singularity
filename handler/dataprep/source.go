package dataprep

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"gorm.io/gorm"
)

// AddSourceStorageHandler associates a given source storage to a Preparation based on the provided ID.
// It first checks if the source storage exists. If it does, it then creates an association
// between the source storage and the specified Preparation.
//
// Parameters:
// - ctx: The context for database transactions and other operations.
// - db: A pointer to the gorm.DB instance representing the database connection.
// - id: The ID of the Preparation to which the source storage should be attached.
// - source: The name of the source storage to be attached.
//
// Returns:
// - A pointer to the updated Preparation model with the new source storage associated.
// - An error, if any occurred during the verification or attachment process.
//
// Note:
// This function ensures that the given source storage exists and that the given Preparation exists
// before creating an association. It also ensures there are no duplicate associations and handles
// potential errors accordingly.
func AddSourceStorageHandler(ctx context.Context, db *gorm.DB, id uint32, source string) (*model.Preparation, error) {
	db = db.WithContext(ctx)
	var storage model.Storage
	err := db.Where("name = ?", source).First(&storage).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "source storage '%s' does not exist", source)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = database.DoRetry(ctx, func() error {
		return db.Create(&model.SourceAttachment{
			StorageID:     storage.ID,
			PreparationID: id,
		}).Error
	})
	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation %d does not exist", id)
	}
	if util.IsDuplicateKeyError(err) {
		return nil, errors.Wrapf(handlererror.ErrDuplicateRecord, "source storage %s is already attached to preparation %d", source, id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var preparation model.Preparation
	err = db.Preload("SourceStorages").Preload("OutputStorages").First(&preparation, id).Error

	return &preparation, errors.WithStack(err)
}
