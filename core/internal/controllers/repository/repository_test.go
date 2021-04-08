package repository

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseEnums "github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"

	repositoryEntities "github.com/ZupIT/horusec-platform/core/internal/entities/repository"
	repositoryRepository "github.com/ZupIT/horusec-platform/core/internal/repositories/repository"
	repositoryUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/repository"
)

func TestCreate(t *testing.T) {
	data := &repositoryEntities.Data{
		AccountID:   uuid.New(),
		Name:        "test",
		Description: "test",
		AuthzMember: []string{"test"},
		AuthzAdmin:  []string{"test"},
		Permissions: []string{"test"},
	}

	t.Run("should success create a new repository", func(t *testing.T) {
		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepositoryByName").Return(
			&repositoryEntities.Repository{}, databaseEnums.ErrorNotFoundRecords)

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})
		databaseMock.On("StartTransaction").Return(databaseMock)
		databaseMock.On("CommitTransaction").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock)

		result, err := controller.Create(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when creating account repository", func(t *testing.T) {
		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepositoryByName").Return(
			&repositoryEntities.Repository{}, databaseEnums.ErrorNotFoundRecords)

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Once().Return(&response.Response{})
		databaseMock.On("Create").Return(
			response.NewResponse(0, errors.New("test"), nil))
		databaseMock.On("StartTransaction").Return(databaseMock)
		databaseMock.On("RollbackTransaction").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock)

		_, err := controller.Create(data)
		assert.Error(t, err)
	})

	t.Run("should return error when creating repository", func(t *testing.T) {
		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepositoryByName").Return(
			&repositoryEntities.Repository{}, databaseEnums.ErrorNotFoundRecords)

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(
			response.NewResponse(0, errors.New("test"), nil))
		databaseMock.On("StartTransaction").Return(databaseMock)
		databaseMock.On("RollbackTransaction").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock)

		_, err := controller.Create(data)
		assert.Error(t, err)
	})

	t.Run("should return error name already in use", func(t *testing.T) {
		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepositoryByName").Return(
			&repositoryEntities.Repository{}, nil)

		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock)

		_, err := controller.Create(data)
		assert.Error(t, err)
	})
}

func TestGet(t *testing.T) {
	data := &repositoryEntities.Data{
		AccountID:    uuid.New(),
		RepositoryID: uuid.New(),
	}

	t.Run("should success get a repository", func(t *testing.T) {
		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepository").Return(&repositoryEntities.Repository{}, nil)
		repositoryMock.On("GetAccountRepository").Return(&repositoryEntities.AccountRepository{}, nil)

		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock)

		result, err := controller.Get(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when failed to get repository", func(t *testing.T) {
		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepository").Return(&repositoryEntities.Repository{}, errors.New("test"))
		repositoryMock.On("GetAccountRepository").Return(&repositoryEntities.AccountRepository{}, nil)

		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock)

		_, err := controller.Get(data)
		assert.Error(t, err)
	})

	t.Run("should return error when failed to get account repository", func(t *testing.T) {
		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetAccountRepository").Return(
			&repositoryEntities.AccountRepository{}, errors.New("test"))

		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock)

		_, err := controller.Get(data)
		assert.Error(t, err)
	})
}
