package token

import (
	"io"

	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/core/internal/entities/token"
)

type IUseCases interface {
	TokenDataFromIOReadCloser(body io.ReadCloser) (*token.Data, error)
}

type UseCases struct {
}

func NewTokenUseCases() IUseCases {
	return &UseCases{}
}

func (u *UseCases) TokenDataFromIOReadCloser(body io.ReadCloser) (*token.Data, error) {
	data := &token.Data{}

	if err := parser.ParseBodyToEntity(body, data); err != nil {
		return nil, err
	}

	return data, data.Validate()
}
