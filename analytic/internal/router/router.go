package router

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"
	dashboardRepository "github.com/ZupIT/horusec-platform/analytic/internal/handlers/dashboard_repository"
	dashboardWorkspace "github.com/ZupIT/horusec-platform/analytic/internal/handlers/dashboard_workspace"

	"github.com/ZupIT/horusec-platform/analytic/docs"
	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/health"

	"github.com/go-chi/chi"

	"github.com/ZupIT/horusec-platform/analytic/internal/router/enums"

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
	healthHandler *health.Handler
	dashboardWorkspaceHandler *dashboardWorkspace.Handler
	dashboardRepositoryHandler *dashboardRepository.Handler
}

func NewHTTPRouter(router http.IRouter, authzMiddleware middlewares.IAuthzMiddleware, healthHandler *health.Handler,
	dashboardWorkspaceHandler *dashboardWorkspace.Handler, dashboardRepositoryHandler *dashboardRepository.Handler) IRouter {
	routes := &Router{
		IRouter:          router,
		IAuthzMiddleware: authzMiddleware,
		ISwagger:         swagger.NewSwagger(router.GetMux(), enums.DefaultPort),
		healthHandler:    healthHandler,
		dashboardWorkspaceHandler:  dashboardWorkspaceHandler,
		dashboardRepositoryHandler: dashboardRepositoryHandler,
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
		router.Options("/", r.dashboardWorkspaceHandler.Options)
		router.Get("/total-developers", r.dashboardWorkspaceHandler.GetTotalDevelopers)
		router.Get("/total-repositories", r.dashboardWorkspaceHandler.GetTotalRepositories)
		router.Get("/vulnerabilities-by-severities", r.dashboardWorkspaceHandler.GetVulnBySeverity)
		router.Get("/vulnerabilities-by-developer", r.dashboardWorkspaceHandler.GetVulnByDeveloper)
		router.Get("/vulnerabilities-by-repositories", r.dashboardWorkspaceHandler.GetVulnByRepository)
		router.Get("/vulnerabilities-by-languages", r.dashboardWorkspaceHandler.GetVulnByLanguage)
		router.Get("/vulnerabilities-by-time", r.dashboardWorkspaceHandler.GetVulnByTime)
		router.Get("/details", r.dashboardWorkspaceHandler.GetVulnDetails)
	})
}

func (r *Router) routerDashboardRepository() {
	r.Route(enums.DashboardRepositoryRouter, func(router chi.Router) {
		router.Options("/", r.dashboardWorkspaceHandler.Options)
		router.Get("/total-developers", r.dashboardWorkspaceHandler.GetTotalDevelopers)
		router.Get("/vulnerabilities-by-severities", r.dashboardWorkspaceHandler.GetVulnBySeverity)
		router.Get("/vulnerabilities-by-developer", r.dashboardWorkspaceHandler.GetVulnByDeveloper)
		router.Get("/vulnerabilities-by-languages", r.dashboardWorkspaceHandler.GetVulnByLanguage)
		router.Get("/vulnerabilities-by-time", r.dashboardWorkspaceHandler.GetVulnByTime)
		router.Get("/details", r.dashboardWorkspaceHandler.GetVulnDetails)
	})
}
