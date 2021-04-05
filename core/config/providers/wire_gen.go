// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package providers

import (
	"github.com/ZupIT/horusec-platform/core/config/cors"
	"github.com/ZupIT/horusec-platform/core/internal/controllers/workspace"
	workspace3 "github.com/ZupIT/horusec-platform/core/internal/handlers/workspace"
	"github.com/ZupIT/horusec-platform/core/internal/router"
	workspace2 "github.com/ZupIT/horusec-platform/core/internal/usecases/workspace"
	"github.com/google/wire"

	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	config2 "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/services/http"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"
)

// Injectors from wire.go:

func Initialize(defaultPort string) (router.IRouter, error) {
	options := cors.NewCorsConfig()
	iRouter := http.NewHTTPRouter(options, defaultPort)
	iConfig := config.NewBrokerConfig()
	clientConnInterface := auth.NewAuthGRPCConnection()
	authServiceClient := proto.NewAuthServiceClient(clientConnInterface)
	appIConfig := app.NewAppConfig(authServiceClient)
	iBroker, err := broker.NewBroker(iConfig, appIConfig)
	if err != nil {
		return nil, err
	}
	configIConfig := config2.NewDatabaseConfig()
	connection, err := database.NewDatabaseReadAndWrite(configIConfig)
	if err != nil {
		return nil, err
	}
	iController := workspace.NewWorkspaceController(iBroker, connection, appIConfig)
	iUseCases := workspace2.NewWorkspaceUseCases()
	handler := workspace3.NewWorkspaceHandler(iController, iUseCases, authServiceClient, appIConfig)
	iAuthzMiddleware := middlewares.NewAuthzMiddleware(clientConnInterface)
	routerIRouter := router.NewHTTPRouter(iRouter, handler, iAuthzMiddleware)
	return routerIRouter, nil
}

// wire.go:

var devKitProviders = wire.NewSet(config.NewBrokerConfig, broker.NewBroker, config2.NewDatabaseConfig, database.NewDatabaseReadAndWrite, http.NewHTTPRouter, auth.NewAuthGRPCConnection, proto.NewAuthServiceClient, app.NewAppConfig, middlewares.NewAuthzMiddleware)

var configProviders = wire.NewSet(cors.NewCorsConfig, router.NewHTTPRouter)

var controllerProviders = wire.NewSet(workspace.NewWorkspaceController)

var handleProviders = wire.NewSet(workspace3.NewWorkspaceHandler)

var useCasesProviders = wire.NewSet(workspace2.NewWorkspaceUseCases)

var repositoriesProviders = wire.NewSet()
