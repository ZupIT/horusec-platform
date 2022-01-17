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
	"github.com/ZupIT/horusec-devkit/pkg/services/http/router"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"
	"github.com/ZupIT/horusec-devkit/pkg/services/swagger"
	"github.com/go-chi/chi"

	"github.com/ZupIT/horusec-platform/analytic/docs"
	"github.com/ZupIT/horusec-platform/analytic/internal/enums/routes"
	eventsdashboard "github.com/ZupIT/horusec-platform/analytic/internal/events/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/health"
)

type IRouter interface {
	router.IRouter
}

type Router struct {
	router.IRouter
	swagger.ISwagger
	middlewares.IAuthzMiddleware
	healthHandler    *health.Handler
	dashboardHandler *dashboard.Handler
	dashboardEvents  *eventsdashboard.Events
}

func NewHTTPRouter(route router.IRouter, authzMiddleware middlewares.IAuthzMiddleware,
	healthHandler *health.Handler, dashboardHandler *dashboard.Handler, eventsDashboard *eventsdashboard.Events) IRouter {
	requestRouter := &Router{
		IRouter:          route,
		IAuthzMiddleware: authzMiddleware,
		ISwagger:         swagger.NewSwagger(route.GetMux(), route.GetPort()),
		healthHandler:    healthHandler,
		dashboardHandler: dashboardHandler,
		dashboardEvents:  eventsDashboard,
	}

	return requestRouter.setRoutes()
}

func (r *Router) setRoutes() IRouter {
	r.routerHealth()
	r.routerSwagger()
	r.routerDashboardWorkspace()

	return r
}

func (r *Router) routerHealth() {
	r.Route(routes.HealthRouter, func(router chi.Router) {
		router.Options("/", r.healthHandler.Options)
		router.Get("/", r.healthHandler.Get)
	})
}

func (r *Router) routerSwagger() {
	r.SetupSwagger()

	docs.SwaggerInfo.Host = r.GetSwaggerHost()
}

func (r *Router) routerDashboardWorkspace() {
	r.Route(routes.DashboardWorkspaceRouter, func(router chi.Router) {
		router.Options("/", r.dashboardHandler.Options)
		router.With(r.IsWorkspaceAdmin).Get("/", r.dashboardHandler.GetAllChartsByWorkspace)
		router.With(r.IsRepositoryMember).Get("/{repositoryID}", r.dashboardHandler.GetAllChartsByRepository)
	})
}
