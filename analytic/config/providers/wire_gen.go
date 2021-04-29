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
	"github.com/google/wire"

	"github.com/ZupIT/horusec-platform/analytic/config/cors"
	dashboard2 "github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
	dashboard4 "github.com/ZupIT/horusec-platform/analytic/internal/events/dashboard"
	dashboard3 "github.com/ZupIT/horusec-platform/analytic/internal/handlers/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/health"
	"github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/router"
)

// Injectors from wire.go:

func Initialize(defaultPort string) (router.IRouter, error) {
	options := cors.NewCorsConfig()
	iRouter := router2.NewHTTPRouter(options, defaultPort)
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
	iReadController := dashboard2.NewControllerDashboardRead(iRepoDashboard)
	dashboardHandler := dashboard3.NewDashboardHandler(iReadController)
	iWriteController := dashboard2.NewControllerDashboardWrite(iRepoDashboard)
	iEvent := dashboard4.NewDashboardEvent(iBroker, iWriteController)
	routerIRouter := router.NewHTTPRouter(iRouter, iAuthzMiddleware, handler, dashboardHandler, iEvent)
	return routerIRouter, nil
}

// wire.go:

var providers = wire.NewSet(auth.NewAuthGRPCConnection, proto.NewAuthServiceClient, app.NewAppConfig, config2.NewBrokerConfig, broker.NewBroker, config.NewDatabaseConfig, database.NewDatabaseReadAndWrite, cors.NewCorsConfig, router2.NewHTTPRouter, middlewares.NewAuthzMiddleware, dashboard.NewRepoDashboard, dashboard2.NewControllerDashboardWrite, dashboard2.NewControllerDashboardRead, health.NewHealthHandler, dashboard3.NewDashboardHandler, dashboard4.NewDashboardEvent, router.NewHTTPRouter)
