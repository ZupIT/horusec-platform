package router

import (
	"testing"

	eventDashboard "github.com/ZupIT/horusec-platform/analytic/internal/events/dashboard"

	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/dashboard"

	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"

	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/health"

	"github.com/go-chi/cors"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/http"
)

func TestNewHTTPRouter(t *testing.T) {
	t.Run("Should add all necessary routes", func(t *testing.T) {
		router := http.NewHTTPRouter(&cors.Options{}, "8009")
		healthMock := &health.Handler{}
		dashboardHandlerMock := &dashboard.Handler{}
		middlewareMock := &middlewares.AuthzMiddleware{}
		eventMock := &eventDashboard.Event{}
		instance := NewHTTPRouter(router, middlewareMock, healthMock, dashboardHandlerMock, eventMock)
		assert.NotEmpty(t, instance)
	})
}
