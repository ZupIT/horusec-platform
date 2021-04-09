package router

import (
	"github.com/go-chi/chi"

	"github.com/ZupIT/horusec-devkit/pkg/services/http"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"
	"github.com/ZupIT/horusec-devkit/pkg/services/swagger"

	"github.com/ZupIT/horusec-platform/core/docs"
	"github.com/ZupIT/horusec-platform/core/internal/enums/routes"
	"github.com/ZupIT/horusec-platform/core/internal/handlers/repository"
	"github.com/ZupIT/horusec-platform/core/internal/handlers/workspace"
)

type IRouter interface {
	http.IRouter
}

type Router struct {
	http.IRouter
	middlewares.IAuthzMiddleware
	workspaceHandler  *workspace.Handler
	repositoryHandler *repository.Handler
	swagger.ISwagger
}

func NewHTTPRouter(router http.IRouter, authzMiddleware middlewares.IAuthzMiddleware,
	workspaceHandler *workspace.Handler, repositoryHandler *repository.Handler) IRouter {
	httpRoutes := &Router{
		IRouter:           router,
		IAuthzMiddleware:  authzMiddleware,
		ISwagger:          swagger.NewSwagger(router.GetMux(), router.GetPort()),
		workspaceHandler:  workspaceHandler,
		repositoryHandler: repositoryHandler,
	}

	return httpRoutes.setRoutes()
}

func (r *Router) setRoutes() IRouter {
	r.swaggerRoutes()
	r.workspaceRoutes()
	r.repositoryRoutes()

	return r
}

func (r *Router) workspaceRoutes() {
	r.Route(routes.WorkspaceHandler, func(router chi.Router) {
		router.Get("/", r.workspaceHandler.List)
		router.With(r.IsApplicationAdmin).Post("/", r.workspaceHandler.Create)
		router.With(r.IsWorkspaceMember).Get("/{workspaceID}", r.workspaceHandler.Get)
		router.With(r.IsWorkspaceAdmin).Patch("/{workspaceID}", r.workspaceHandler.Update)
		router.With(r.IsWorkspaceAdmin).Delete("/{workspaceID}", r.workspaceHandler.Delete)
		router.With(r.IsWorkspaceAdmin).Get("/{workspaceID}/roles", r.workspaceHandler.GetUsers)
		router.With(r.IsWorkspaceAdmin).Patch("/{workspaceID}/roles/{accountID}", r.workspaceHandler.UpdateRole)
		router.With(r.IsWorkspaceAdmin).Post("/{workspaceID}/roles", r.workspaceHandler.InviteUser)
		router.With(r.IsWorkspaceAdmin).Delete("/{workspaceID}/roles/{accountID}", r.workspaceHandler.RemoveUser)
	})
}

func (r *Router) repositoryRoutes() {
	r.Route(routes.RepositoryHandler, func(router chi.Router) {
		router.With(r.IsWorkspaceAdmin).Post("/", r.repositoryHandler.Create)
		router.With(r.IsWorkspaceMember).Get("/", r.repositoryHandler.List)
		router.With(r.IsRepositoryMember).Get("/{repositoryID}", r.repositoryHandler.Get)
		router.With(r.IsRepositoryAdmin).Patch("/{repositoryID}", r.repositoryHandler.Update)
		router.With(r.IsRepositoryAdmin).Delete("/{repositoryID}", r.repositoryHandler.Delete)
		router.With(r.IsRepositoryAdmin).Post("/{repositoryID}/roles", r.repositoryHandler.InviteUser)
		router.With(r.IsRepositoryAdmin).Patch("/{repositoryID}/roles/{accountID}", r.repositoryHandler.UpdateRole)
	})
}

func (r *Router) swaggerRoutes() {
	r.SetupSwagger()

	docs.SwaggerInfo.Host = r.GetSwaggerHost()
}
