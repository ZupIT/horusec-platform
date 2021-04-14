package router

import (
	dashboardRepository "github.com/ZupIT/horusec-platform/analytic/internal/handlers/dashboard_repository"
	dashboardWorkspace "github.com/ZupIT/horusec-platform/analytic/internal/handlers/dashboard_workspace"
	"testing"

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
		dashboardWorkspaceMock := &dashboardWorkspace.Handler{}
		dashboardRepositoryMock := &dashboardRepository.Handler{}
		middlewareMock := &middlewares.AuthzMiddleware{}
		instance := NewHTTPRouter(router, middlewareMock, healthMock, dashboardWorkspaceMock, dashboardRepositoryMock)
		assert.NotEmpty(t, instance)
	})
}
