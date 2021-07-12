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

package router

import (
	"github.com/go-chi/chi"

	"github.com/ZupIT/horusec-devkit/pkg/services/http/router"
	"github.com/ZupIT/horusec-devkit/pkg/services/swagger"

	"github.com/ZupIT/horusec-platform/auth/config/grpc"
	"github.com/ZupIT/horusec-platform/auth/docs"
	"github.com/ZupIT/horusec-platform/auth/internal/enums/routes"
	accountHandler "github.com/ZupIT/horusec-platform/auth/internal/handlers/account"
	authHandler "github.com/ZupIT/horusec-platform/auth/internal/handlers/authentication"
	"github.com/ZupIT/horusec-platform/auth/internal/handlers/health"
)

type IRouter interface {
	router.IRouter
}

type Router struct {
	router.IRouter
	swagger.ISwagger
	grpc.IAuthGRPCServer
	authHandler    *authHandler.Handler
	accountHandler *accountHandler.Handler
	healthHandler  *health.Handler
}

func NewHTTPRouter(routerConnection router.IRouter, authGRPCServer grpc.IAuthGRPCServer,
	handlerAuth *authHandler.Handler, handlerAccount *accountHandler.Handler, handlerHealth *health.Handler) IRouter {
	httpRouter := &Router{
		IRouter:         routerConnection,
		ISwagger:        swagger.NewSwagger(routerConnection.GetMux(), routerConnection.GetPort()),
		IAuthGRPCServer: authGRPCServer,
		authHandler:     handlerAuth,
		accountHandler:  handlerAccount,
		healthHandler:   handlerHealth,
	}

	httpRouter.startGRPCServer()
	return httpRouter.setRoutes()
}

func (r *Router) setRoutes() IRouter {
	r.swaggerRoutes()

	return r
}

func (r *Router) startGRPCServer() {
	go r.ListenAndServeGRPCServer()
}

func (r *Router) swaggerRoutes() {
	r.SetupSwagger()
	r.authenticationRoutes()
	r.accountRoutes()
	r.healthRoutes()

	docs.SwaggerInfo.Host = r.GetSwaggerHost()
}

func (r *Router) authenticationRoutes() {
	r.Route(routes.AuthenticationHandler, func(router chi.Router) {
		router.Post("/login", r.authHandler.Login)
		router.Get("/config", r.authHandler.GetConfig)
	})
}

func (r *Router) accountRoutes() {
	r.Route(routes.AccountHandler, func(router chi.Router) {
		router.Post("/create-account-keycloak", r.accountHandler.CreateAccountKeycloak)
		router.Post("/create-account-horusec", r.accountHandler.CreateAccountHorusec)
		router.Get("/validate/{accountID}", r.accountHandler.ValidateAccountEmail)
		router.Post("/send-reset-code", r.accountHandler.SendResetPasswordCode)
		router.Post("/check-reset-code", r.accountHandler.CheckResetPasswordCode)
		router.Post("/change-password", r.accountHandler.ChangePassword)
		router.Post("/refresh-token", r.accountHandler.RefreshToken)
		router.Post("/logout", r.accountHandler.Logout)
		router.Delete("/delete", r.accountHandler.DeleteAccount)
		router.Post("/verify-already-used", r.accountHandler.CheckExistingEmailOrUsername)
		router.Patch("/update", r.accountHandler.UpdateAccount)
	})
}

func (r *Router) healthRoutes() {
	r.Route(routes.HealthHandler, func(router chi.Router) {
		router.Get("/", r.healthHandler.Get)
	})
}
