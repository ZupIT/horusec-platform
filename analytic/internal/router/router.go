package router

import (
	"github.com/go-chi/chi"

	httpRouter "github.com/ZupIT/horusec-devkit/pkg/services/http/router"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"
	"github.com/ZupIT/horusec-devkit/pkg/services/swagger"

	"github.com/ZupIT/horusec-platform/analytic/docs"
	"github.com/ZupIT/horusec-platform/analytic/internal/enums/routes"
	eventsDashboard "github.com/ZupIT/horusec-platform/analytic/internal/events/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/health"
)

type IRouter interface {
	httpRouter.IRouter
}

type Router struct {
	httpRouter.IRouter
	swagger.ISwagger
	middlewares.IAuthzMiddleware
	healthHandler    *health.Handler
	dashboardHandler *dashboard.Handler
	dashboardEvent   eventsDashboard.IEvents
}

func NewHTTPRouter(router httpRouter.IRouter, authzMiddleware middlewares.IAuthzMiddleware,
	healthHandler *health.Handler, dashboardHandler *dashboard.Handler, dashboardEvent eventsDashboard.IEvents) IRouter {
	requestRouter := &Router{
		IRouter:          router,
		IAuthzMiddleware: authzMiddleware,
		ISwagger:         swagger.NewSwagger(router.GetMux(), router.GetPort()),
		healthHandler:    healthHandler,
		dashboardHandler: dashboardHandler,
		dashboardEvent:   dashboardEvent,
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
