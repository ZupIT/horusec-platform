//+build wireinject

package providers

import (
	dashboardfilter "github.com/ZupIT/horusec-platform/analytic/internal/usecase/dashboard"
	"github.com/google/wire"

	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseConfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	routerHttp "github.com/ZupIT/horusec-devkit/pkg/services/http/router"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"

	"github.com/ZupIT/horusec-platform/analytic/config/cors"
	controllerDashboard "github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
	dashboardEvent "github.com/ZupIT/horusec-platform/analytic/internal/events/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/health"
	dashboardRepository "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/router"
)

var devKitProviders = wire.NewSet(
	auth.NewAuthGRPCConnection,
	proto.NewAuthServiceClient,
	app.NewAppConfig,
	config.NewBrokerConfig,
	broker.NewBroker,
	databaseConfig.NewDatabaseConfig,
	database.NewDatabaseReadAndWrite,
	routerHttp.NewHTTPRouter,
	middlewares.NewAuthzMiddleware,
)

var configProviders = wire.NewSet(
	cors.NewCorsConfig,
	router.NewHTTPRouter,
)

var repositoriesProviders = wire.NewSet(
	dashboardRepository.NewRepoDashboard,
)

var controllersProviders = wire.NewSet(
	controllerDashboard.NewControllerDashboardRead,
)

var handlersProviders = wire.NewSet(
	health.NewHealthHandler,
	dashboard.NewDashboardHandler,
)

var eventsProviders = wire.NewSet(
	dashboardEvent.NewDashboardEvents,
)

var useCasesProviders = wire.NewSet(
	dashboardfilter.NewUseCaseDashboard,
)

func Initialize(_ string) (router.IRouter, error) {
	wire.Build(devKitProviders, configProviders, repositoriesProviders, controllersProviders,
		handlersProviders, eventsProviders, useCasesProviders)

	return &router.Router{}, nil
}
