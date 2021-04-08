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

	repositoryController "github.com/ZupIT/horusec-platform/core/internal/controllers/repository"
	repositoryEntities "github.com/ZupIT/horusec-platform/core/internal/entities/repository"
	repositoryEnums "github.com/ZupIT/horusec-platform/core/internal/enums/repository"
	repositoryUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/repository"
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
			appConfigMock, authGRPCMock)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

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
			appConfigMock, authGRPCMock)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

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
			appConfigMock, authGRPCMock)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

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
			appConfigMock, authGRPCMock)

		data.AuthzAdmin = []string{"test2"}
		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.Create(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when invalid request body", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		appConfigMock := &app.Mock{}

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, nil)

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader([]byte("")))
		w := httptest.NewRecorder()

		handler.Create(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when failed to get account data", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		appConfigMock := &app.Mock{}

		authGRPCMock := &proto.Mock{}
		authGRPCMock.On("GetAccountInfo").Return(accountData, errors.New("test"))

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader([]byte("")))
		w := httptest.NewRecorder()

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
			appConfigMock, authGRPCMock)

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("repositoryID", uuid.NewString())
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
			appConfigMock, authGRPCMock)

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("repositoryID", uuid.NewString())
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
			appConfigMock, authGRPCMock)

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("repositoryID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Get(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when invalid repository id", func(t *testing.T) {
		controllerMock := &repositoryController.Mock{}
		appConfigMock := &app.Mock{}
		authGRPCMock := &proto.Mock{}

		handler := NewRepositoryHandler(repositoryUseCases.NewRepositoryUseCases(), controllerMock,
			appConfigMock, authGRPCMock)

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("repositoryID", "test")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Get(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
