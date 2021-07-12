// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package router

import (
	"github.com/go-chi/chi"

	httpRouter "github.com/ZupIT/horusec-devkit/pkg/services/http/router"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"
	"github.com/ZupIT/horusec-devkit/pkg/services/swagger"

	"github.com/ZupIT/horusec-platform/core/docs"
	"github.com/ZupIT/horusec-platform/core/internal/enums/routes"
	"github.com/ZupIT/horusec-platform/core/internal/handlers/health"
	"github.com/ZupIT/horusec-platform/core/internal/handlers/repository"
	"github.com/ZupIT/horusec-platform/core/internal/handlers/workspace"
)

type IRouter interface {
	httpRouter.IRouter
}

type Router struct {
	httpRouter.IRouter
	middlewares.IAuthzMiddleware
	workspaceHandler  *workspace.Handler
	repositoryHandler *repository.Handler
	healthHandler     *health.Handler
	swagger.ISwagger
}

func NewHTTPRouter(router httpRouter.IRouter, authzMiddleware middlewares.IAuthzMiddleware,
	workspaceHandler *workspace.Handler, repositoryHandler *repository.Handler, healthHandler *health.Handler) IRouter {
	httpRoutes := &Router{
		IRouter:           router,
		IAuthzMiddleware:  authzMiddleware,
		ISwagger:          swagger.NewSwagger(router.GetMux(), router.GetPort()),
		workspaceHandler:  workspaceHandler,
		repositoryHandler: repositoryHandler,
		healthHandler:     healthHandler,
	}

	return httpRoutes.setRoutes()
}

func (r *Router) setRoutes() IRouter {
	r.swaggerRoutes()
	r.workspaceRoutes()
	r.repositoryRoutes()
	r.healthRoutes()

	return r
}

func (r *Router) workspaceRoutes() {
	r.Route(routes.WorkspaceHandler, func(router chi.Router) {
		router.Get("/", r.workspaceHandler.List)
		router.Options("/", r.workspaceHandler.Options)
		router.With(r.IsApplicationAdmin).Post("/", r.workspaceHandler.Create)
		router.With(r.IsWorkspaceMember).Get("/{workspaceID}", r.workspaceHandler.Get)
		router.With(r.IsWorkspaceAdmin).Patch("/{workspaceID}", r.workspaceHandler.Update)
		router.With(r.IsWorkspaceAdmin).Delete("/{workspaceID}", r.workspaceHandler.Delete)
		router.With(r.IsWorkspaceAdmin).Get("/{workspaceID}/roles", r.workspaceHandler.GetUsers)
		router.With(r.IsWorkspaceAdmin).Patch("/{workspaceID}/roles/{accountID}", r.workspaceHandler.UpdateRole)
		router.With(r.IsWorkspaceAdmin).Post("/{workspaceID}/roles", r.workspaceHandler.InviteUser)
		router.With(r.IsWorkspaceAdmin).Delete("/{workspaceID}/roles/{accountID}", r.workspaceHandler.RemoveUser)
		router.With(r.IsWorkspaceAdmin).Post("/{workspaceID}/tokens", r.workspaceHandler.CreateToken)
		router.With(r.IsWorkspaceAdmin).Delete("/{workspaceID}/tokens/{tokenID}", r.workspaceHandler.DeleteToken)
		router.With(r.IsWorkspaceAdmin).Get("/{workspaceID}/tokens", r.workspaceHandler.ListTokens)
	})
}

func (r *Router) repositoryRoutes() {
	r.Route(routes.RepositoryHandler, func(router chi.Router) {
		router.Options("/", r.repositoryHandler.Options)
		router.With(r.IsWorkspaceAdmin).Post("/", r.repositoryHandler.Create)
		router.With(r.IsWorkspaceMember).Get("/", r.repositoryHandler.List)
		router.With(r.IsRepositoryMember).Get("/{repositoryID}", r.repositoryHandler.Get)
		router.With(r.IsRepositoryAdmin).Patch("/{repositoryID}", r.repositoryHandler.Update)
		router.With(r.IsRepositoryAdmin).Delete("/{repositoryID}", r.repositoryHandler.Delete)
		router.With(r.IsRepositoryAdmin).Post("/{repositoryID}/roles", r.repositoryHandler.InviteUser)
		router.With(r.IsRepositoryAdmin).Patch("/{repositoryID}/roles/{accountID}", r.repositoryHandler.UpdateRole)
		router.With(r.IsRepositoryAdmin).Get("/{repositoryID}/roles", r.repositoryHandler.GetUsers)
		router.With(r.IsRepositoryAdmin).Delete("/{repositoryID}/roles/{accountID}", r.repositoryHandler.RemoveUser)
		router.With(r.IsRepositoryAdmin).Post("/{repositoryID}/tokens", r.repositoryHandler.CreateToken)
		router.With(r.IsRepositoryAdmin).Delete("/{repositoryID}/tokens/{tokenID}", r.repositoryHandler.DeleteToken)
		router.With(r.IsRepositoryAdmin).Get("/{repositoryID}/tokens", r.repositoryHandler.ListTokens)
	})
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
