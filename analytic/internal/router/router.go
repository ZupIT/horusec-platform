package router

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"

	eventDashboard "github.com/ZupIT/horusec-platform/analytic/internal/events/dashboard"

	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/dashboard"

	"github.com/ZupIT/horusec-platform/analytic/docs"
	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/health"

	"github.com/go-chi/chi"

	"github.com/ZupIT/horusec-platform/analytic/internal/enums"

	"github.com/ZupIT/horusec-devkit/pkg/services/swagger"

	"github.com/ZupIT/horusec-devkit/pkg/services/http"
)

type IRouter interface {
	http.IRouter
}

type Router struct {
	http.IRouter
	swagger.ISwagger
	middlewares.IAuthzMiddleware
	healthHandler    *health.Handler
	dashboardHandler *dashboard.Handler
	dashboardEvent   eventDashboard.IEvent
}

func NewHTTPRouter(router http.IRouter, authzMiddleware middlewares.IAuthzMiddleware, healthHandler *health.Handler,
	dashboardHandler *dashboard.Handler, dashboardEvent eventDashboard.IEvent) IRouter {
	routes := &Router{
		IRouter:          router,
		IAuthzMiddleware: authzMiddleware,
		ISwagger:         swagger.NewSwagger(router.GetMux(), enums.DefaultPort),
		healthHandler:    healthHandler,
		dashboardHandler: dashboardHandler,
		dashboardEvent:   dashboardEvent,
	}
	return routes.setRoutes()
}

func (r *Router) setRoutes() IRouter {
	r.routerHealth()
	r.routerSwagger()
	r.routerDashboardWorkspace()
	return r
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

func (r *Router) routerDashboardWorkspace() {
	r.Route(enums.DashboardWorkspaceRouter, func(router chi.Router) {
		router.Options("/", r.dashboardHandler.Options)
		router.With(r.IsWorkspaceAdmin).Get("/dashboard-charts", r.dashboardHandler.GetAllChartsByWorkspace)
		router.With(r.IsRepositoryAdmin).Get("/{repositoryID}/dashboard-charts", r.dashboardHandler.GetAllChartsByRepository)
	})
}