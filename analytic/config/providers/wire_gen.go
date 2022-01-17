// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package providers

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	config2 "github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	router2 "github.com/ZupIT/horusec-devkit/pkg/services/http/router"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"
	"github.com/google/wire"

	"github.com/ZupIT/horusec-platform/analytic/config/cors"
	dashboard3 "github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
	dashboard5 "github.com/ZupIT/horusec-platform/analytic/internal/events/dashboard"
	dashboard4 "github.com/ZupIT/horusec-platform/analytic/internal/handlers/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/health"
	"github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/router"
	dashboard2 "github.com/ZupIT/horusec-platform/analytic/internal/usecases/dashboard"
)

// Injectors from wire.go:

func Initialize(string2 string) (router.IRouter, error) {
	options := cors.NewCorsConfig()
	iRouter := router2.NewHTTPRouter(options, string2)
	clientConnInterface := auth.NewAuthGRPCConnection()
	iAuthzMiddleware := middlewares.NewAuthzMiddleware(clientConnInterface)
	iConfig := config.NewDatabaseConfig()
	connection, err := database.NewDatabaseReadAndWrite(iConfig)
	if err != nil {
		return nil, err
	}
	configIConfig := config2.NewBrokerConfig()
	iBroker, err := broker.NewBroker(configIConfig)
	if err != nil {
		return nil, err
	}
	handler := health.NewHealthHandler(connection, iBroker)
	iRepoRepository := dashboard.NewRepoDashboard(connection)
	iWorkspaceRepository := dashboard.NewWorkspaceDashboard(connection)
	iUseCases := dashboard2.NewUseCaseDashboard()
	iController := dashboard3.NewDashboardController(iRepoRepository, iWorkspaceRepository, connection, iUseCases)
	dashboardHandler := dashboard4.NewDashboardHandler(iController)
	events := dashboard5.NewDashboardEvents(iBroker, iController)
	routerIRouter := router.NewHTTPRouter(iRouter, iAuthzMiddleware, handler, dashboardHandler, events)
	return routerIRouter, nil
}

// wire.go:

var devKitProviders = wire.NewSet(auth.NewAuthGRPCConnection, proto.NewAuthServiceClient, app.NewAppConfig, config2.NewBrokerConfig, broker.NewBroker, config.NewDatabaseConfig, database.NewDatabaseReadAndWrite, router2.NewHTTPRouter, middlewares.NewAuthzMiddleware)

var configProviders = wire.NewSet(cors.NewCorsConfig, router.NewHTTPRouter)

var repositoriesProviders = wire.NewSet(dashboard.NewRepoDashboard, dashboard.NewWorkspaceDashboard)

var controllersProviders = wire.NewSet(dashboard3.NewDashboardController)

var handlersProviders = wire.NewSet(health.NewHealthHandler, dashboard4.NewDashboardHandler)

var eventsProviders = wire.NewSet(dashboard5.NewDashboardEvents)

var useCasesProviders = wire.NewSet(dashboard2.NewUseCaseDashboard)
