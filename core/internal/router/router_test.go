package router

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/http"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"

	"github.com/ZupIT/horusec-platform/core/config/cors"
	"github.com/ZupIT/horusec-platform/core/internal/handlers/workspace"
)

func TestNewHTTPRouter(t *testing.T) {
	t.Run("should success create a new http router and set routes", func(t *testing.T) {
		routerService := http.NewHTTPRouter(cors.NewCorsConfig(), "9999")
		workspaceHandler := &workspace.Handler{}
		middlewareService := middlewares.NewAuthzMiddleware(nil)

		assert.NotPanics(t, func() {
			assert.NotNil(t, NewHTTPRouter(routerService, workspaceHandler, middlewareService))
		})
	})
}
