package router

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/http/router"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"
	"github.com/go-chi/chi"

	webhookEvent "github.com/ZupIT/horusec-platform/webhook/internal/events/webhook"
	"github.com/ZupIT/horusec-platform/webhook/internal/handlers/webhook"

	"github.com/ZupIT/horusec-platform/webhook/internal/handlers/health"

	"github.com/ZupIT/horusec-platform/webhook/docs"

	"github.com/ZupIT/horusec-platform/webhook/internal/enums"

	"github.com/ZupIT/horusec-devkit/pkg/services/swagger"
)

type IRouter interface {
	router.IRouter
}

type Router struct {
	router.IRouter
	swagger.ISwagger
	middlewares.IAuthzMiddleware
	healthHandler  *health.Handler
	webhookHandler *webhook.Handler
	webhookEvents  webhookEvent.IEvent
}

func NewHTTPRouter(routerConn router.IRouter, authzMiddleware middlewares.IAuthzMiddleware,
	healthHandler *health.Handler, webhookHandler *webhook.Handler, webhookEvents webhookEvent.IEvent) IRouter {
	routes := &Router{
		IRouter:          routerConn,
		IAuthzMiddleware: authzMiddleware,
		ISwagger:         swagger.NewSwagger(routerConn.GetMux(), enums.DefaultPort),
		healthHandler:    healthHandler,
		webhookHandler:   webhookHandler,
		webhookEvents:    webhookEvents,
	}
	return routes.setRoutes()
}

func (r *Router) setRoutes() IRouter {
	r.routerSwagger()
	r.routerHealth()
	r.routerWebhook()
	return r
}

func (r *Router) routerSwagger() {
	r.SetupSwagger()
	docs.SwaggerInfo.Host = r.GetSwaggerHost()
}

func (r *Router) routerHealth() {
	r.Route(enums.HealthRouter, func(router chi.Router) {
		router.Options("/", r.healthHandler.Options)
		router.Get("/", r.healthHandler.Get)
	})
}

func (r *Router) routerWebhook() {
	r.Route(enums.WebhookRouter, func(router chi.Router) {
		router.Options("/", r.webhookHandler.Options)
		router.With(r.IsWorkspaceAdmin).Get("/", r.webhookHandler.ListAll)
		router.With(r.IsWorkspaceAdmin).Post("/", r.webhookHandler.Save)
		router.With(r.IsWorkspaceAdmin).Put("/{webhookID}", r.webhookHandler.Update)
		router.With(r.IsWorkspaceAdmin).Delete("/{webhookID}", r.webhookHandler.Remove)
	})
}
