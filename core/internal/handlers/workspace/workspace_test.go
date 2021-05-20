package workspace

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"

	workspaceController "github.com/ZupIT/horusec-platform/core/internal/controllers/workspace"
	"github.com/ZupIT/horusec-platform/core/internal/entities/role"
	tokenEntities "github.com/ZupIT/horusec-platform/core/internal/entities/token"
	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
	roleUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/role"
	tokenUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/token"
	workspaceUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/workspace"
)

func TestNewWorkspaceHandler(t *testing.T) {
	t.Run("should success create a new workspace handler", func(t *testing.T) {
		assert.NotNil(t, NewWorkspaceHandler(nil, nil, nil,
			nil, nil, nil))
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

	t.Run("should return 201 when everything it is ok", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		controllerMock.On("Create").Return(&workspaceEntities.Response{}, nil)

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}
		appConfigMock.On("GetAuthenticationType").Return(auth.Horusec)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(workspaceData.ToBytes()))
		w := httptest.NewRecorder()

		handler.Create(w, r)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("should return 500 when something went wrong while creating workspace", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		controllerMock.On("Create").Return(&workspaceEntities.Response{}, errors.New("test"))

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}
		appConfigMock.On("GetAuthenticationType").Return(auth.Horusec)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(workspaceData.ToBytes()))
		w := httptest.NewRecorder()

		handler.Create(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when failed to get account data", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, errors.New("test"))

		appConfigMock := &app.Mock{}
		appConfigMock.On("GetAuthenticationType").Return(auth.Horusec)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(workspaceData.ToBytes()))
		w := httptest.NewRecorder()

		handler.Create(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when failed to get workspace data from request body", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}
		appConfigMock.On("GetAuthenticationType").Return(auth.Horusec)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader([]byte("")))
		w := httptest.NewRecorder()

		handler.Create(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when invalid ldap groups", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}

		authGRPCMock := &proto.Mock{}

		accountData.Permissions = []string{"test2"}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}
		appConfigMock.On("GetAuthenticationType").Return(auth.Ldap)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(workspaceData.ToBytes()))
		w := httptest.NewRecorder()

		handler.Create(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestGet(t *testing.T) {
	accountData := &proto.GetAccountDataResponse{
		AccountID:   uuid.New().String(),
		Permissions: []string{"test"},
	}

	t.Run("should return 200 when everything it is ok", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		controllerMock.On("Get").Return(&workspaceEntities.Response{}, nil)

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Get(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		controllerMock.On("Get").Return(&workspaceEntities.Response{}, errors.New("test"))

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Get(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when failed to get account data", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, errors.New("test"))

		appConfigMock := &app.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Get(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when invalid workspace id", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", "test")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Get(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestUpdate(t *testing.T) {
	accountData := &proto.GetAccountDataResponse{
		AccountID:   uuid.New().String(),
		Permissions: []string{"test"},
	}

	workspaceData := &workspaceEntities.Data{
		Name:        "test",
		Description: "test",
		AuthzMember: []string{"test"},
		AuthzAdmin:  []string{"test"},
	}

	t.Run("should return 200 when everything it is ok", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		controllerMock.On("Update").Return(&workspaceEntities.Response{}, nil)

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}
		appConfigMock.On("GetAuthenticationType").Return(auth.Horusec)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(workspaceData.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Update(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		controllerMock.On("Update").Return(&workspaceEntities.Response{}, errors.New("test"))

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}
		appConfigMock.On("GetAuthenticationType").Return(auth.Horusec)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(workspaceData.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Update(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid workspace id", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}
		appConfigMock.On("GetAuthenticationType").Return(auth.Horusec)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(workspaceData.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", "test")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Update(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 failed to get account data", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, errors.New("test"))

		appConfigMock := &app.Mock{}
		appConfigMock.On("GetAuthenticationType").Return(auth.Horusec)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(workspaceData.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Update(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should return 204 when everything it is ok", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		controllerMock.On("Delete").Return(nil)

		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodDelete, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Delete(w, r)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		controllerMock.On("Delete").Return(errors.New("test"))

		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodDelete, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Delete(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid workspace id", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodDelete, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", "test")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Delete(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestList(t *testing.T) {
	accountData := &proto.GetAccountDataResponse{
		AccountID:   uuid.New().String(),
		Permissions: []string{"test"},
	}

	t.Run("should return 200 when everything it is ok", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		controllerMock.On("List").Return(&[]workspaceEntities.Response{}, nil)

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		handler.List(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		controllerMock.On("List").Return(&[]workspaceEntities.Response{}, errors.New("test"))

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		handler.List(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when failed to get account data", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		appConfigMock := &app.Mock{}

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, errors.New("test"))

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		handler.List(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestUpdateRole(t *testing.T) {
	roleData := &role.Data{
		Role: account.Member,
	}

	t.Run("should return 200 when everything it is ok", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		controllerMock.On("UpdateRole").Return(&role.Response{}, nil)

		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(roleData.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("accountID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.UpdateRole(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		controllerMock.On("UpdateRole").Return(&role.Response{}, errors.New("test"))

		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(roleData.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("accountID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.UpdateRole(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 invalid account id", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		appConfigMock := &app.Mock{}
		authGRPCMock := &proto.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(roleData.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("accountID", "test")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.UpdateRole(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when invalid body", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		appConfigMock := &app.Mock{}
		authGRPCMock := &proto.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader([]byte("test")))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("accountID", "test")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.UpdateRole(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestInviteUser(t *testing.T) {
	roleData := &role.UserData{
		Role:      account.Member,
		Email:     "test@test.com",
		AccountID: uuid.New(),
		Username:  "test",
	}

	accountData := &proto.GetAccountDataResponse{
		AccountID:   uuid.New().String(),
		Permissions: []string{"test"},
		Username:    "test",
		Email:       "test@test.com",
	}

	t.Run("should return 200 when everything it is ok", func(t *testing.T) {
		appConfigMock := &app.Mock{}

		controllerMock := &workspaceController.Mock{}
		controllerMock.On("InviteUser").Return(&role.Response{}, nil)

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(roleData.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.InviteUser(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		appConfigMock := &app.Mock{}

		controllerMock := &workspaceController.Mock{}
		controllerMock.On("InviteUser").Return(&role.Response{}, errors.New("test"))

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(roleData.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.InviteUser(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 request body", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader([]byte("")))
		w := httptest.NewRecorder()

		handler.InviteUser(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when failed to get account by email", func(t *testing.T) {
		controllerMock := &workspaceController.Mock{}
		appConfigMock := &app.Mock{}

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, errors.New("test"))

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(roleData.ToBytes()))
		w := httptest.NewRecorder()

		handler.InviteUser(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestGetUsers(t *testing.T) {
	t.Run("should return 200 when everything it is ok", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		controllerMock := &workspaceController.Mock{}
		controllerMock.On("GetUsers").Return(&[]role.Response{}, nil)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.GetUsers(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		controllerMock := &workspaceController.Mock{}
		controllerMock.On("GetUsers").Return(&[]role.Response{}, errors.New("test"))

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.GetUsers(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid workspace id", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}
		controllerMock := &workspaceController.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", "test")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.GetUsers(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestRemoveUser(t *testing.T) {
	t.Run("should return 204 success remove user", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		controllerMock := &workspaceController.Mock{}
		controllerMock.On("RemoveUser").Return(nil)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodDelete, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("accountID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.RemoveUser(w, r)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		controllerMock := &workspaceController.Mock{}
		controllerMock.On("RemoveUser").Return(errors.New("test"))

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodDelete, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("accountID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.RemoveUser(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid account id", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}
		controllerMock := &workspaceController.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodDelete, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("accountID", "test")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.RemoveUser(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when invalid workspace id", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}
		controllerMock := &workspaceController.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodDelete, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", "test")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.RemoveUser(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestCreateToken(t *testing.T) {
	data := &tokenEntities.Data{
		Description: "test",
		IsExpirable: false,
		ExpiresAt:   time.Time{},
	}

	t.Run("should return 201 when everything it is ok", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		controllerMock := &workspaceController.Mock{}
		controllerMock.On("CreateToken").Return(uuid.NewString(), nil)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToByes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.CreateToken(w, r)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		controllerMock := &workspaceController.Mock{}
		controllerMock.On("CreateToken").Return(uuid.NewString(), errors.New("test"))

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToByes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.CreateToken(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid request body", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}
		controllerMock := &workspaceController.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader([]byte("test")))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.CreateToken(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when invalid workspace id", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}
		controllerMock := &workspaceController.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(nil))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", "test")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.CreateToken(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestDeleteToken(t *testing.T) {
	t.Run("should return 204 when everything it is ok", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		controllerMock := &workspaceController.Mock{}
		controllerMock.On("DeleteToken").Return(nil)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodDelete, "test", bytes.NewReader(nil))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("tokenID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.DeleteToken(w, r)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		controllerMock := &workspaceController.Mock{}
		controllerMock.On("DeleteToken").Return(errors.New("test"))

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodDelete, "test", bytes.NewReader(nil))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("tokenID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.DeleteToken(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid token id", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}
		controllerMock := &workspaceController.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodDelete, "test", bytes.NewReader(nil))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("tokenID", "test")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.DeleteToken(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when invalid workspace id", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}
		controllerMock := &workspaceController.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodDelete, "test", bytes.NewReader(nil))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", "test")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.DeleteToken(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestListTokens(t *testing.T) {
	t.Run("should return 200 when everything it is ok", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		controllerMock := &workspaceController.Mock{}
		controllerMock.On("ListTokens").Return(&[]tokenEntities.Response{}, nil)

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", bytes.NewReader(nil))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.ListTokens(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		controllerMock := &workspaceController.Mock{}
		controllerMock.On("ListTokens").Return(&[]tokenEntities.Response{}, errors.New("test"))

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", bytes.NewReader(nil))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.ListTokens(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid workspace id", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}
		controllerMock := &workspaceController.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", bytes.NewReader(nil))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", "test")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.ListTokens(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestOptions(t *testing.T) {
	t.Run("should return 204 when options", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}
		controllerMock := &workspaceController.Mock{}

		handler := NewWorkspaceHandler(controllerMock, workspaceUseCases.NewWorkspaceUseCases(),
			authGRPCMock, appConfigMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodOptions, "test", nil)
		w := httptest.NewRecorder()

		handler.Options(w, r)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}
