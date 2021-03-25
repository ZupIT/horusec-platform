package router

import (
	"github.com/ZupIT/horusec-platform/api/internal/handlers/analysis"
	"github.com/ZupIT/horusec-platform/api/internal/handlers/health"
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
}

func NewHTTPRouter(router http.IRouter, analysisHandler *analysis.Handler, healthHandler *health.Handler) IRouter {
	routes := &Router{
		IRouter:         router,
		analysisHandler: analysisHandler,
		healthHandler:   healthHandler,
	}

	return routes.setRoutes()
}

func (r *Router) setRoutes() IRouter {
	r.routerHealth()

	return r
}

func (r *Router) routerHealth() {
	r.Route("/api", func(router chi.Router) {
		router.Get("/health", r.healthHandler.Get)
	})
}
