package router

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/http"
	"github.com/ZupIT/horusec-platform/core/internal/handlers/workspace"
	"github.com/go-chi/chi"
)

type IRouter interface {
	http.IRouter
}

type Router struct {
	http.IRouter
	workspaceHandler *workspace.Handler
}

func NewHTTPRouter(router http.IRouter, workspaceHandler *workspace.Handler) IRouter {
	routes := &Router{
		IRouter:          router,
		workspaceHandler: workspaceHandler,
	}

	return routes.setRoutes()
}

func (r *Router) setRoutes() IRouter {
	r.routerTest()

	return r
}

func (r *Router) routerTest() {
	r.Route("/test", func(router chi.Router) {
		router.Get("/", r.workspaceHandler.Get)
	})
}
