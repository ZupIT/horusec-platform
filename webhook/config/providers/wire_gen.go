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

	"github.com/ZupIT/horusec-platform/webhook/config/cors"
	"github.com/ZupIT/horusec-platform/webhook/internal/controllers/dispatcher"
	webhook2 "github.com/ZupIT/horusec-platform/webhook/internal/controllers/webhook"
	webhook4 "github.com/ZupIT/horusec-platform/webhook/internal/events/webhook"
	"github.com/ZupIT/horusec-platform/webhook/internal/handlers/health"
	webhook3 "github.com/ZupIT/horusec-platform/webhook/internal/handlers/webhook"
	"github.com/ZupIT/horusec-platform/webhook/internal/repositories/webhook"
	"github.com/ZupIT/horusec-platform/webhook/internal/router"
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
	iWebhookRepository := webhook.NewWebhookRepository(connection)
	iWebhookController := webhook2.NewWebhookController(iWebhookRepository)
	webhookHandler := webhook3.NewWebhookHandler(iWebhookController)
	iDispatcherController := dispatcher.NewDispatcherController(iWebhookRepository)
	iEvent := webhook4.NewWebhookEvent(iBroker, iDispatcherController)
	routerIRouter := router.NewHTTPRouter(iRouter, iAuthzMiddleware, handler, webhookHandler, iEvent)
	return routerIRouter, nil
}

// wire.go:

var providers = wire.NewSet(auth.NewAuthGRPCConnection, proto.NewAuthServiceClient, app.NewAppConfig, config2.NewBrokerConfig, broker.NewBroker, config.NewDatabaseConfig, database.NewDatabaseReadAndWrite, cors.NewCorsConfig, router2.NewHTTPRouter, middlewares.NewAuthzMiddleware, webhook.NewWebhookRepository, webhook2.NewWebhookController, dispatcher.NewDispatcherController, webhook4.NewWebhookEvent, health.NewHealthHandler, webhook3.NewWebhookHandler, router.NewHTTPRouter)