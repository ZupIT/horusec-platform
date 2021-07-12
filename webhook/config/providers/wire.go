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
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	routerHttp "github.com/ZupIT/horusec-devkit/pkg/services/http/router"
	"github.com/google/wire"

	"github.com/ZupIT/horusec-platform/webhook/internal/controllers/dispatcher"
	webhookController "github.com/ZupIT/horusec-platform/webhook/internal/controllers/webhook"
	webhookEvent "github.com/ZupIT/horusec-platform/webhook/internal/events/webhook"
	"github.com/ZupIT/horusec-platform/webhook/internal/handlers/webhook"
	webhookRepository "github.com/ZupIT/horusec-platform/webhook/internal/repositories/webhook"

	"github.com/ZupIT/horusec-platform/webhook/internal/handlers/health"
	"github.com/ZupIT/horusec-platform/webhook/internal/router"

	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"

	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"

	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseConfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth"

	"github.com/ZupIT/horusec-platform/webhook/config/cors"
)

var providers = wire.NewSet(
	auth.NewAuthGRPCConnection,
	proto.NewAuthServiceClient,
	app.NewAppConfig,

	config.NewBrokerConfig,
	broker.NewBroker,

	databaseConfig.NewDatabaseConfig,
	database.NewDatabaseReadAndWrite,

	cors.NewCorsConfig,
	routerHttp.NewHTTPRouter,

	middlewares.NewAuthzMiddleware,

	webhookRepository.NewWebhookRepository,

	webhookController.NewWebhookController,
	dispatcher.NewDispatcherController,

	webhookEvent.NewWebhookEvent,

	health.NewHealthHandler,
	webhook.NewWebhookHandler,

	router.NewHTTPRouter,
)

func Initialize(defaultPort string) (router.IRouter, error) {
	wire.Build(providers)
	return &router.Router{}, nil
}
