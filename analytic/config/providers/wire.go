//+build wireinject

package providers

import (
	"github.com/google/wire"

	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/health"

	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"

	"github.com/ZupIT/horusec-platform/analytic/internal/router"

	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseConfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/http"

	"github.com/ZupIT/horusec-platform/analytic/config/cors"
)

var providers = wire.NewSet(
	databaseConfig.NewDatabaseConfig,
	database.NewDatabaseReadAndWrite,
	auth.NewAuthGRPCConnection,
	proto.NewAuthServiceClient,
	cors.NewCorsConfig,
	http.NewHTTPRouter,
	middlewares.NewAuthzMiddleware,

	health.NewHealthHandler,

	router.NewHTTPRouter,
)

func Initialize(defaultPort string) (router.IRouter, error) {
	wire.Build(providers)
	return &router.Router{}, nil
}
