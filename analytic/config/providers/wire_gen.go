// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

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
	"github.com/ZupIT/horusec-platform/analytic/config/cors"
	dashboard2 "github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
	dashboard4 "github.com/ZupIT/horusec-platform/analytic/internal/events/dashboard"
	dashboard3 "github.com/ZupIT/horusec-platform/analytic/internal/handlers/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/health"
	"github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/router"
	"github.com/ZupIT/horusec-platform/analytic/internal/usecase/dashboard"
	"github.com/google/wire"
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
	handler := health.NewHealthHandler(connection, clientConnInterface, iBroker)
	iRepoDashboard := dashboard.NewRepoDashboard(connection)
	iUseCases := dashboardfilter.NewUseCaseDashboard()
	iController := dashboard2.NewControllerDashboardRead(iRepoDashboard, connection, iUseCases)
	dashboardHandler := dashboard3.NewDashboardHandler(iController)
	iEvents := dashboard4.NewDashboardEvents(iBroker, iController)
	routerIRouter := router.NewHTTPRouter(iRouter, iAuthzMiddleware, handler, dashboardHandler, iEvents)
	return routerIRouter, nil
}

// wire.go:

var devKitProviders = wire.NewSet(auth.NewAuthGRPCConnection, proto.NewAuthServiceClient, app.NewAppConfig, config2.NewBrokerConfig, broker.NewBroker, config.NewDatabaseConfig, database.NewDatabaseReadAndWrite, router2.NewHTTPRouter, middlewares.NewAuthzMiddleware)

var configProviders = wire.NewSet(cors.NewCorsConfig, router.NewHTTPRouter)

var repositoriesProviders = wire.NewSet(dashboard.NewRepoDashboard)

var controllersProviders = wire.NewSet(dashboard2.NewControllerDashboardRead)

var handlersProviders = wire.NewSet(health.NewHealthHandler, dashboard3.NewDashboardHandler)

var eventsProviders = wire.NewSet(dashboard4.NewDashboardEvents)

var useCasesProviders = wire.NewSet(dashboardfilter.NewUseCaseDashboard)
