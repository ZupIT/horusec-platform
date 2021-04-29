package router

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/http/router"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"

	eventDashboard "github.com/ZupIT/horusec-platform/analytic/internal/events/dashboard"

	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/dashboard"

	"github.com/ZupIT/horusec-platform/analytic/docs"
	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/health"

	"github.com/go-chi/chi"

	"github.com/ZupIT/horusec-platform/analytic/internal/enums"

	"github.com/ZupIT/horusec-devkit/pkg/services/swagger"
)

type IRouter interface {
	router.IRouter
}

type Router struct {
	router.IRouter
	swagger.ISwagger
	middlewares.IAuthzMiddleware
	healthHandler    *health.Handler
	dashboardHandler *dashboard.Handler
	dashboardEvent   eventDashboard.IEvent
}

func NewHTTPRouter(routerConn router.IRouter, authzMiddleware middlewares.IAuthzMiddleware,
	healthHandler *health.Handler, dashboardHandler *dashboard.Handler, dashboardEvent eventDashboard.IEvent) IRouter {
	routes := &Router{
		IRouter:          routerConn,
		IAuthzMiddleware: authzMiddleware,
		ISwagger:         swagger.NewSwagger(routerConn.GetMux(), enums.DefaultPort),
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
