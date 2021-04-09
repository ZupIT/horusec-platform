package repository

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	mockUtils "github.com/ZupIT/horusec-devkit/pkg/utils/mock"

	repositoryEntities "github.com/ZupIT/horusec-platform/core/internal/entities/repository"
	roleEntities "github.com/ZupIT/horusec-platform/core/internal/entities/role"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) Create(_ *repositoryEntities.Data) (*repositoryEntities.Response, error) {
	args := m.MethodCalled("Create")
	return args.Get(0).(*repositoryEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) Get(_ *repositoryEntities.Data) (*repositoryEntities.Response, error) {
	args := m.MethodCalled("Get")
	return args.Get(0).(*repositoryEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) Update(_ *repositoryEntities.Data) (*repositoryEntities.Response, error) {
	args := m.MethodCalled("Update")
	return args.Get(0).(*repositoryEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) Delete(_ uuid.UUID) error {
	args := m.MethodCalled("Delete")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (m *Mock) List(_ *repositoryEntities.Data) (*[]repositoryEntities.Response, error) {
	args := m.MethodCalled("List")
	return args.Get(0).(*[]repositoryEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) UpdateRole(_ *roleEntities.Data) (*roleEntities.Response, error) {
	args := m.MethodCalled("UpdateRole")
	return args.Get(0).(*roleEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) InviteUser(_ *roleEntities.UserData) (*roleEntities.Response, error) {
	args := m.MethodCalled("InviteUser")
	return args.Get(0).(*roleEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}
