package workspace

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"

	workspaceController "github.com/ZupIT/horusec-platform/core/internal/controllers/workspace"
	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
	workspaceUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/workspace"
)

func TestNewWorkspaceHandler(t *testing.T) {
	t.Run("should success create a new workspace handler", func(t *testing.T) {
		assert.NotNil(t, NewWorkspaceHandler(nil, nil, nil, nil))
	})
}

func TestCreate(t *testing.T) {
	workspaceData := &workspaceEntities.Data{
		Name:        "test",
		Description: "test",
		AuthzMember: []string{"test"},
		AuthzAdmin:  []string{"test"},
	}

	accountData := &proto.GetAccountDataResponse{
		AccountID:   uuid.New().String(),
		Permissions: []string{"test"},
	}

	authConfig := &proto.GetAuthConfigResponse{}

	t.Run("should return 200 when everything it is ok", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		controllerMock.On("Create").Return(&workspaceEntities.Workspace{}, nil)

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)
		authGRPCMock.On("GetAuthConfig").Return(authConfig, nil)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, app.NewAppConfig(authGRPCMock))

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(workspaceData.ToBytes()))
		w := httptest.NewRecorder()

		handler.Create(w, r)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("should return 500 when something went wrong while creating workspace", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		controllerMock.On("Create").Return(&workspaceEntities.Workspace{}, errors.New("test"))

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)
		authGRPCMock.On("GetAuthConfig").Return(authConfig, nil)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, app.NewAppConfig(authGRPCMock))

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(workspaceData.ToBytes()))
		w := httptest.NewRecorder()

		handler.Create(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when failed to get account data", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, errors.New("test"))
		authGRPCMock.On("GetAuthConfig").Return(authConfig, nil)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, app.NewAppConfig(authGRPCMock))

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(workspaceData.ToBytes()))
		w := httptest.NewRecorder()

		handler.Create(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when failed to get workspace data from request body", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAuthConfig").Return(authConfig, nil)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, app.NewAppConfig(authGRPCMock))

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader([]byte("")))
		w := httptest.NewRecorder()

		handler.Create(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when invalid ldap groups", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		controllerMock.On("Create").Return(&workspaceEntities.Workspace{}, nil)

		authGRPCMock := &proto.Mock{}
		accountData.Permissions = []string{"test2"}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)
		authConfig.AuthType = auth.Ldap.ToString()
		authGRPCMock.On("GetAuthConfig").Return(authConfig, nil)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, app.NewAppConfig(authGRPCMock))

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(workspaceData.ToBytes()))
		w := httptest.NewRecorder()

		handler.Create(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
