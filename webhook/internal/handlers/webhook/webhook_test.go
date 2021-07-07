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

package webhook

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	enums2 "github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	enumsParser "github.com/ZupIT/horusec-devkit/pkg/utils/parser/enums"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-platform/webhook/internal/controllers/webhook"
	webhookEntity "github.com/ZupIT/horusec-platform/webhook/internal/entities/webhook"
	"github.com/ZupIT/horusec-platform/webhook/internal/enums"
	useCaseWebhook "github.com/ZupIT/horusec-platform/webhook/internal/usecases/webhook"
)

func TestHandler_Options(t *testing.T) {
	t.Run("Should return no content when call options", func(t *testing.T) {
		controllerMock := &webhook.Mock{}
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("", "/test", nil)
		NewWebhookHandler(controllerMock).Options(w, r)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}

func TestHandler_ListAll(t *testing.T) {
	t.Run("Should return status ok when call ListAll", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("", "/test", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		controllerMock := &webhook.Mock{}
		controllerMock.On("ListAll").Return(&[]webhookEntity.WithRepository{}, nil)
		useCaseMock := &useCaseWebhook.Mock{}
		useCaseMock.On("ExtractWorkspaceIDFromURL").Return(uuid.New(), nil)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		handler.ListAll(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("Should return status bad request when call ListAll and not exists workspace", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("", "/test", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", "")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		controllerMock := &webhook.Mock{}
		controllerMock.On("ListAll").Return(&[]webhookEntity.Webhook{}, nil)
		useCaseMock := &useCaseWebhook.Mock{}
		useCaseMock.On("ExtractWorkspaceIDFromURL").Return(uuid.Nil, enums.ErrorWrongWorkspaceID)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		handler.ListAll(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("Should return status internal server error when call ListAll in controller unexpected error", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("", "/test", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		controllerMock := &webhook.Mock{}
		controllerMock.On("ListAll").Return(&[]webhookEntity.WithRepository{}, errors.New("unexpected error"))
		useCaseMock := &useCaseWebhook.Mock{}
		useCaseMock.On("ExtractWorkspaceIDFromURL").Return(uuid.New(), nil)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		handler.ListAll(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestHandler_Remove(t *testing.T) {
	t.Run("Should return status No Content when call Remove", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("", "/test", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("webhookID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		controllerMock := &webhook.Mock{}
		controllerMock.On("Remove").Return(nil)
		useCaseMock := &useCaseWebhook.Mock{}
		useCaseMock.On("ExtractWebhookIDFromURL").Return(uuid.New(), nil)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		handler.Remove(w, r)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
	t.Run("Should return status bad request when call Remove and not exists webhookID", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("", "/test", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("webhookID", "")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		controllerMock := &webhook.Mock{}
		controllerMock.On("Remove").Return(nil)
		useCaseMock := &useCaseWebhook.Mock{}
		useCaseMock.On("ExtractWebhookIDFromURL").Return(uuid.Nil, enums.ErrorWrongWebhookID)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		handler.Remove(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("Should return status not found when call Remove but not exists webhook", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("", "/test", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("webhookID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		controllerMock := &webhook.Mock{}
		controllerMock.On("Remove").Return(enums2.ErrorNotFoundRecords)
		useCaseMock := &useCaseWebhook.Mock{}
		useCaseMock.On("ExtractWebhookIDFromURL").Return(uuid.New(), nil)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		handler.Remove(w, r)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
	t.Run("Should return no status internal server error when call ListAll in controller unexpected error", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("", "/test", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("webhookID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		controllerMock := &webhook.Mock{}
		controllerMock.On("Remove").Return(errors.New("unexpected error"))
		useCaseMock := &useCaseWebhook.Mock{}
		useCaseMock.On("ExtractWebhookIDFromURL").Return(uuid.New(), nil)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		handler.Remove(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestHandler_Save(t *testing.T) {
	t.Run("Should return status OK when call Save", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("", "/test", nil)
		ctx := chi.NewRouteContext()
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		controllerMock := &webhook.Mock{}
		controllerMock.On("Save").Return(uuid.New(), nil)
		useCaseMock := &useCaseWebhook.Mock{}
		useCaseMock.On("DecodeWebhookFromIoRead").Return(&webhookEntity.Webhook{}, nil)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		handler.Save(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("Should return status bad request when call Save and not exists body", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("", "/test", nil)
		ctx := chi.NewRouteContext()
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		controllerMock := &webhook.Mock{}
		controllerMock.On("Save").Return(uuid.New(), nil)
		useCaseMock := &useCaseWebhook.Mock{}
		useCaseMock.On("DecodeWebhookFromIoRead").Return(&webhookEntity.Webhook{}, enumsParser.ErrorBodyEmpty)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		handler.Save(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("Should return status conflict when call Save but already exists webhook", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("", "/test", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("webhookID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		controllerMock := &webhook.Mock{}
		controllerMock.On("Save").Return(uuid.Nil, enums.ErrorWebhookDuplicate)
		useCaseMock := &useCaseWebhook.Mock{}
		useCaseMock.On("DecodeWebhookFromIoRead").Return(&webhookEntity.Webhook{}, nil)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		handler.Save(w, r)
		assert.Equal(t, http.StatusConflict, w.Code)
	})
	t.Run("Should return no status internal server error when call ListAll in controller unexpected error", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("", "/test", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("webhookID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		controllerMock := &webhook.Mock{}
		controllerMock.On("Save").Return(uuid.Nil, errors.New("unexpected error"))
		useCaseMock := &useCaseWebhook.Mock{}
		useCaseMock.On("DecodeWebhookFromIoRead").Return(&webhookEntity.Webhook{}, nil)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		handler.Save(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestHandler_Update(t *testing.T) {
	t.Run("Should return status no content when call Update", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("", "/test", nil)
		ctx := chi.NewRouteContext()
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		controllerMock := &webhook.Mock{}
		controllerMock.On("Update").Return(nil)
		useCaseMock := &useCaseWebhook.Mock{}
		useCaseMock.On("ExtractWebhookIDFromURL").Return(uuid.New(), nil)
		useCaseMock.On("DecodeWebhookFromIoRead").Return(&webhookEntity.Webhook{}, nil)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		handler.Update(w, r)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
	t.Run("Should return status bad request when call Update and not exists body", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("", "/test", nil)
		ctx := chi.NewRouteContext()
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		controllerMock := &webhook.Mock{}
		controllerMock.On("Update").Return(nil)
		useCaseMock := &useCaseWebhook.Mock{}
		useCaseMock.On("ExtractWebhookIDFromURL").Return(uuid.New(), nil)
		useCaseMock.On("DecodeWebhookFromIoRead").Return(&webhookEntity.Webhook{}, enumsParser.ErrorBodyEmpty)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		handler.Update(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("Should return status bad request when call Update and not webhookID on query", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("", "/test", nil)
		ctx := chi.NewRouteContext()
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		controllerMock := &webhook.Mock{}
		controllerMock.On("Update").Return(nil)
		useCaseMock := &useCaseWebhook.Mock{}
		useCaseMock.On("ExtractWebhookIDFromURL").Return(uuid.Nil, enums.ErrorWrongWebhookID)
		useCaseMock.On("DecodeWebhookFromIoRead").Return(&webhookEntity.Webhook{}, nil)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		handler.Update(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("Should return status not found when call Update but not exists on system", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("", "/test", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("webhookID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		controllerMock := &webhook.Mock{}
		controllerMock.On("Update").Return(enums2.ErrorNotFoundRecords)
		useCaseMock := &useCaseWebhook.Mock{}
		useCaseMock.On("ExtractWebhookIDFromURL").Return(uuid.New(), nil)
		useCaseMock.On("DecodeWebhookFromIoRead").Return(&webhookEntity.Webhook{}, nil)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		handler.Update(w, r)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
	t.Run("Should return no status internal server error when call ListAll in controller unexpected error", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("", "/test", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("webhookID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		controllerMock := &webhook.Mock{}
		controllerMock.On("Update").Return(errors.New("unexpected error"))
		useCaseMock := &useCaseWebhook.Mock{}
		useCaseMock.On("ExtractWebhookIDFromURL").Return(uuid.New(), nil)
		useCaseMock.On("DecodeWebhookFromIoRead").Return(&webhookEntity.Webhook{}, nil)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		handler.Update(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
