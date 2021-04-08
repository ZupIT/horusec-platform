package repository

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	mockUtils "github.com/ZupIT/horusec-devkit/pkg/utils/mock"

	repositoryEntities "github.com/ZupIT/horusec-platform/core/internal/entities/repository"
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
