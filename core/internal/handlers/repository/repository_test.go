// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package repository

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

	repositoryController "github.com/ZupIT/horusec-platform/core/internal/controllers/repository"
	repositoryEntities "github.com/ZupIT/horusec-platform/core/internal/entities/repository"
	"github.com/ZupIT/horusec-platform/core/internal/entities/role"
	tokenEntities "github.com/ZupIT/horusec-platform/core/internal/entities/token"
	repositoryEnums "github.com/ZupIT/horusec-platform/core/internal/enums/repository"
	repositoryUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/repository"
	roleUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/role"
	tokenUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/token"
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
		appConfigMock.On("GetAuthenticationType").Return(auth.Horusec)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
		appConfigMock.On("GetAuthenticationType").Return(auth.Horusec)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
		appConfigMock.On("GetAuthenticationType").Return(auth.Horusec)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
		appConfigMock.On("GetAuthenticationType").Return(auth.Ldap)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
		appConfigMock.On("GetAuthenticationType").Return(auth.Horusec)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
		appConfigMock.On("GetAuthenticationType").Return(auth.Horusec)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
		appConfigMock.On("GetAuthenticationType").Return(auth.Horusec)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
		appConfigMock.On("GetAuthenticationType").Return(auth.Ldap)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

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
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader([]byte("")))
		w := httptest.NewRecorder()

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

	t.Run("should return 200 when everything it is ok", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		controllerMock.On("InviteUser").Return(&role.Response{}, nil)

		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(roleData.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.InviteUser(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		controllerMock.On("InviteUser").Return(&role.Response{}, errors.New("test"))

		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(roleData.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.InviteUser(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when user does not belong to workspace", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		controllerMock.On("InviteUser").Return(
			&role.Response{}, repositoryEnums.ErrorUserDoesNotBelongToWorkspace)

		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(roleData.ToBytes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.InviteUser(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when invalid request body", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		controllerMock.On("InviteUser").Return(
			&role.Response{}, repositoryEnums.ErrorUserDoesNotBelongToWorkspace)

		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader([]byte("test")))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.InviteUser(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestGetUsers(t *testing.T) {
	t.Run("should return 200 when everything it is ok", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		controllerMock := &repositoryController.Mock{}
		controllerMock.On("GetUsers").Return(&[]role.Response{}, nil)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.GetUsers(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		controllerMock := &repositoryController.Mock{}
		controllerMock.On("GetUsers").Return(&[]role.Response{}, errors.New("test"))

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.GetUsers(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid repository id", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		controllerMock := &repositoryController.Mock{}
		controllerMock.On("GetUsers").Return(&[]role.Response{}, errors.New("test"))

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", "test")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.GetUsers(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestRemoveUser(t *testing.T) {
	t.Run("should return 204 success remove user", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		controllerMock := &repositoryController.Mock{}
		controllerMock.On("RemoveUser").Return(nil)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodDelete, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		ctx.URLParams.Add("accountID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.RemoveUser(w, r)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		controllerMock := &repositoryController.Mock{}
		controllerMock.On("RemoveUser").Return(errors.New("test"))

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodDelete, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		ctx.URLParams.Add("accountID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.RemoveUser(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid account id", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}
		controllerMock := &repositoryController.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodDelete, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		ctx.URLParams.Add("accountID", "test")
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

		controllerMock := &repositoryController.Mock{}
		controllerMock.On("CreateToken").Return(uuid.NewString(), nil)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToByes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.CreateToken(w, r)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		controllerMock := &repositoryController.Mock{}
		controllerMock.On("CreateToken").Return(uuid.NewString(), errors.New("test"))

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToByes()))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.CreateToken(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid request body", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}
		controllerMock := &repositoryController.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader([]byte("")))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.CreateToken(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestDeleteToken(t *testing.T) {
	t.Run("should return 204 when everything it is ok", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		controllerMock := &repositoryController.Mock{}
		controllerMock.On("DeleteToken").Return(nil)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodDelete, "test", bytes.NewReader(nil))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("tokenID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.DeleteToken(w, r)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		controllerMock := &repositoryController.Mock{}
		controllerMock.On("DeleteToken").Return(errors.New("test"))

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodDelete, "test", bytes.NewReader(nil))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("tokenID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.DeleteToken(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid token id", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}
		controllerMock := &repositoryController.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodDelete, "test", bytes.NewReader(nil))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("tokenID", "test")
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.DeleteToken(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestListTokens(t *testing.T) {
	t.Run("should return 200 when everything it is ok", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		controllerMock := &repositoryController.Mock{}
		controllerMock.On("ListTokens").Return(&[]tokenEntities.Response{}, nil)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", bytes.NewReader(nil))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.ListTokens(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}

		controllerMock := &repositoryController.Mock{}
		controllerMock.On("ListTokens").Return(&[]tokenEntities.Response{}, errors.New("test"))

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", bytes.NewReader(nil))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.ListTokens(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid repository id", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}
		controllerMock := &repositoryController.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", bytes.NewReader(nil))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		ctx.URLParams.Add("repositoryID", "test")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.ListTokens(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when invalid workspace id", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}
		controllerMock := &repositoryController.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodGet, "test", bytes.NewReader(nil))
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", "test")
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.ListTokens(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestOptions(t *testing.T) {
	t.Run("should return 204 when options", func(t *testing.T) {
		authGRPCMock := &proto.Mock{}
		appConfigMock := &app.Mock{}
		controllerMock := &repositoryController.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock, roleUseCases.NewRoleUseCases(), tokenUseCases.NewTokenUseCases())

		r, _ := http.NewRequest(http.MethodOptions, "test", nil)
		w := httptest.NewRecorder()

		handler.Options(w, r)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}
