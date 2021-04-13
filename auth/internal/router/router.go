package router

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/http"
	"github.com/ZupIT/horusec-devkit/pkg/services/swagger"
	"github.com/ZupIT/horusec-platform/auth/config/grpc"
)

type IRouter interface {
	http.IRouter
}

type Router struct {
	http.IRouter
	swagger.ISwagger
	grpc.IAuthGRPCServer
}

func NewHTTPRouter(router http.IRouter, authGRPCServer grpc.IAuthGRPCServer) IRouter {
	httpRouter := &Router{
		IRouter:         router,
		ISwagger:        swagger.NewSwagger(router.GetMux(), router.GetPort()),
		IAuthGRPCServer: authGRPCServer,
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

	//docs.SwaggerInfo.Host = r.GetSwaggerHost()
}
