package router

import (
	"testing"

	webhook2 "github.com/ZupIT/horusec-platform/webhook/internal/events/webhook"
	"github.com/ZupIT/horusec-platform/webhook/internal/handlers/webhook"

	"github.com/ZupIT/horusec-platform/webhook/internal/handlers/health"

	"github.com/ZupIT/horusec-devkit/pkg/services/http/router"

	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"

	"github.com/go-chi/cors"
	"github.com/stretchr/testify/assert"
)

func TestNewHTTPRouter(t *testing.T) {
	t.Run("Should add all necessary routes", func(t *testing.T) {
		routerConn := router.NewHTTPRouter(&cors.Options{}, "8005")
		middlewareMock := &middlewares.AuthzMiddleware{}
		healthMock := &health.Handler{}
		webhookHandlerMock := &webhook.Handler{}
		webhookEventMock := &webhook2.Event{}
		instance := NewHTTPRouter(routerConn, middlewareMock, healthMock, webhookHandlerMock, webhookEventMock)
		assert.NotEmpty(t, instance)
	})
}
