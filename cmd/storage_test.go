package cmd

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func swapStorageHandler(mockHandler storage.Handler) func() {
	actual := storage.Default
	storage.Default = mockHandler
	return func() {
		storage.Default = actual
	}
}

func TestStorageExploreHandler(t *testing.T) {
	runner := Runner{}
	defer runner.Save(t)
	ctx := context.Background()
	mockHandler := new(MockStorage)
	defer swapStorageHandler(mockHandler)()
	mockHandler.On("ExploreHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]storage.DirEntry{
		{
			Path:         "path1",
			LastModified: time.Time{},
			Size:         100,
			IsDir:        true,
			DirID:        "xx",
			NumItems:     100,
			Hash:         "hash1",
		},
		{
			Path:         "path2",
			LastModified: time.Time{},
			Size:         100,
			IsDir:        false,
			DirID:        "",
			NumItems:     -1,
			Hash:         "hash2",
		},
	}, nil)
	_, _, err := runner.Run(ctx, "singularity storage explore name path")
	require.NoError(t, err)

	_, _, err = runner.Run(ctx, "singularity --verbose storage explore name path")
	require.NoError(t, err)

}

func TestStorageCreateHandler_S3Provider(t *testing.T) {
	runner := Runner{}
	defer runner.Save(t)
	ctx := context.Background()
	mockHandler := new(MockStorage)
	defer swapStorageHandler(mockHandler)()
	mockHandler.On("CreateStorageHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Storage{
		ID:        1,
		Name:      "name",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Type:      "s3",
		Path:      "bucket",
		Config:    map[string]string{"region": "us-east-1"},
	}, nil)
	_, _, err := runner.Run(ctx, "singularity storage create s3 aws --region us-east-1 name bucket")
	require.NoError(t, err)

	_, _, err = runner.Run(ctx, "singularity --verbose storage create s3 aws --region us-east-1 name bucket")
	require.NoError(t, err)

}

func TestStorageCreateHandler_Local(t *testing.T) {
	runner := Runner{}
	defer runner.Save(t)
	ctx := context.Background()
	mockHandler := new(MockStorage)
	defer swapStorageHandler(mockHandler)()
	mockHandler.On("CreateStorageHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Storage{
		ID:        1,
		Name:      "name",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Type:      "local",
		Path:      "path",
	}, nil)
	_, _, err := runner.Run(ctx, "singularity storage create local name path")
	require.NoError(t, err)

	_, _, err = runner.Run(ctx, "singularity --verbose storage create local name path")
	require.NoError(t, err)

}

func TestStorageListHandler(t *testing.T) {
	runner := Runner{}
	defer runner.Save(t)
	ctx := context.Background()
	mockHandler := new(MockStorage)
	defer swapStorageHandler(mockHandler)()
	mockHandler.On("ListStoragesHandler", mock.Anything, mock.Anything).Return([]model.Storage{{
		ID:        1,
		Name:      "name1",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Type:      "local",
		Path:      "path",
		PreparationsAsSource: []model.Preparation{
			{
				ID:                1,
				CreatedAt:         time.Time{},
				UpdatedAt:         time.Time{},
				DeleteAfterExport: true,
				MaxSize:           100,
				PieceSize:         200,
			},
		},
		PreparationsAsOutput: []model.Preparation{
			{
				ID:                2,
				CreatedAt:         time.Time{},
				UpdatedAt:         time.Time{},
				DeleteAfterExport: true,
				MaxSize:           300,
				PieceSize:         400,
			}},
	}, {
		ID:        2,
		Name:      "name",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Type:      "local",
		Path:      "path",
	}}, nil)
	_, _, err := runner.Run(ctx, "singularity storage list")
	require.NoError(t, err)

	_, _, err = runner.Run(ctx, "singularity --verbose storage list")
	require.NoError(t, err)

}

func TestStorageRemoveHandler(t *testing.T) {
	runner := Runner{}
	defer runner.Save(t)
	ctx := context.Background()
	mockHandler := new(MockStorage)
	defer swapStorageHandler(mockHandler)()
	mockHandler.On("RemoveHandler", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	_, _, err := runner.Run(ctx, "singularity storage remove name")
	require.NoError(t, err)
}

func TestStorageUpdateHandler_S3Provider(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := Runner{}
		defer runner.Save(t)
		err := db.Create(&model.Storage{
			Name:   "name",
			Type:   "s3",
			Config: map[string]string{"provider": "AWS"},
			Path:   "bucket",
		}).Error
		require.NoError(t, err)
		mockHandler := new(MockStorage)
		defer swapStorageHandler(mockHandler)()
		mockHandler.On("UpdateStorageHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Storage{
			ID:        1,
			Name:      "name",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			Type:      "s3",
			Path:      "bucket",
			Config:    map[string]string{"region": "us-east-1"},
		}, nil)
		_, _, err = runner.Run(ctx, "singularity storage update s3 aws --region us-east-1 name")
		require.NoError(t, err)
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose storage update s3 aws --region us-east-1 name")
		require.NoError(t, err)

	})
}
func TestStorageUpdateHandler_Local(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := Runner{}
		defer runner.Save(t)
		err := db.Create(&model.Storage{
			Name: "name",
			Type: "local",
			Path: "/tmp",
		}).Error
		require.NoError(t, err)
		mockHandler := new(MockStorage)
		defer swapStorageHandler(mockHandler)()
		mockHandler.On("UpdateStorageHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Storage{
			ID:        1,
			Name:      "name",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			Type:      "local",
			Path:      "/tmp",
		}, nil)
		_, _, err = runner.Run(ctx, "singularity storage update local name")
		require.NoError(t, err)
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose storage update local name")
		require.NoError(t, err)

	})
}