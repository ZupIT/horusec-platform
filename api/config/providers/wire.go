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
	routerHttp "github.com/ZupIT/horusec-devkit/pkg/services/http/router"
	"github.com/google/wire"

	analysisHandler "github.com/ZupIT/horusec-platform/api/internal/handlers/analysis"
	healthHandler "github.com/ZupIT/horusec-platform/api/internal/handlers/health"
	"github.com/ZupIT/horusec-platform/api/internal/middelwares/token"
	"github.com/ZupIT/horusec-platform/api/internal/repositories/analysis"
	"github.com/ZupIT/horusec-platform/api/internal/repositories/repository"
	repositoriesToken "github.com/ZupIT/horusec-platform/api/internal/repositories/token"
	"github.com/ZupIT/horusec-platform/api/internal/router"

	appConfiguration "github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"

	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	brokerConfig "github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseConfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth"

	"github.com/ZupIT/horusec-platform/api/config/cors"
	analysisController "github.com/ZupIT/horusec-platform/api/internal/controllers/analysis"
)

var providers = wire.NewSet(
	brokerConfig.NewBrokerConfig,
	broker.NewBroker,
	databaseConfig.NewDatabaseConfig,
	database.NewDatabaseReadAndWrite,
	auth.NewAuthGRPCConnection,
	proto.NewAuthServiceClient,
	token.NewTokenAuthz,
	analysis.NewRepositoriesAnalysis,
	repository.NewRepositoriesRepository,
	repositoriesToken.NewRepositoriesToken,
	cors.NewCorsConfig,
	routerHttp.NewHTTPRouter,
	appConfiguration.NewAppConfig,
	analysisController.NewAnalysisController,
	analysisHandler.NewAnalysisHandler,
	healthHandler.NewHealthHandler,
	router.NewHTTPRouter,
)

func Initialize(defaultPort string) (router.IRouter, error) {
	wire.Build(providers)
	return &router.Router{}, nil
}
