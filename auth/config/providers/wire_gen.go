// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package providers

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	config2 "github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/http"
	"github.com/google/wire"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	"github.com/ZupIT/horusec-platform/auth/config/cors"
	"github.com/ZupIT/horusec-platform/auth/config/grpc"
	account3 "github.com/ZupIT/horusec-platform/auth/internal/controllers/account"
	authentication3 "github.com/ZupIT/horusec-platform/auth/internal/controllers/authentication"
	account4 "github.com/ZupIT/horusec-platform/auth/internal/handlers/account"
	authentication4 "github.com/ZupIT/horusec-platform/auth/internal/handlers/authentication"
	account2 "github.com/ZupIT/horusec-platform/auth/internal/repositories/account"
	authentication2 "github.com/ZupIT/horusec-platform/auth/internal/repositories/authentication"
	"github.com/ZupIT/horusec-platform/auth/internal/router"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/horusec"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/keycloak"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/ldap"
	"github.com/ZupIT/horusec-platform/auth/internal/usecases/account"
	"github.com/ZupIT/horusec-platform/auth/internal/usecases/authentication"
)

// Injectors from wire.go:

func Initialize(string2 string) (router.IRouter, error) {
	options := cors.NewCorsConfig()
	iRouter := http.NewHTTPRouter(options, string2)
	iConfig := app.NewAuthAppConfig()
	iUseCases := authentication.NewAuthenticationUseCases()
	configIConfig := config.NewDatabaseConfig()
	connection, err := database.NewDatabaseReadAndWrite(configIConfig)
	if err != nil {
		return nil, err
	}
	accountIUseCases := account.NewAccountUseCases()
	iRepository := account2.NewAccountRepository(connection, accountIUseCases)
	authenticationIRepository := authentication2.NewAuthenticationRepository(connection, iUseCases)
	iService := horusec.NewHorusecAuthenticationService(iRepository, iConfig, iUseCases, authenticationIRepository)
	ldapIService := ldap.NewLDAPAuthenticationService(iRepository, iUseCases, iConfig, authenticationIRepository)
	keycloakIService := keycloak.NewKeycloakAuthenticationService(iRepository, iConfig, iUseCases, authenticationIRepository)
	iController := authentication3.NewAuthenticationController(iConfig, iService, ldapIService, keycloakIService)
	handler := authentication4.NewAuthenticationHandler(iConfig, iUseCases, iController)
	iAuthGRPCServer := grpc.NewAuthGRPCServer(handler)
	routerIRouter := router.NewHTTPRouter(iRouter, iAuthGRPCServer, handler)
	return routerIRouter, nil
}

// wire.go:

var devKitProviders = wire.NewSet(http.NewHTTPRouter, config.NewDatabaseConfig, config2.NewBrokerConfig, broker.NewBroker, database.NewDatabaseReadAndWrite)

var configProviders = wire.NewSet(grpc.NewAuthGRPCServer, cors.NewCorsConfig, router.NewHTTPRouter, app.NewAuthAppConfig)

var controllerProviders = wire.NewSet(authentication3.NewAuthenticationController, account3.NewAccountController)

var handleProviders = wire.NewSet(authentication4.NewAuthenticationHandler, account4.NewAccountHandler)

var useCasesProviders = wire.NewSet(authentication.NewAuthenticationUseCases, account.NewAccountUseCases)

var repositoriesProviders = wire.NewSet(account2.NewAccountRepository, authentication2.NewAuthenticationRepository)

var serviceProviders = wire.NewSet(horusec.NewHorusecAuthenticationService, ldap.NewLDAPAuthenticationService, keycloak.NewKeycloakAuthenticationService)
