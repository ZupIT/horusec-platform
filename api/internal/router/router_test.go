package router

import (
	"testing"

	"github.com/ZupIT/horusec-devkit/pkg/services/http/router"

	"github.com/go-chi/cors"
	"github.com/stretchr/testify/assert"

	analysisHandler "github.com/ZupIT/horusec-platform/api/internal/handlers/analysis"
	healthHandler "github.com/ZupIT/horusec-platform/api/internal/handlers/health"
	"github.com/ZupIT/horusec-platform/api/internal/middelwares/token"
)

func TestNewHTTPRouter(t *testing.T) {
	t.Run("Should add all necessary routes", func(t *testing.T) {
		routerConnection := router.NewHTTPRouter(&cors.Options{}, "8000")
		healthMock := &healthHandler.Handler{}
		analysisMock := &analysisHandler.Handler{}
		tokenMiddlewareMock := token.NewTokenAuthz(nil)
		instance := NewHTTPRouter(routerConnection, tokenMiddlewareMock, analysisMock, healthMock)
		assert.NotEmpty(t, instance)
	})
}
