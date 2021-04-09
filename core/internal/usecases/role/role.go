package role

import (
	"io"

	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/core/internal/entities/role"
)

type IUseCases interface {
	NewRoleData(accountID, workspaceID, repositoryID uuid.UUID) *role.Data
	InviteUserDataFromIOReadCloser(body io.ReadCloser) (*role.InviteUserData, error)
	RoleDataFromIOReadCloser(body io.ReadCloser) (*role.Data, error)
}

type UseCases struct {
}

func NewRoleUseCases() IUseCases {
	return &UseCases{}
}

func (u *UseCases) NewRoleData(accountID, workspaceID, repositoryID uuid.UUID) *role.Data {
	return &role.Data{
		WorkspaceID:  workspaceID,
		RepositoryID: repositoryID,
		AccountID:    accountID,
	}
}

func (u *UseCases) InviteUserDataFromIOReadCloser(body io.ReadCloser) (*role.InviteUserData, error) {
	data := &role.InviteUserData{}

	if err := parser.ParseBodyToEntity(body, data); err != nil {
		return nil, err
	}

	return data, data.Validate()
}

func (u *UseCases) RoleDataFromIOReadCloser(body io.ReadCloser) (*role.Data, error) {
	data := &role.Data{}

	if err := parser.ParseBodyToEntity(body, data); err != nil {
		return nil, err
	}

	return data, data.Validate()
}
