package workspace

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"

	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
	workspaceRepository "github.com/ZupIT/horusec-platform/core/internal/repositories/workspace"
	workspaceUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/workspace"
)

func TestNewWorkspaceController(t *testing.T) {
	t.Run("should success create a new workspace controller", func(t *testing.T) {
		assert.NotNil(t, NewWorkspaceController(&broker.Broker{}, &database.Connection{},
			&app.Config{}, workspaceUseCases.NewWorkspaceUseCases(), &workspaceRepository.Repository{}))
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

	authConfig := &proto.GetAuthConfigResponse{}

	t.Run("should success create a new workspace with horusec auth type", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})
		databaseMock.On("StartTransaction").Return(databaseMock)
		databaseMock.On("CommitTransaction").Return(&response.Response{})

		authGRPCMock := &proto.Mock{}
		authConfig.AuthType = auth.Horusec.ToString()
		authGRPCMock.On("GetAuthConfig").Return(authConfig, nil)

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection,
			app.NewAppConfig(authGRPCMock), workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

		result, err := controller.Create(workspaceData)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should success create a new workspace with ldap auth type", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})
		databaseMock.On("StartTransaction").Return(databaseMock)
		databaseMock.On("CommitTransaction").Return(&response.Response{})

		authGRPCMock := &proto.Mock{}
		authConfig.AuthType = auth.Ldap.ToString()
		authGRPCMock.On("GetAuthConfig").Return(authConfig, nil)

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection,
			app.NewAppConfig(authGRPCMock), workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

		result, err := controller.Create(workspaceData)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should success create a new workspace with keycloak auth type", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})
		databaseMock.On("StartTransaction").Return(databaseMock)
		databaseMock.On("CommitTransaction").Return(&response.Response{})

		authGRPCMock := &proto.Mock{}
		authConfig.AuthType = auth.Keycloak.ToString()
		authGRPCMock.On("GetAuthConfig").Return(authConfig, nil)

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection,
			app.NewAppConfig(authGRPCMock), workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

		result, err := controller.Create(workspaceData)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should success create a new workspace with keycloak auth type", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})
		databaseMock.On("StartTransaction").Return(databaseMock)
		databaseMock.On("CommitTransaction").Return(&response.Response{})

		authGRPCMock := &proto.Mock{}
		authConfig.AuthType = auth.Keycloak.ToString()
		authGRPCMock.On("GetAuthConfig").Return(authConfig, nil)

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection,
			app.NewAppConfig(authGRPCMock), workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

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

		authGRPCMock := &proto.Mock{}
		authConfig.AuthType = auth.Horusec.ToString()
		authGRPCMock.On("GetAuthConfig").Return(authConfig, nil)

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection,
			app.NewAppConfig(authGRPCMock), workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

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

		authGRPCMock := &proto.Mock{}
		authConfig.AuthType = auth.Horusec.ToString()
		authGRPCMock.On("GetAuthConfig").Return(authConfig, nil)

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection,
			app.NewAppConfig(authGRPCMock), workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

		result, err := controller.Create(workspaceData)
		assert.Error(t, err)
		assert.Equal(t, errors.New("test"), err)
		assert.Nil(t, result)
	})
}
