//+build wireinject

package providers

import (
	"github.com/google/wire"

	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	brokerConfig "github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseConfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/http"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	"github.com/ZupIT/horusec-platform/auth/config/cors"
	"github.com/ZupIT/horusec-platform/auth/config/grpc"
	authController "github.com/ZupIT/horusec-platform/auth/internal/controllers/authentication"
	authHandler "github.com/ZupIT/horusec-platform/auth/internal/handlers/authentication"
	accountRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/account"
	"github.com/ZupIT/horusec-platform/auth/internal/router"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/horusec"
	accountUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/account"
	authUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/authentication"
)

var devKitProviders = wire.NewSet(
	http.NewHTTPRouter,
	databaseConfig.NewDatabaseConfig,
	brokerConfig.NewBrokerConfig,
	broker.NewBroker,
	database.NewDatabaseReadAndWrite,
)

var configProviders = wire.NewSet(
	grpc.NewAuthGRPCServer,
	cors.NewCorsConfig,
	router.NewHTTPRouter,
	app.NewAuthAppConfig,
)

var controllerProviders = wire.NewSet(
	authController.NewAuthenticationController,
)

var handleProviders = wire.NewSet(
	authHandler.NewAuthenticationHandler,
)

var useCasesProviders = wire.NewSet(
	authUseCases.NewAuthenticationUseCases,
	accountUseCases.NewAccountUseCases,
)

var repositoriesProviders = wire.NewSet(
	accountRepository.NewAccountRepository,
)

var serviceProviders = wire.NewSet(
	horusec.NewHorusecAuthenticationService,
)

func Initialize(_ string) (router.IRouter, error) {
	wire.Build(devKitProviders, configProviders, controllerProviders, handleProviders,
		useCasesProviders, repositoriesProviders, serviceProviders)

	return &router.Router{}, nil
}
