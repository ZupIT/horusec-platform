package webhook

import (
	"context"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"
	"github.com/ZupIT/horusec-platform/webhook/internal/entities/webhook"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestUseCaseWebhook_DecodeWebhookFromIoRead(t *testing.T) {
	t.Run("Should decode webhook code without error", func(t *testing.T) {
		wh := &webhook.Webhook{
			URL:          "http://google.com",
			Method:       "POST",
			Headers:      []webhook.Headers{
				{Key: "x-authorization", Value: "1243567890"},
			},
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
		}
		body, err := parser.ParseEntityToIOReadCloser(wh)
		assert.NoError(t, err)
		r, _ := http.NewRequest(http.MethodPost, "/test", body)
		uc := NewUseCaseWebhook()
		entity, err := uc.DecodeWebhookFromIoRead(r)
		assert.NoError(t, err)
		assert.NotEmpty(t, entity)
	})
	t.Run("Should decode webhook code with error invalid method type", func(t *testing.T) {
		wh := &webhook.Webhook{
			URL:          "http://google.com",
			Method:       "GET",
			Headers:      []webhook.Headers{
				{Key: "x-authorization", Value: "1243567890"},
			},
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
		}
		body, err := parser.ParseEntityToIOReadCloser(wh)
		assert.NoError(t, err)
		r, _ := http.NewRequest(http.MethodPost, "/test", body)
		uc := NewUseCaseWebhook()
		entity, err := uc.DecodeWebhookFromIoRead(r)
		assert.NoError(t, err)
		assert.NotEmpty(t, entity)
	})
	t.Run("Should decode body empty and return nil value", func(t *testing.T) {
		r, _ := http.NewRequest(http.MethodPost, "/test", ioutil.NopCloser(strings.NewReader(string("some wrong type"))))
		uc := NewUseCaseWebhook()
		entity, err := uc.DecodeWebhookFromIoRead(r)
		assert.Error(t, err)
		assert.Empty(t, entity)
	})
}

func TestUseCaseWebhook_ExtractWebhookIDFromURL(t *testing.T) {
	t.Run("Should get webhookID from url param without error", func(t *testing.T) {
		r, _ := http.NewRequest(http.MethodPost, "/test", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("webhookID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		uc := NewUseCaseWebhook()
		entity, err := uc.ExtractWebhookIDFromURL(r)
		assert.NoError(t, err)
		assert.NotEmpty(t, entity)
	})
	t.Run("Should get webhookID from url param with error", func(t *testing.T) {
		r, _ := http.NewRequest(http.MethodPost, "/test", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("webhookID", "wrong data type")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		uc := NewUseCaseWebhook()
		entity, err := uc.ExtractWebhookIDFromURL(r)
		assert.Error(t, err)
		assert.Equal(t, entity, uuid.Nil)
	})
}

func TestUseCaseWebhook_ExtractWorkspaceIDFromURL(t *testing.T) {
	t.Run("Should get workspaceID from url param without error", func(t *testing.T) {
		r, _ := http.NewRequest(http.MethodPost, "/test", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		uc := NewUseCaseWebhook()
		entity, err := uc.ExtractWorkspaceIDFromURL(r)
		assert.NoError(t, err)
		assert.NotEmpty(t, entity)
	})
	t.Run("Should get workspaceID from url param with error", func(t *testing.T) {
		r, _ := http.NewRequest(http.MethodPost, "/test", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", "wrong data type")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		uc := NewUseCaseWebhook()
		entity, err := uc.ExtractWorkspaceIDFromURL(r)
		assert.Error(t, err)
		assert.Equal(t, entity, uuid.Nil)
	})
}
