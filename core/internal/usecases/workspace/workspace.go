package workspace

import (
	"io"

	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
)

type IUseCases interface {
	GetCreateWorkspaceData(body io.ReadCloser) (data *workspace.CreateWorkspaceData, err error)
}

type UseCases struct {
}

func NewWorkspaceUseCases() IUseCases {
	return &UseCases{}
}

func (u *UseCases) GetCreateWorkspaceData(body io.ReadCloser) (*workspace.CreateWorkspaceData, error) {
	data := &workspace.CreateWorkspaceData{}
	if err := parser.ParseBodyToEntity(body, data); err != nil {
		return nil, err
	}

	return data, data.Validate()
}
