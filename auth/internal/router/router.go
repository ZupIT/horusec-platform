package router

import (
	"github.com/go-chi/chi"

	"github.com/ZupIT/horusec-devkit/pkg/services/http"
	"github.com/ZupIT/horusec-devkit/pkg/services/swagger"

	"github.com/ZupIT/horusec-platform/auth/config/grpc"
	"github.com/ZupIT/horusec-platform/auth/docs"
	"github.com/ZupIT/horusec-platform/auth/internal/enums/routes"
	authHandler "github.com/ZupIT/horusec-platform/auth/internal/handlers/authentication"
)

type IRouter interface {
	http.IRouter
}

type Router struct {
	http.IRouter
	swagger.ISwagger
	grpc.IAuthGRPCServer
	authHandler *authHandler.Handler
}

func NewHTTPRouter(router http.IRouter, authGRPCServer grpc.IAuthGRPCServer, handlerAuth *authHandler.Handler) IRouter {
	httpRouter := &Router{
		IRouter:         router,
		ISwagger:        swagger.NewSwagger(router.GetMux(), router.GetPort()),
		IAuthGRPCServer: authGRPCServer,
		authHandler:     handlerAuth,
	}

	httpRouter.startGRPCServer()
	return httpRouter.setRoutes()
}

func (r *Router) setRoutes() IRouter {
	r.swaggerRoutes()

	return r
}

func (r *Router) startGRPCServer() {
	go r.ListenAndServeGRPCServer()
}

func (r *Router) swaggerRoutes() {
	r.SetupSwagger()
	r.authenticationRoutes()

	docs.SwaggerInfo.Host = r.GetSwaggerHost()
}

func (r *Router) authenticationRoutes() {
	r.Route(routes.AuthenticationHandler, func(router chi.Router) {
		router.Post("/login", r.authHandler.Login)
	})
}
