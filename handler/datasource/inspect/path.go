package inspect

import (
	"context"
	"strconv"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
)

type DirDetail struct {
	Current model.Directory
	Dirs    []model.Directory
	Files   []model.File
}

type GetPathRequest struct {
	Path string `json:"path"`
}

func GetPathHandler(
	ctx context.Context,
	db *gorm.DB,
	id string,
	request GetPathRequest,
) (*DirDetail, error) {
	return getPathHandler(db.WithContext(ctx), id, request)
}

// @Summary Get all file details inside a data source path
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Param request body GetPathRequest true "GetPathRequest"
// @Success 200 {object} DirDetail
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /source/{id}/path [get]
func getPathHandler(
	db *gorm.DB,
	id string,
	request GetPathRequest,
) (*DirDetail, error) {
	path := request.Path
	sourceID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handlererror.NewInvalidParameterErr("invalid source id")
	}
	var source model.Source
	err = db.Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handlererror.NewInvalidParameterErr("source not found")
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	segments := underscore.Filter(strings.Split(path, "/"), func(s string) bool { return s != "" })
	err = source.LoadRootDirectory(db)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	dirID := source.RootDirectory().ID
	for i, segment := range segments {
		var subdir model.Directory
		// TODO: name is not indexed. We may not need to add index if this path is not used frequently.
		err = db.Where("parent_id = ? AND name = ?", dirID, segment).First(&subdir).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if i == len(segments)-1 {
				var files []model.File
				err = db.Where("source_id = ? AND path = ?", sourceID, strings.Join(segments, "/")).Find(&files).Error
				if err != nil {
					return nil, errors.WithStack(err)
				}
				if len(files) == 0 {
					return nil, handlererror.NewInvalidParameterErr("entry not found with given path")
				}
				return &DirDetail{
					Current: subdir,
					Dirs:    nil,
					Files:   files,
				}, nil
			}
			return nil, handlererror.NewInvalidParameterErr("entry not found with given path")
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}
		dirID = subdir.ID
	}

	var current model.Directory
	var dirs []model.Directory
	var files []model.File
	err = db.Where("id = ?", dirID).First(&current).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	err = db.Where("parent_id = ?", dirID).Find(&dirs).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	err = db.Where("directory_id = ?", dirID).Find(&files).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &DirDetail{
		Current: current,
		Dirs:    dirs,
		Files:   files,
	}, nil
}
