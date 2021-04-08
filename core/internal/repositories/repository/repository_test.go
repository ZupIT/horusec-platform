package repository

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"

	repositoryEntities "github.com/ZupIT/horusec-platform/core/internal/entities/repository"
	repositoryUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/repository"
)

func TestNewRepositoryRepository(t *testing.T) {
	t.Run("should success create a repository repository", func(t *testing.T) {
		assert.NotNil(t, NewRepositoryRepository(&database.Connection{}, repositoryUseCases.NewRepositoryUseCases()))
	})
}

func TestGetRepositoryByName(t *testing.T) {
	t.Run("should success get repository by name", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Find").
			Return(response.NewResponse(1, nil, &repositoryEntities.Repository{}))

		repository := NewRepositoryRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			repositoryUseCases.NewRepositoryUseCases())

		result, err := repository.GetRepositoryByName(uuid.New(), "test")
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestGetRepository(t *testing.T) {
	t.Run("should success get a repository", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Find").
			Return(response.NewResponse(1, nil, &repositoryEntities.Repository{}))

		repository := NewRepositoryRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			repositoryUseCases.NewRepositoryUseCases())

		result, err := repository.GetRepository(uuid.New())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestGetAccountRepository(t *testing.T) {
	t.Run("should success get a account repository", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Find").
			Return(response.NewResponse(1, nil, &repositoryEntities.Repository{}))

		repository := NewRepositoryRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			repositoryUseCases.NewRepositoryUseCases())

		result, err := repository.GetAccountRepository(uuid.New(), uuid.New())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}
