package router

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/http"

	"github.com/ZupIT/horusec-platform/auth/config/cors"
	"github.com/ZupIT/horusec-platform/auth/config/grpc"
	authHandler "github.com/ZupIT/horusec-platform/auth/internal/handlers/authentication"
)

func TestNewHTTPRouter(t *testing.T) {
	t.Run("should success create a new http router and set routes", func(t *testing.T) {
		routerService := http.NewHTTPRouter(cors.NewCorsConfig(), "9999")

		assert.NotPanics(t, func() {
			assert.NotNil(t, NewHTTPRouter(routerService, &grpc.AuthGRPCServer{}, &authHandler.Handler{}))
		})
	})
}
