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
