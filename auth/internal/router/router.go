package router

import (
	"github.com/go-chi/chi"

	"github.com/ZupIT/horusec-devkit/pkg/services/http"
	"github.com/ZupIT/horusec-devkit/pkg/services/swagger"

	"github.com/ZupIT/horusec-platform/auth/config/grpc"
	"github.com/ZupIT/horusec-platform/auth/docs"
	"github.com/ZupIT/horusec-platform/auth/internal/enums/routes"
	accountHandler "github.com/ZupIT/horusec-platform/auth/internal/handlers/account"
	authHandler "github.com/ZupIT/horusec-platform/auth/internal/handlers/authentication"
)

type IRouter interface {
	http.IRouter
}

type Router struct {
	http.IRouter
	swagger.ISwagger
	grpc.IAuthGRPCServer
	authHandler    *authHandler.Handler
	accountHandler *accountHandler.Handler
}

func NewHTTPRouter(router http.IRouter, authGRPCServer grpc.IAuthGRPCServer, handlerAuth *authHandler.Handler,
	handlerAccount *accountHandler.Handler) IRouter {
	httpRouter := &Router{
		IRouter:         router,
		ISwagger:        swagger.NewSwagger(router.GetMux(), router.GetPort()),
		IAuthGRPCServer: authGRPCServer,
		authHandler:     handlerAuth,
		accountHandler:  handlerAccount,
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
	r.accountRoutes()

	docs.SwaggerInfo.Host = r.GetSwaggerHost()
}

func (r *Router) authenticationRoutes() {
	r.Route(routes.AuthenticationHandler, func(router chi.Router) {
		router.Post("/login", r.authHandler.Login)
		router.Get("/config", r.authHandler.GetConfig)
	})
}

func (r *Router) accountRoutes() {
	r.Route(routes.AccountHandler, func(router chi.Router) {
		router.Post("/create-account-keycloak", r.accountHandler.CreateAccountKeycloak)
		router.Post("/create-account-horusec", r.accountHandler.CreateAccountHorusec)
		router.Get("/validate/{accountID}", r.accountHandler.ValidateAccountEmail)
		router.Post("/send-reset-code", r.accountHandler.SendResetPasswordCode)
		router.Post("/check-reset-code", r.accountHandler.CheckResetPasswordCode)
		router.Post("/change-password", r.accountHandler.ChangePassword)
		router.Post("/refresh-token", r.accountHandler.RefreshToken)
		router.Post("/logout", r.accountHandler.Logout)
		router.Delete("/delete", r.accountHandler.DeleteAccount)
		router.Post("/verify-already-used", r.accountHandler.CheckExistingEmailOrUsername)
		router.Patch("/update", r.accountHandler.UpdateAccount)
	})
}
