package wallet

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// AttachHandler associates a wallet with a specific preparation based on given preparationID and wallet address or ID.
//
// Parameters:
// - ctx: The context for database transactions and other operations.
// - db: A pointer to the gorm.DB instance representing the database connection.
// - preparationID: The ID of the preparation to which the wallet will be attached.
// - wallet: The address or ID of the wallet to be attached to the preparation.
//
// Returns:
// - A pointer to the updated Preparation instance.
// - An error, if any occurred during the association operation.
func AttachHandler(
	ctx context.Context,
	db *gorm.DB,
	preparationID uint32,
	wallet string,
) (*model.Preparation, error) {
	db = db.WithContext(ctx)
	var preparation model.Preparation
	err := db.Preload("SourceStorages").Preload("OutputStorages").Preload("Wallets").First(&preparation, preparationID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation %d not found", preparationID)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var w model.Wallet
	err = db.Where("address = ? OR id = ?", wallet, wallet).First(&w).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "wallet %s not found", wallet)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = database.DoRetry(ctx, func() error {
		return db.Model(&preparation).Association("Wallets").Append(&w)
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &preparation, nil
}

// @Summary Attach a new wallet with a preparation
// @Tags Wallet Association
// @Produce json
// @Accept json
// @Param id path int true "Preparation ID"
// @Param wallet path string true "Wallet Address"
// @Success 200 {object} model.Preparation
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/wallet/{wallet} [post]
func _() {}
