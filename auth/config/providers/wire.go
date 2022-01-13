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
	"github.com/google/wire"

	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	brokerConfig "github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/cache"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseConfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	httpRouter "github.com/ZupIT/horusec-devkit/pkg/services/http/router"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	"github.com/ZupIT/horusec-platform/auth/config/cors"
	"github.com/ZupIT/horusec-platform/auth/config/grpc"
	accountController "github.com/ZupIT/horusec-platform/auth/internal/controllers/account"
	authController "github.com/ZupIT/horusec-platform/auth/internal/controllers/authentication"
	accountHandler "github.com/ZupIT/horusec-platform/auth/internal/handlers/account"
	authHandler "github.com/ZupIT/horusec-platform/auth/internal/handlers/authentication"
	healthHandler "github.com/ZupIT/horusec-platform/auth/internal/handlers/health"
	accountRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/account"
	authRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/authentication"
	"github.com/ZupIT/horusec-platform/auth/internal/router"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/horusec"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/keycloak"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/ldap"
	accountUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/account"
	authUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/authentication"
)

var devKitProviders = wire.NewSet(
	databaseConfig.NewDatabaseConfig,
	brokerConfig.NewBrokerConfig,
	broker.NewBroker,
	database.NewDatabaseReadAndWrite,
	cache.NewCache,
	httpRouter.NewHTTPRouter,
)

var configProviders = wire.NewSet(
	grpc.NewAuthGRPCServer,
	cors.NewCorsConfig,
	app.NewAuthAppConfig,
	router.NewHTTPRouter,
)

var controllerProviders = wire.NewSet(
	authController.NewAuthenticationController,
	accountController.NewAccountController,
)

var handleProviders = wire.NewSet(
	authHandler.NewAuthenticationHandler,
	accountHandler.NewAccountHandler,
	healthHandler.NewHealthHandler,
)

var useCasesProviders = wire.NewSet(
	authUseCases.NewAuthenticationUseCases,
	accountUseCases.NewAccountUseCases,
)

var repositoriesProviders = wire.NewSet(
	accountRepository.NewAccountRepository,
	authRepository.NewAuthenticationRepository,
)

var serviceProviders = wire.NewSet(
	horusec.NewHorusecAuthenticationService,
	ldap.NewLDAPAuthenticationService,
	keycloak.NewKeycloakAuthenticationService,
)

func Initialize(_ string) (router.IRouter, error) {
	wire.Build(devKitProviders, configProviders, controllerProviders, handleProviders,
		useCasesProviders, repositoriesProviders, serviceProviders)

	return &router.Router{}, nil
}
