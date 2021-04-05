// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package providers

import (
	"github.com/ZupIT/horusec-platform/api/config/cors"
	"github.com/ZupIT/horusec-platform/api/internal/controllers/analysis"
	analysis2 "github.com/ZupIT/horusec-platform/api/internal/handlers/analysis"
	"github.com/ZupIT/horusec-platform/api/internal/handlers/health"
	"github.com/ZupIT/horusec-platform/api/internal/router"
	"github.com/google/wire"

	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	config2 "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/services/http"
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
	iController := analysis.NewAnalysisController(iBroker, iConfig, connection, appIConfig)
	handler := analysis2.NewAnalysisHandler(iController)
	healthHandler := health.NewHealthHandler(iBroker, iConfig, connection, clientConnInterface, appIConfig)
	routerIRouter := router.NewHTTPRouter(iRouter, handler, healthHandler)
	return routerIRouter, nil
}

// wire.go:

var providers = wire.NewSet(config.NewBrokerConfig, broker.NewBroker, config2.NewDatabaseConfig, database.NewDatabaseReadAndWrite, auth.NewAuthGRPCConnection, proto.NewAuthServiceClient, cors.NewCorsConfig, http.NewHTTPRouter, app.NewAppConfig, analysis.NewAnalysisController, analysis2.NewAnalysisHandler, health.NewHealthHandler, router.NewHTTPRouter)
