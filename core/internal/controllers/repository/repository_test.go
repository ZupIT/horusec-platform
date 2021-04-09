package repository

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
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

func TestUpdate(t *testing.T) {
	data := &repositoryEntities.Data{
		AccountID:       uuid.New(),
		Name:            "test",
		Description:     "test",
		AuthzMember:     []string{"test"},
		AuthzAdmin:      []string{"test"},
		AuthzSupervisor: []string{"test"},
		Permissions:     []string{"test"},
	}

	t.Run("should success update repository", func(t *testing.T) {
		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepository").Return(&repositoryEntities.Repository{Name: "test2"}, nil)
		repositoryMock.On("GetRepositoryByName").Return(
			&repositoryEntities.Repository{}, databaseEnums.ErrorNotFoundRecords)

		databaseMock := &database.Mock{}
		databaseMock.On("Update").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock)

		result, err := controller.Update(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error name already in use", func(t *testing.T) {
		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepository").Return(&repositoryEntities.Repository{}, nil)
		repositoryMock.On("GetRepositoryByName").Return(
			&repositoryEntities.Repository{}, nil)

		databaseMock := &database.Mock{}
		databaseMock.On("Update").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock)

		_, err := controller.Update(data)
		assert.Error(t, err)
	})

	t.Run("should return error while getting repository", func(t *testing.T) {
		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepository").Return(&repositoryEntities.Repository{}, errors.New("test"))

		databaseMock := &database.Mock{}
		databaseMock.On("Update").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock)

		_, err := controller.Update(data)
		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should success delete repository", func(t *testing.T) {
		repositoryMock := &repositoryRepository.Mock{}
		appConfig := &app.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Delete").Return(&response.Response{})

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock)

		assert.NoError(t, controller.Delete(uuid.New()))
	})
}

func TestList(t *testing.T) {
	data := &repositoryEntities.Data{
		AccountID:   uuid.New(),
		WorkspaceID: uuid.New(),
		Permissions: []string{"test"},
	}

	t.Run("should success list repositories when horusec auth type", func(t *testing.T) {
		databaseMock := &database.Mock{}

		appConfig := &app.Mock{}
		appConfig.On("GetAuthorizationType").Return(auth.Horusec)

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("ListRepositoriesAuthTypeHorusec").Return(&[]repositoryEntities.Response{}, nil)

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock)

		result, err := controller.List(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should success list repositories when ldap auth type", func(t *testing.T) {
		databaseMock := &database.Mock{}

		appConfig := &app.Mock{}
		appConfig.On("GetAuthorizationType").Return(auth.Ldap)

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("ListRepositoriesAuthTypeLdap").Return(&[]repositoryEntities.Response{}, nil)

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock)

		result, err := controller.List(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}
