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
	"github.com/ZupIT/horusec-devkit/pkg/services/swagger"
	"github.com/ZupIT/horusec-platform/api/docs"
	"github.com/ZupIT/horusec-platform/api/internal/enums"
	"github.com/ZupIT/horusec-platform/api/internal/handlers/analysis"
	"github.com/ZupIT/horusec-platform/api/internal/handlers/health"
	"github.com/ZupIT/horusec-platform/api/internal/middelwares/token"
	"github.com/go-chi/chi/v5"
)

type IRouter interface {
	router.IRouter
}

type Router struct {
	router.IRouter
	swagger.ISwagger
	analysisHandler *analysis.Handler
	healthHandler   *health.Handler
	tokenAuthz      token.ITokenAuthz
}

func NewHTTPRouter(routerConnection router.IRouter, tokenAuthz token.ITokenAuthz,
	analysisHandler *analysis.Handler, healthHandler *health.Handler) IRouter {
	routes := &Router{
		IRouter:         routerConnection,
		ISwagger:        swagger.NewSwagger(routerConnection.GetMux(), enums.DefaultPort),
		analysisHandler: analysisHandler,
		healthHandler:   healthHandler,
		tokenAuthz:      tokenAuthz,
	}
	return routes.setRoutes()
}

func (r *Router) setRoutes() IRouter {
	r.routerHealth()
	r.routerAnalysis()
	r.routerSwagger()
	return r
}

func (r *Router) routerAnalysis() {
	r.Route(enums.AnalysisRouter, func(router chi.Router) {
		router.Use(r.tokenAuthz.IsAuthorized)
		router.Options("/", r.analysisHandler.Options)
		router.Post("/", r.analysisHandler.Post)
		router.Get("/{analysisID}", r.analysisHandler.Get)
	})
}

func (r *Router) routerHealth() {
	r.Route(enums.HealthRouter, func(router chi.Router) {
		router.Options("/", r.healthHandler.Options)
		router.Get("/", r.healthHandler.Get)
	})
}

func (r *Router) routerSwagger() {
	r.SetupSwagger()
	docs.SwaggerInfo.Host = r.GetSwaggerHost()
}
