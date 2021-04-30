package router

import (
	"testing"

	"github.com/stretchr/testify/assert"

	httpRouter "github.com/ZupIT/horusec-devkit/pkg/services/http/router"

	"github.com/ZupIT/horusec-platform/messages/config/cors"
	"github.com/ZupIT/horusec-platform/messages/internal/handlers/health"
)

func TestNewHTTPRouter(t *testing.T) {
	t.Run("should success create a new http router and set routes", func(t *testing.T) {
		routerService := httpRouter.NewHTTPRouter(cors.NewCorsConfig(), "9999")

		assert.NotPanics(t, func() {
			assert.NotNil(t, NewHTTPRouter(routerService, &health.Handler{}))
		})
	})
}
