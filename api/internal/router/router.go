package router

import (
	"github.com/ZupIT/horusec-platform/api/docs"
	"github.com/ZupIT/horusec-platform/api/internal/handlers/analysis"
	"github.com/ZupIT/horusec-platform/api/internal/handlers/health"
	"github.com/ZupIT/horusec-platform/api/internal/middelwares/token"
	"github.com/ZupIT/horusec-platform/api/internal/router/enums"
	"github.com/go-chi/chi"

	"github.com/ZupIT/horusec-devkit/pkg/services/swagger"

	"github.com/ZupIT/horusec-devkit/pkg/services/http"
)

type IRouter interface {
	http.IRouter
}

type Router struct {
	http.IRouter
	swagger.ISwagger
	analysisHandler *analysis.Handler
	healthHandler   *health.Handler
	tokenAuthz      token.ITokenAuthz
}

func NewHTTPRouter(router http.IRouter, tokenAuthz token.ITokenAuthz,
	analysisHandler *analysis.Handler, healthHandler *health.Handler) IRouter {
	routes := &Router{
		IRouter:         router,
		ISwagger:        swagger.NewSwagger(router.GetMux(), enums.DefaultPort),
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
