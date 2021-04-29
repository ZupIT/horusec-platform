package router

import (
	"testing"

	"github.com/stretchr/testify/assert"

	httpRouter "github.com/ZupIT/horusec-devkit/pkg/services/http/router"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"

	"github.com/ZupIT/horusec-platform/core/config/cors"
	"github.com/ZupIT/horusec-platform/core/internal/handlers/health"
	"github.com/ZupIT/horusec-platform/core/internal/handlers/repository"
	"github.com/ZupIT/horusec-platform/core/internal/handlers/workspace"
)

func TestNewHTTPRouter(t *testing.T) {
	t.Run("should success create a new http router and set routes", func(t *testing.T) {
		routerService := httpRouter.NewHTTPRouter(cors.NewCorsConfig(), "9999")
		middlewareService := middlewares.NewAuthzMiddleware(nil)
		workspaceHandler := &workspace.Handler{}
		repositoryHandler := &repository.Handler{}
		healthHandler := &health.Handler{}

		assert.NotPanics(t, func() {
			assert.NotNil(t, NewHTTPRouter(routerService, middlewareService, workspaceHandler,
				repositoryHandler, healthHandler))
		})
	})
}
