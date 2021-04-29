// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package providers

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	config2 "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	router2 "github.com/ZupIT/horusec-devkit/pkg/services/http/router"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"
	"github.com/google/wire"

	"github.com/ZupIT/horusec-platform/core/config/cors"
	repository3 "github.com/ZupIT/horusec-platform/core/internal/controllers/repository"
	workspace3 "github.com/ZupIT/horusec-platform/core/internal/controllers/workspace"
	"github.com/ZupIT/horusec-platform/core/internal/handlers/health"
	repository4 "github.com/ZupIT/horusec-platform/core/internal/handlers/repository"
	workspace4 "github.com/ZupIT/horusec-platform/core/internal/handlers/workspace"
	repository2 "github.com/ZupIT/horusec-platform/core/internal/repositories/repository"
	workspace2 "github.com/ZupIT/horusec-platform/core/internal/repositories/workspace"
	"github.com/ZupIT/horusec-platform/core/internal/router"
	"github.com/ZupIT/horusec-platform/core/internal/usecases/repository"
	"github.com/ZupIT/horusec-platform/core/internal/usecases/role"
	"github.com/ZupIT/horusec-platform/core/internal/usecases/token"
	"github.com/ZupIT/horusec-platform/core/internal/usecases/workspace"
)

// Injectors from wire.go:

func Initialize(string2 string) (router.IRouter, error) {
	options := cors.NewCorsConfig()
	iRouter := router2.NewHTTPRouter(options, string2)
	clientConnInterface := auth.NewAuthGRPCConnection()
	iAuthzMiddleware := middlewares.NewAuthzMiddleware(clientConnInterface)
	iConfig := config.NewBrokerConfig()
	iBroker, err := broker.NewBroker(iConfig)
	if err != nil {
		return nil, err
	}
	configIConfig := config2.NewDatabaseConfig()
	connection, err := database.NewDatabaseReadAndWrite(configIConfig)
	if err != nil {
		return nil, err
	}
	authServiceClient := proto.NewAuthServiceClient(clientConnInterface)
	appIConfig := app.NewAppConfig(authServiceClient)
	iUseCases := workspace.NewWorkspaceUseCases()
	iRepository := workspace2.NewWorkspaceRepository(connection, iUseCases)
	tokenIUseCases := token.NewTokenUseCases()
	iController := workspace3.NewWorkspaceController(iBroker, connection, appIConfig, iUseCases, iRepository, tokenIUseCases)
	roleIUseCases := role.NewRoleUseCases()
	handler := workspace4.NewWorkspaceHandler(iController, iUseCases, authServiceClient, appIConfig, roleIUseCases, tokenIUseCases)
	repositoryIUseCases := repository.NewRepositoryUseCases()
	repositoryIRepository := repository2.NewRepositoryRepository(connection, repositoryIUseCases, iRepository)
	repositoryIController := repository3.NewRepositoryController(iBroker, connection, appIConfig, repositoryIUseCases, repositoryIRepository, tokenIUseCases)
	repositoryHandler := repository4.NewRepositoryHandler(repositoryIUseCases, repositoryIController, appIConfig, authServiceClient, roleIUseCases, tokenIUseCases)
	healthHandler := health.NewHealthHandler(connection, iBroker)
	routerIRouter := router.NewHTTPRouter(iRouter, iAuthzMiddleware, handler, repositoryHandler, healthHandler)
	return routerIRouter, nil
}

// wire.go:

var devKitProviders = wire.NewSet(config.NewBrokerConfig, broker.NewBroker, config2.NewDatabaseConfig, database.NewDatabaseReadAndWrite, router2.NewHTTPRouter, auth.NewAuthGRPCConnection, proto.NewAuthServiceClient, app.NewAppConfig, middlewares.NewAuthzMiddleware)

var configProviders = wire.NewSet(cors.NewCorsConfig, router.NewHTTPRouter)

var controllerProviders = wire.NewSet(workspace3.NewWorkspaceController, repository3.NewRepositoryController)

var handleProviders = wire.NewSet(workspace4.NewWorkspaceHandler, repository4.NewRepositoryHandler, health.NewHealthHandler)

var useCasesProviders = wire.NewSet(workspace.NewWorkspaceUseCases, repository.NewRepositoryUseCases, role.NewRoleUseCases, token.NewTokenUseCases)

var repositoriesProviders = wire.NewSet(workspace2.NewWorkspaceRepository, repository2.NewRepositoryRepository)
