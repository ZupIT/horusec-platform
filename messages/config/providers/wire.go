//+build wireinject

package providers

import (
	"github.com/google/wire"

	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	brokerConfig "github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	httpRouter "github.com/ZupIT/horusec-devkit/pkg/services/http/router"

	"github.com/ZupIT/horusec-platform/messages/config/cors"
	emailController "github.com/ZupIT/horusec-platform/messages/internal/controllers/email"
	"github.com/ZupIT/horusec-platform/messages/internal/events/email"
	"github.com/ZupIT/horusec-platform/messages/internal/handlers/health"
	"github.com/ZupIT/horusec-platform/messages/internal/router"
	"github.com/ZupIT/horusec-platform/messages/internal/services/mailer"
)

var devKitProviders = wire.NewSet(
	brokerConfig.NewBrokerConfig,
	broker.NewBroker,
	httpRouter.NewHTTPRouter,
)

var configProviders = wire.NewSet(
	cors.NewCorsConfig,
	router.NewHTTPRouter,
)

var controllerProviders = wire.NewSet(
	emailController.NewEmailController,
)

var handleProviders = wire.NewSet(
	health.NewHealthHandler,
)

var eventProviders = wire.NewSet(
	email.NewEmailEventHandler,
)

var serviceProviders = wire.NewSet(
	mailer.NewMailerService,
)

func Initialize(_ string) (router.IRouter, error) {
	wire.Build(serviceProviders, handleProviders, eventProviders, controllerProviders, configProviders,
		devKitProviders)

	return &router.Router{}, nil
}
