package repository

import (
	"io"

	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/core/internal/entities/repository"
)

type IUseCases interface {
	RepositoryDataFromIOReadCloser(body io.ReadCloser) (*repository.Data, error)
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
