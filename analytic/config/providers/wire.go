//+build wireinject

package providers

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	routerHttp "github.com/ZupIT/horusec-devkit/pkg/services/http/router"
	"github.com/google/wire"

	dashboardEvent "github.com/ZupIT/horusec-platform/analytic/internal/events/dashboard"

	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/dashboard"

	controllerDashboard "github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
	repoDashboard "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"

	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/health"

	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"

	"github.com/ZupIT/horusec-platform/analytic/internal/router"

	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"

	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseConfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth"

	"github.com/ZupIT/horusec-platform/analytic/config/cors"
)

var providers = wire.NewSet(
	auth.NewAuthGRPCConnection,
	proto.NewAuthServiceClient,
	app.NewAppConfig,

	config.NewBrokerConfig,
	broker.NewBroker,

	databaseConfig.NewDatabaseConfig,
	database.NewDatabaseReadAndWrite,

	cors.NewCorsConfig,
	routerHttp.NewHTTPRouter,

	middlewares.NewAuthzMiddleware,

	repoDashboard.NewRepoDashboard,

	controllerDashboard.NewControllerDashboardWrite,
	controllerDashboard.NewControllerDashboardRead,

	health.NewHealthHandler,
	dashboard.NewDashboardHandler,

	dashboardEvent.NewDashboardEvent,

	router.NewHTTPRouter,
)

func Initialize(defaultPort string) (router.IRouter, error) {
	wire.Build(providers)
	return &router.Router{}, nil
}
