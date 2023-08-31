package storage

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	t.Run("not supported storage type", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			_, err := Default.CreateStorageHandler(ctx, db, "not_supported", CreateRequest{
				"", "", "", nil,
			})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})
	})
	t.Run("local path", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			tmp := t.TempDir()
			storage, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "name", tmp, nil})
			require.NoError(t, err)
			require.Greater(t, storage.ID, uint32(0))
		})
	})
	t.Run("local path with config", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			tmp := t.TempDir()
			storage, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "name", tmp,
				map[string]string{
					"copy_links": "true",
				}})
			require.NoError(t, err)
			require.Greater(t, storage.ID, uint32(0))
		})
	})

	t.Run("local path with invalid config", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			tmp := t.TempDir()
			_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "name", tmp,
				map[string]string{
					"copy_links": "invalid",
				}})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})
	})

	t.Run("local path with inaccessible path", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "name", "/invalid/path", nil})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})
	})

	t.Run("invalid provider", func(t *testing.T) {
		tmp := t.TempDir()
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"invalid", "name", tmp, nil})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})
	})

	t.Run("duplicate name", func(t *testing.T) {
		tmp := t.TempDir()
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "name", tmp, nil})
			require.NoError(t, err)
			_, err = Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "name", tmp, nil})
			require.ErrorIs(t, err, handlererror.ErrDuplicateRecord)
		})
	})
}