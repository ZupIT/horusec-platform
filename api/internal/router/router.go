package router

import (
	"github.com/ZupIT/horusec-platform/api/internal/handlers/analysis"
	"github.com/ZupIT/horusec-platform/api/internal/handlers/health"
	"github.com/ZupIT/horusec-platform/api/internal/middelwares/token"
	"github.com/go-chi/chi"

	"github.com/ZupIT/horusec-devkit/pkg/services/http"
)

type IRouter interface {
	http.IRouter
}

type Router struct {
	http.IRouter
	analysisHandler *analysis.Handler
	healthHandler   *health.Handler
	tokenAuthz      token.ITokenAuthz
}

func NewHTTPRouter(router http.IRouter, analysisHandler *analysis.Handler, healthHandler *health.Handler,
	tokenAuthz token.ITokenAuthz) IRouter {
	routes := &Router{
		IRouter:         router,
		analysisHandler: analysisHandler,
		healthHandler:   healthHandler,
		tokenAuthz:      tokenAuthz,
	}

	return routes.setRoutes()
}

func (r *Router) setRoutes() IRouter {
	r.routerHealth()
	r.routerAnalysis()

	return r
}

func (r *Router) routerAnalysis() {
	r.Route("/api/analysis", func(router chi.Router) {
		router.Use(r.tokenAuthz.IsAuthorized)
		router.Options("/", r.analysisHandler.Options)
		router.Post("/", r.analysisHandler.Post)
		router.Get("/{analysisID}", r.analysisHandler.Get)
	})
}

func (r *Router) routerHealth() {
	r.Route("/api/health", func(router chi.Router) {
		router.Options("/", r.healthHandler.Options)
		router.Get("/", r.healthHandler.Get)
	})
}
