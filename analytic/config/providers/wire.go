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

//go:build wireinject
// +build wireinject

package providers

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseConfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	routerHttp "github.com/ZupIT/horusec-devkit/pkg/services/http/router"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"
	"github.com/google/wire"

	"github.com/ZupIT/horusec-platform/analytic/config/cors"
	dashboardController "github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
	dashboardEvents "github.com/ZupIT/horusec-platform/analytic/internal/events/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/health"
	dashboardRepository "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/router"
	dashboardUseCases "github.com/ZupIT/horusec-platform/analytic/internal/usecases/dashboard"
)

var devKitProviders = wire.NewSet(
	auth.NewAuthGRPCConnection,
	proto.NewAuthServiceClient,
	app.NewAppConfig,
	config.NewBrokerConfig,
	broker.NewBroker,
	databaseConfig.NewDatabaseConfig,
	database.NewDatabaseReadAndWrite,
	routerHttp.NewHTTPRouter,
	middlewares.NewAuthzMiddleware,
)

var configProviders = wire.NewSet(
	cors.NewCorsConfig,
	router.NewHTTPRouter,
)

var repositoriesProviders = wire.NewSet(
	dashboardRepository.NewRepoDashboard,
	dashboardRepository.NewWorkspaceDashboard,
)

var controllersProviders = wire.NewSet(
	dashboardController.NewDashboardController,
)

var handlersProviders = wire.NewSet(
	health.NewHealthHandler,
	dashboard.NewDashboardHandler,
)

var eventsProviders = wire.NewSet(
	dashboardEvents.NewDashboardEvents,
)

var useCasesProviders = wire.NewSet(
	dashboardUseCases.NewUseCaseDashboard,
)

func Initialize(_ string) (router.IRouter, error) {
	wire.Build(devKitProviders, configProviders, repositoriesProviders, controllersProviders,
		handlersProviders, eventsProviders, useCasesProviders)

	return &router.Router{}, nil
}
