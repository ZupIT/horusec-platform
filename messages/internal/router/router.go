package router

import (
	"github.com/go-chi/chi"

	httpRouter "github.com/ZupIT/horusec-devkit/pkg/services/http/router"
	"github.com/ZupIT/horusec-devkit/pkg/services/swagger"

	"github.com/ZupIT/horusec-platform/messages/internal/enums/routes"
	"github.com/ZupIT/horusec-platform/messages/internal/handlers/health"
)

type IRouter interface {
	httpRouter.IRouter
}

type Router struct {
	httpRouter.IRouter
	swagger.ISwagger
	healthHandler *health.Handler
}

func NewHTTPRouter(router httpRouter.IRouter, handlerHealth *health.Handler) IRouter {
	httpRoutes := &Router{
		IRouter:       router,
		ISwagger:      swagger.NewSwagger(router.GetMux(), router.GetPort()),
		healthHandler: handlerHealth,
	}

	return httpRoutes.setRoutes()
}

func (r *Router) setRoutes() IRouter {
	r.swaggerRoutes()
	r.healthRoutes()

	return r
}

func (r *Router) healthRoutes() {
	r.Route(routes.HealthHandler, func(router chi.Router) {
		router.Get("/", r.healthHandler.Get)
	})
}

func (r *Router) swaggerRoutes() {
	//r.SetupSwagger()
	//
	//docs.SwaggerInfo.Host = r.GetSwaggerHost()
}
