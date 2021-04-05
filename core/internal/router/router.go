package router

import (
	"github.com/go-chi/chi"

	"github.com/ZupIT/horusec-devkit/pkg/services/http"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"
	"github.com/ZupIT/horusec-devkit/pkg/services/swagger"

	"github.com/ZupIT/horusec-platform/core/docs"
	"github.com/ZupIT/horusec-platform/core/internal/enums/routes"
	"github.com/ZupIT/horusec-platform/core/internal/handlers/workspace"
)

type IRouter interface {
	http.IRouter
}

type Router struct {
	http.IRouter
	middlewares.IAuthzMiddleware
	workspaceHandler *workspace.Handler
	swagger.ISwagger
}

func NewHTTPRouter(router http.IRouter, workspaceHandler *workspace.Handler,
	authzMiddleware middlewares.IAuthzMiddleware) IRouter {
	httpRoutes := &Router{
		IRouter:          router,
		IAuthzMiddleware: authzMiddleware,
		workspaceHandler: workspaceHandler,
		ISwagger:         swagger.NewSwagger(router.GetMux(), router.GetPort()),
	}

	return httpRoutes.setRoutes()
}

func (r *Router) setRoutes() IRouter {
	r.swaggerRoutes()
	r.workspaceRoutes()

	return r
}

func (r *Router) workspaceRoutes() {
	r.Route(routes.WorkspaceHandler, func(router chi.Router) {
		router.With(r.IsApplicationAdmin).Post("/", r.workspaceHandler.Create)
	})
}

func (r *Router) swaggerRoutes() {
	r.SetupSwagger()
	docs.SwaggerInfo.Host = r.GetSwaggerHost()
}
