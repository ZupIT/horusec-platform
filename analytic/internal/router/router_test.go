package router

import (
	"testing"

	"github.com/go-chi/cors"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/http/router"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"

	eventDashboard "github.com/ZupIT/horusec-platform/analytic/internal/events/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/health"
)

func TestNewHTTPRouter(t *testing.T) {
	t.Run("should add all necessary routes", func(t *testing.T) {
		routerConn := router.NewHTTPRouter(&cors.Options{}, "8009")
		healthMock := &health.Handler{}
		dashboardHandlerMock := &dashboard.Handler{}
		middlewareMock := &middlewares.AuthzMiddleware{}
		eventMock := &eventDashboard.Events{}
		instance := NewHTTPRouter(routerConn, middlewareMock, healthMock, dashboardHandlerMock, eventMock)
		assert.NotEmpty(t, instance)
	})
}
