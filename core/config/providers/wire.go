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

	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	brokerConfig "github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseConfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	httpRouter "github.com/ZupIT/horusec-devkit/pkg/services/http/router"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"

	"github.com/ZupIT/horusec-platform/core/config/cors"
	repositoryController "github.com/ZupIT/horusec-platform/core/internal/controllers/repository"
	workspaceController "github.com/ZupIT/horusec-platform/core/internal/controllers/workspace"
	healthHandler "github.com/ZupIT/horusec-platform/core/internal/handlers/health"
	repositoryHandler "github.com/ZupIT/horusec-platform/core/internal/handlers/repository"
	workspaceHandler "github.com/ZupIT/horusec-platform/core/internal/handlers/workspace"
	repositoryRepository "github.com/ZupIT/horusec-platform/core/internal/repositories/repository"
	workspaceRepository "github.com/ZupIT/horusec-platform/core/internal/repositories/workspace"
	"github.com/ZupIT/horusec-platform/core/internal/router"
	repositoryUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/repository"
	roleUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/role"
	"github.com/ZupIT/horusec-platform/core/internal/usecases/token"
	workspaceUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/workspace"
)

var devKitProviders = wire.NewSet(
	brokerConfig.NewBrokerConfig,
	broker.NewBroker,
	databaseConfig.NewDatabaseConfig,
	database.NewDatabaseReadAndWrite,
	httpRouter.NewHTTPRouter,
	auth.NewAuthGRPCConnection,
	proto.NewAuthServiceClient,
	app.NewAppConfig,
	middlewares.NewAuthzMiddleware,
)

var configProviders = wire.NewSet(
	cors.NewCorsConfig,
	router.NewHTTPRouter,
)

var controllerProviders = wire.NewSet(
	workspaceController.NewWorkspaceController,
	repositoryController.NewRepositoryController,
)

var handleProviders = wire.NewSet(
	workspaceHandler.NewWorkspaceHandler,
	repositoryHandler.NewRepositoryHandler,
	healthHandler.NewHealthHandler,
)

var useCasesProviders = wire.NewSet(
	workspaceUseCases.NewWorkspaceUseCases,
	repositoryUseCases.NewRepositoryUseCases,
	roleUseCases.NewRoleUseCases,
	token.NewTokenUseCases,
)

var repositoriesProviders = wire.NewSet(
	workspaceRepository.NewWorkspaceRepository,
	repositoryRepository.NewRepositoryRepository,
)

func Initialize(_ string) (router.IRouter, error) {
	wire.Build(devKitProviders, configProviders, controllerProviders, handleProviders,
		useCasesProviders, repositoriesProviders)

	return &router.Router{}, nil
}
