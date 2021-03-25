package workspace

import (
	"net/http"

	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
)

type IUseCases interface {
	GetCreateWorkspaceData(r *http.Request) (data *workspace.CreateWorkspaceData, err error)
}

type UseCases struct {
}

func NewWorkspaceUseCases() IUseCases {
	return &UseCases{}
}

func (u *UseCases) GetCreateWorkspaceData(r *http.Request) (*workspace.CreateWorkspaceData, error) {
	data := &workspace.CreateWorkspaceData{}
	err := parser.ParseBodyToEntity(r.Body, data)
	if err != nil {
		return nil, err
	}

	return data, data.Validate()
}
