package workspace

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"

	"github.com/ZupIT/horusec-platform/core/internal/entities/role"
	tokenEntities "github.com/ZupIT/horusec-platform/core/internal/entities/token"
	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
	workspaceRepository "github.com/ZupIT/horusec-platform/core/internal/repositories/workspace"
	tokenUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/token"
	workspaceUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/workspace"
)

func TestNewWorkspaceController(t *testing.T) {
	t.Run("should success create a new workspace controller", func(t *testing.T) {
		assert.NotNil(t, NewWorkspaceController(&broker.Broker{}, &database.Connection{}, &app.Config{},
			workspaceUseCases.NewWorkspaceUseCases(), &workspaceRepository.Repository{}, tokenUseCases.NewTokenUseCases()))
	})
}

func TestCreate(t *testing.T) {
	workspaceData := &workspaceEntities.Data{
		AccountID:   uuid.New(),
		Name:        "test",
		Description: "test",
		AuthzMember: []string{"test"},
		AuthzAdmin:  []string{"test"},
		Permissions: []string{"test"},
	}

	t.Run("should success create a new workspace", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})
		databaseMock.On("StartTransaction").Return(databaseMock)
		databaseMock.On("CommitTransaction").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		result, err := controller.Create(workspaceData)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when creating account workspace relation", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Once().Return(&response.Response{})
		databaseMock.On("Create").Return(response.NewResponse(1,
			errors.New("test"), nil))
		databaseMock.On("StartTransaction").Return(databaseMock)
		databaseMock.On("RollbackTransaction").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		result, err := controller.Create(workspaceData)
		assert.Error(t, err)
		assert.Equal(t, errors.New("test"), err)
		assert.Nil(t, result)
	})

	t.Run("should return error when creating workspace", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(response.NewResponse(1,
			errors.New("test"), nil))
		databaseMock.On("StartTransaction").Return(databaseMock)
		databaseMock.On("RollbackTransaction").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		result, err := controller.Create(workspaceData)
		assert.Error(t, err)
		assert.Equal(t, errors.New("test"), err)
		assert.Nil(t, result)
	})
}

func TestGet(t *testing.T) {
	workspaceData := &workspaceEntities.Data{
		AccountID:   uuid.New(),
		Name:        "test",
		Description: "test",
		AuthzMember: []string{"test"},
		AuthzAdmin:  []string{"test"},
		Permissions: []string{"test"},
	}

	accountWorkspace := &workspaceEntities.AccountWorkspace{
		WorkspaceID: uuid.New(),
		AccountID:   uuid.New(),
		Role:        account.Admin,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	workspace := &workspaceEntities.Workspace{
		WorkspaceID: uuid.New(),
		Name:        "test",
		Description: "test",
		AuthzMember: []string{"test"},
		AuthzAdmin:  []string{"test"},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("should success get workspace with role", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		repositoryMock.On("GetAccountWorkspace").Return(accountWorkspace, nil)
		repositoryMock.On("GetWorkspace").Return(workspace, nil)

		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		result, err := controller.Get(workspaceData)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when failed to get workspace", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		repositoryMock.On("GetAccountWorkspace").Return(accountWorkspace, nil)
		repositoryMock.On("GetWorkspace").Return(workspace, errors.New("test"))

		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		_, err := controller.Get(workspaceData)
		assert.Error(t, err)
		assert.Equal(t, errors.New("test"), err)
	})

	t.Run("should return error when failed to get workspace", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		repositoryMock.On("GetAccountWorkspace").Return(accountWorkspace, errors.New("test"))

		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		_, err := controller.Get(workspaceData)
		assert.Error(t, err)
		assert.Equal(t, errors.New("test"), err)
	})

	t.Run("should success get workspace when application admin", func(t *testing.T) {
		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("GetWorkspace").Return(workspace, nil)

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		workspaceData.IsApplicationAdmin = true
		result, err := controller.Get(workspaceData)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when failed to get workspace and user is application admin", func(t *testing.T) {
		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("GetWorkspace").Return(workspace, errors.New("test"))

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		workspaceData.IsApplicationAdmin = true
		_, err := controller.Get(workspaceData)
		assert.Error(t, err)
		assert.Equal(t, errors.New("test"), err)
	})
}

func TestUpdate(t *testing.T) {
	workspaceData := &workspaceEntities.Data{
		AccountID:   uuid.New(),
		Name:        "test",
		Description: "test",
		AuthzMember: []string{"test"},
		AuthzAdmin:  []string{"test"},
		Permissions: []string{"test"},
	}

	workspace := &workspaceEntities.Workspace{
		WorkspaceID: uuid.New(),
		Name:        "test",
		Description: "test",
		AuthzMember: []string{"test"},
		AuthzAdmin:  []string{"test"},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("should success update workspace", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("GetWorkspace").Return(workspace, nil)

		databaseMock := &database.Mock{}
		databaseMock.On("Update").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		result, err := controller.Update(workspaceData)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when failed to get workspace", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("GetWorkspace").Return(workspace, errors.New("test"))

		databaseMock := &database.Mock{}
		databaseMock.On("Update").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		_, err := controller.Update(workspaceData)
		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should success delete workspace", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Delete").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		assert.NoError(t, controller.Delete(uuid.New()))
	})
}

func TestList(t *testing.T) {
	workspaceData := &workspaceEntities.Data{
		AccountID:   uuid.New(),
		Name:        "test",
		Description: "test",
		AuthzMember: []string{"test"},
		AuthzAdmin:  []string{"test"},
		Permissions: []string{"test"},
	}

	workspaceResponse := &[]workspaceEntities.Response{
		{
			WorkspaceID: uuid.New(),
			Name:        "test",
			Role:        account.Admin,
			Description: "test",
			AuthzMember: []string{"test"},
			AuthzAdmin:  []string{"test"},
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	t.Run("should list workspaces when horusec auth type", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("ListWorkspacesAuthTypeHorusec").Return(workspaceResponse, nil)

		appConfig := &app.Mock{}
		appConfig.On("GetAuthenticationType").Return(auth.Horusec)

		databaseMock := &database.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		result, err := controller.List(workspaceData)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should list workspaces when ldap auth type", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("ListWorkspacesAuthTypeLdap").Return(workspaceResponse, nil)

		appConfig := &app.Mock{}
		appConfig.On("GetAuthenticationType").Return(auth.Ldap)

		databaseMock := &database.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		result, err := controller.List(workspaceData)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when failed to list with horusec auth type", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("ListWorkspacesAuthTypeHorusec").Return(
			workspaceResponse, errors.New("test"))

		appConfig := &app.Mock{}
		appConfig.On("GetAuthenticationType").Return(auth.Horusec)

		databaseMock := &database.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		_, err := controller.List(workspaceData)
		assert.Error(t, err)
	})

	t.Run("should return error when failed to list with ldap auth type", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("ListWorkspacesAuthTypeLdap").Return(workspaceResponse, errors.New("test"))

		appConfig := &app.Mock{}
		appConfig.On("GetAuthenticationType").Return(auth.Ldap)

		databaseMock := &database.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		_, err := controller.List(workspaceData)
		assert.Error(t, err)
	})

	t.Run("should list workspaces when application admin", func(t *testing.T) {
		appConfig := &app.Mock{}
		databaseMock := &database.Mock{}

		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("ListWorkspacesApplicationAdmin").Return(workspaceResponse, nil)

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		workspaceData.IsApplicationAdmin = true
		result, err := controller.List(workspaceData)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestUpdateRole(t *testing.T) {
	data := &role.Data{
		Role: account.Admin,
	}

	accountWorkspace := &workspaceEntities.AccountWorkspace{
		WorkspaceID: uuid.New(),
		AccountID:   uuid.New(),
		Role:        account.Admin,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("should success update user role", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("GetAccountWorkspace").Return(accountWorkspace, nil)

		databaseMock := &database.Mock{}
		databaseMock.On("Update").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		result, err := controller.UpdateRole(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when failed to get account workspace", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("GetAccountWorkspace").Return(accountWorkspace, errors.New("test"))

		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		_, err := controller.UpdateRole(data)
		assert.Error(t, err)
	})
}

func TestInviteUser(t *testing.T) {
	data := &role.UserData{
		Role: account.Admin,
	}

	workspace := &workspaceEntities.Workspace{
		WorkspaceID: uuid.New(),
		Name:        "test",
		Description: "test",
		AuthzMember: []string{"test"},
		AuthzAdmin:  []string{"test"},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("should success create new account workspace without email", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("GetWorkspace").Return(workspace, nil)

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})

		appConfig := &app.Mock{}
		appConfig.On("IsEmailsDisabled").Return(true)

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		result, err := controller.InviteUser(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should success create new account workspace with email", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("GetWorkspace").Return(workspace, nil)

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})

		appConfig := &app.Mock{}
		appConfig.On("IsEmailsDisabled").Return(false)

		brokerMock := &broker.Mock{}
		brokerMock.On("Publish").Return(nil)

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(brokerMock, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		result, err := controller.InviteUser(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when failed to create account workspace", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("GetWorkspace").Return(workspace, nil)

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(
			response.NewResponse(0, errors.New("test"), nil))

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		_, err := controller.InviteUser(data)
		assert.Error(t, err)
	})

	t.Run("should return error when failed to get workspace", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("GetWorkspace").Return(workspace, errors.New("test"))

		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		_, err := controller.InviteUser(data)
		assert.Error(t, err)
	})
}

func TestGetUsers(t *testing.T) {
	usersResponse := &[]role.Response{
		{
			AccountID: uuid.New(),
			Email:     "test@test.com",
			Username:  "test",
			Role:      account.Admin,
		},
	}

	t.Run("should success get all users of workspace", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("ListAllWorkspaceUsers").Return(usersResponse, nil)

		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		result, err := controller.GetUsers(uuid.New())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestRemoveUser(t *testing.T) {
	data := &role.Data{
		AccountID:   uuid.New(),
		WorkspaceID: uuid.New(),
	}

	t.Run("should success remove user from repositories and workspace", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		appConfig := &app.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Delete").Return(&response.Response{})

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		assert.NoError(t, controller.RemoveUser(data))
	})

	t.Run("should return error when failed to remove user from workspace", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		appConfig := &app.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Delete").Once().Return(&response.Response{})
		databaseMock.On("Delete").Return(
			response.NewResponse(0, errors.New("test"), nil))

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		assert.Error(t, controller.RemoveUser(data))
	})

	t.Run("should return error when failed to remove user from repositories", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		appConfig := &app.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Delete").Return(
			response.NewResponse(0, errors.New("test"), nil))

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		assert.Error(t, controller.RemoveUser(data))
	})
}

func TestCreateToken(t *testing.T) {
	data := &tokenEntities.Data{}

	t.Run("should success create a new workspace token ", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		appConfig := &app.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		result, err := controller.CreateToken(data)
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("should return error while creating a new workspace token ", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		appConfig := &app.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(
			response.NewResponse(0, errors.New("test"), nil))

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		_, err := controller.CreateToken(data)
		assert.Error(t, err)
	})
}

func TestDeleteToken(t *testing.T) {
	t.Run("should success delete a workspace token ", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		appConfig := &app.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Delete").Return(&response.Response{})

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		assert.NoError(t, controller.DeleteToken(&tokenEntities.Data{}))
	})
}

func TestListTokens(t *testing.T) {
	t.Run("should success list workspace tokens", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		appConfig := &app.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Find").Return(&response.Response{})

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock, tokenUseCases.NewTokenUseCases())

		result, err := controller.ListTokens(uuid.New())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}
