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

func (m *Mock) GetRepositoryByName(_ uuid.UUID, _ string) (*repositoryEntities.Repository, error) {
	args := m.MethodCalled("GetRepositoryByName")
	return args.Get(0).(*repositoryEntities.Repository), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetRepository(_ uuid.UUID) (*repositoryEntities.Repository, error) {
	args := m.MethodCalled("GetRepository")
	return args.Get(0).(*repositoryEntities.Repository), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetAccountRepository(_, _ uuid.UUID) (*repositoryEntities.AccountRepository, error) {
	args := m.MethodCalled("GetAccountRepository")
	return args.Get(0).(*repositoryEntities.AccountRepository), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) ListRepositoriesAuthTypeHorusec(_, _ uuid.UUID) (*[]repositoryEntities.Response, error) {
	args := m.MethodCalled("ListRepositoriesAuthTypeHorusec")
	return args.Get(0).(*[]repositoryEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) ListRepositoriesAuthTypeLdap(_ uuid.UUID, _ []string) (*[]repositoryEntities.Response, error) {
	args := m.MethodCalled("ListRepositoriesAuthTypeLdap")
	return args.Get(0).(*[]repositoryEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) IsNotMemberOfWorkspace(_, _ uuid.UUID) bool {
	args := m.MethodCalled("IsNotMemberOfWorkspace")
	return args.Get(0).(bool)
}

func (m *Mock) ListAllRepositoryUsers(_ uuid.UUID) (*[]roleEntities.Response, error) {
	args := m.MethodCalled("ListAllRepositoryUsers")
	return args.Get(0).(*[]roleEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}
