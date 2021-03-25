//+build wireinject

package providers

import (
	workspaceUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/workspace"
	"github.com/google/wire"

	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	brokerConfig "github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseConfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/services/http"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"

	"github.com/ZupIT/horusec-platform/core/config/cors"
	workspaceController "github.com/ZupIT/horusec-platform/core/internal/controllers/workspace"
	workspaceHandler "github.com/ZupIT/horusec-platform/core/internal/handlers/workspace"
	"github.com/ZupIT/horusec-platform/core/internal/router"
)

var devKitProviders = wire.NewSet(
	brokerConfig.NewBrokerConfig,
	broker.NewBroker,
	databaseConfig.NewDatabaseConfig,
	database.NewDatabaseReadAndWrite,
	http.NewHTTPRouter,
	auth.NewAuthGRPCConnection,
	proto.NewAuthServiceClient,
	app.NewAppConfig,
	middlewares.NewAuthzMiddleware,
)

var configProviders = wire.NewSet(
	cors.NewCorsConfig,
	router.NewHTTPRouter,
)

var controllerProviders = wire.NewSet(
	workspaceController.NewWorkspaceController,
)

var handleProviders = wire.NewSet(
	workspaceHandler.NewWorkspaceHandler,
)

var useCasesProviders = wire.NewSet(
	workspaceUseCases.NewWorkspaceUseCases,
)

var repositoriesProviders = wire.NewSet()

func Initialize(defaultPort string) (router.IRouter, error) {
	wire.Build(devKitProviders, configProviders, controllerProviders, handleProviders, useCasesProviders)
	return &router.Router{}, nil
}
