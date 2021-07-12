// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
