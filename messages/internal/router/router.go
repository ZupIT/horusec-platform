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

	httpRouter "github.com/ZupIT/horusec-devkit/pkg/services/http/router"
	"github.com/ZupIT/horusec-devkit/pkg/services/swagger"

	"github.com/ZupIT/horusec-platform/messages/docs"
	"github.com/ZupIT/horusec-platform/messages/internal/enums/routes"
	"github.com/ZupIT/horusec-platform/messages/internal/events/email"
	"github.com/ZupIT/horusec-platform/messages/internal/handlers/health"
)

type IRouter interface {
	httpRouter.IRouter
}

type Router struct {
	httpRouter.IRouter
	swagger.ISwagger
	healthHandler     *health.Handler
	emailEventHandler *email.EventHandler
}

func NewHTTPRouter(router httpRouter.IRouter, handlerHealth *health.Handler,
	emailEventHandler *email.EventHandler) IRouter {
	httpRoutes := &Router{
		IRouter:           router,
		ISwagger:          swagger.NewSwagger(router.GetMux(), router.GetPort()),
		healthHandler:     handlerHealth,
		emailEventHandler: emailEventHandler,
	}

	return httpRoutes.setRoutes()
}

func (r *Router) setRoutes() IRouter {
	r.swaggerRoutes()
	r.healthRoutes()
	r.emailEventHandler.StartConsumers()

	return r
}

func (r *Router) healthRoutes() {
	r.Route(routes.HealthHandler, func(router chi.Router) {
		router.Options("/", r.healthHandler.Options)
		router.Get("/", r.healthHandler.Get)
	})
}

func (r *Router) swaggerRoutes() {
	r.SetupSwagger()

	docs.SwaggerInfo.Host = r.GetSwaggerHost()
}
