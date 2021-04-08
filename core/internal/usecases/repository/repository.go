package repository

import (
	"io"

	"github.com/google/uuid"

	databaseEnums "github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/core/internal/entities/repository"
)

type IUseCases interface {
	RepositoryDataFromIOReadCloser(body io.ReadCloser) (*repository.Data, error)
	FilterRepositoryByName(workspaceID uuid.UUID, name string) map[string]interface{}
	IsNotFoundError(err error) bool
}

type UseCases struct {
}

func NewRepositoryUseCases() IUseCases {
	return &UseCases{}
}

func (u *UseCases) RepositoryDataFromIOReadCloser(body io.ReadCloser) (*repository.Data, error) {
	data := &repository.Data{}

	if err := parser.ParseBodyToEntity(body, data); err != nil {
		return nil, err
	}

	return data, data.Validate()
}

func (u *UseCases) FilterRepositoryByName(workspaceID uuid.UUID, name string) map[string]interface{} {
	return map[string]interface{}{"workspace_id": workspaceID, "name": name}
}

func (u *UseCases) IsNotFoundError(err error) bool {
	if err != nil {
		if err == databaseEnums.ErrorNotFoundRecords {
			return true
		}
	}

	return false
}
