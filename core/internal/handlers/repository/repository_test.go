package repository

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-platform/core/internal/entities/role"

	repositoryController "github.com/ZupIT/horusec-platform/core/internal/controllers/repository"
	repositoryEntities "github.com/ZupIT/horusec-platform/core/internal/entities/repository"
	repositoryEnums "github.com/ZupIT/horusec-platform/core/internal/enums/repository"
	repositoryUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/repository"
	roleUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/role"
)

func TestCreate(t *testing.T) {
	data := &repositoryEntities.Data{
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
		controllerMock := &repositoryController.Mock{}
		controllerMock.On("Create").Return(&repositoryEntities.Response{}, nil)

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}
		appConfigMock.On("GetAuthorizationType").Return(auth.Horusec)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Create(w, r)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		controllerMock.On("Create").Return(&repositoryEntities.Response{}, errors.New("test"))

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}
		appConfigMock.On("GetAuthorizationType").Return(auth.Horusec)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Create(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when name already in use", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		controllerMock.On("Create").Return(
			&repositoryEntities.Response{}, repositoryEnums.ErrorRepositoryNameAlreadyInUse)

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}
		appConfigMock.On("GetAuthorizationType").Return(auth.Horusec)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Create(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when invalid ldap groups", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}
		appConfigMock.On("GetAuthorizationType").Return(auth.Ldap)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		data.AuthzAdmin = []string{"test2"}
		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Create(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when invalid request body", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		appConfigMock := &app.Mock{}

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader([]byte("")))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Create(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when failed to get account data", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		appConfigMock := &app.Mock{}

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, errors.New("test"))

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader([]byte("")))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Create(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestGet(t *testing.T) {
	accountData := &proto.GetAccountDataResponse{
		AccountID: uuid.New().String(),
	}

	t.Run("should return 200 when everything it is ok", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		controllerMock.On("Get").Return(&repositoryEntities.Response{}, nil)

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Get(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		controllerMock.On("Get").Return(&repositoryEntities.Response{}, errors.New("test"))

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Get(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when failed to get account data", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		appConfigMock := &app.Mock{}

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, errors.New("test"))

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Get(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when invalid repository id", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		appConfigMock := &app.Mock{}
		authGRPCMock := &proto.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("repositoryID", "test")
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

	data := &repositoryEntities.Data{
		Name:            "test",
		Description:     "test",
		AuthzMember:     []string{"test"},
		AuthzAdmin:      []string{"test"},
		AuthzSupervisor: []string{"test"},
	}

	t.Run("should return 200 when everything it is ok", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		controllerMock.On("Update").Return(&repositoryEntities.Response{}, nil)

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}
		appConfigMock.On("GetAuthorizationType").Return(auth.Horusec)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Update(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		controllerMock.On("Update").Return(&repositoryEntities.Response{}, errors.New("test"))

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}
		appConfigMock.On("GetAuthorizationType").Return(auth.Horusec)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Update(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when name already in use", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		controllerMock.On("Update").Return(
			&repositoryEntities.Response{}, repositoryEnums.ErrorRepositoryNameAlreadyInUse)

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}
		appConfigMock.On("GetAuthorizationType").Return(auth.Horusec)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Update(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when invalid ldap groups", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}
		appConfigMock.On("GetAuthorizationType").Return(auth.Ldap)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		data.AuthzAdmin = []string{"test2"}
		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Update(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when failed to get account data", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		appConfigMock := &app.Mock{}

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, errors.New("test"))

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Update(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when failed to get repository id", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", "test")
		ctx.URLParams.Add("repositoryID", "test")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Update(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should return 204 when everything it is ok", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		controllerMock.On("Delete").Return(nil)

		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodDelete, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Delete(w, r)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("should return 500 whe something went wrong", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		controllerMock.On("Delete").Return(errors.New("test"))

		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodDelete, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Delete(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when failed to get repository id", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodDelete, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", "test")
		ctx.URLParams.Add("repositoryID", "test")
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
		controllerMock := &repositoryController.Mock{}
		controllerMock.On("List").Return(&[]repositoryEntities.Response{}, nil)

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.New().String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.List(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		controllerMock.On("List").Return(&[]repositoryEntities.Response{}, errors.New("test"))

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		appConfigMock := &app.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.New().String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.List(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when failed to get account data", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		appConfigMock := &app.Mock{}

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, errors.New("test"))

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.New().String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.List(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestUpdateRole(t *testing.T) {
	roleData := &role.Data{
		Role: account.Member,
	}

	t.Run("should return 200 when everything it is ok", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		controllerMock.On("UpdateRole").Return(&role.Response{}, nil)

		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(roleData.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		ctx.URLParams.Add("accountID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.UpdateRole(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		controllerMock.On("UpdateRole").Return(&role.Response{}, errors.New("test"))

		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(roleData.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		ctx.URLParams.Add("accountID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.UpdateRole(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when user does not belong to the workspace", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		controllerMock.On("UpdateRole").Return(
			&role.Response{}, repositoryEnums.ErrorUserDoesNotBelongToWorkspace)

		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(roleData.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		ctx.URLParams.Add("accountID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.UpdateRole(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when invalid account id", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(roleData.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("accountID", "test")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.UpdateRole(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when invalid request body", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases())

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader([]byte("")))
		w := httptest.NewRecorder()

		handler.UpdateRole(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
