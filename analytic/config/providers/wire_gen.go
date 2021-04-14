// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package providers

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/services/http"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"
	"github.com/google/wire"

	"github.com/ZupIT/horusec-platform/analytic/config/cors"
	dashboard2 "github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
	dashboardrepository "github.com/ZupIT/horusec-platform/analytic/internal/handlers/dashboard_repository"
	dashboardworkspace "github.com/ZupIT/horusec-platform/analytic/internal/handlers/dashboard_workspace"
	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/health"
	"github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/router"
)

// Injectors from wire.go:

func Initialize(defaultPort string) (router.IRouter, error) {
	options := cors.NewCorsConfig()
	iRouter := http.NewHTTPRouter(options, defaultPort)
	clientConnInterface := auth.NewAuthGRPCConnection()
	iAuthzMiddleware := middlewares.NewAuthzMiddleware(clientConnInterface)
	iConfig := config.NewDatabaseConfig()
	connection, err := database.NewDatabaseReadAndWrite(iConfig)
	if err != nil {
		return nil, err
	}
	handler := health.NewHealthHandler(connection, clientConnInterface)
	iRepoDashboard := dashboard.NewRepoDashboard(connection)
	iController := dashboard2.NewControllerDashboard(iRepoDashboard)
	dashboardworkspaceHandler := dashboardworkspace.NewDashboardWorkspaceHandler(iController)
	dashboardrepositoryHandler := dashboardrepository.NewDashboardRepositoryHandler(iController)
	routerIRouter := router.NewHTTPRouter(iRouter, iAuthzMiddleware, handler, dashboardworkspaceHandler, dashboardrepositoryHandler)
	return routerIRouter, nil
}

// wire.go:

var providers = wire.NewSet(config.NewDatabaseConfig, database.NewDatabaseReadAndWrite, auth.NewAuthGRPCConnection, proto.NewAuthServiceClient, cors.NewCorsConfig, http.NewHTTPRouter, middlewares.NewAuthzMiddleware, dashboard.NewRepoDashboard, dashboard2.NewControllerDashboard, health.NewHealthHandler, dashboardworkspace.NewDashboardWorkspaceHandler, dashboardrepository.NewDashboardRepositoryHandler, router.NewHTTPRouter)
