package authentication

import (
	"context"
	"net/http"

	authTypes "github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	databaseEnums "github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	authController "github.com/ZupIT/horusec-platform/auth/internal/controllers/authentication"
	"github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	horusecAuthEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication/horusec"
	ldapEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication/ldap"
	authUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/authentication"
)

type Handler struct {
	proto.UnimplementedAuthServiceServer
	useCases   authUseCases.IUseCases
	appConfig  app.IConfig
	controller authController.IController
}

func NewAuthenticationHandler(appConfig app.IConfig, useCases authUseCases.IUseCases,
	controller authController.IController) *Handler {
	return &Handler{
		useCases:   useCases,
		appConfig:  appConfig,
		controller: controller,
	}
}

// @Tags Authenticate
// @Description Get auth config
// @ID get-config
// @Accept  json
// @Produce  json
// @Success 200 {object} entities.Response
// @Router /auth/authenticate/config [get]
// @Security ApiKeyAuth
func (h *Handler) GetConfig(w http.ResponseWriter, _ *http.Request) {
	httpUtil.StatusOK(w, h.appConfig.ToConfigResponse())
}

// @Tags Authenticate
// @Description Login in into a horusec account
// @ID login
// @Accept  json
// @Produce  json
// @Param Credentials body authentication.LoginCredentials true "login data"
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 403 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /auth/authenticate/login [post]
// @Security ApiKeyAuth
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	credentials, err := h.getLoginCredentials(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	response, err := h.controller.Login(credentials)
	if err != nil {
		h.checkLoginErrors(w, err)
		return
	}

	httpUtil.StatusOK(w, response)
}

func (h *Handler) getLoginCredentials(r *http.Request) (*authentication.LoginCredentials, error) {
	credentials, err := h.useCases.LoginCredentialsFromIOReadCloser(r.Body)
	if err != nil {
		return credentials, err
	}

	return credentials, nil
}

func (h *Handler) checkLoginErrors(w http.ResponseWriter, err error) {
	switch h.appConfig.GetAuthType() {
	case authTypes.Horusec:
		h.checkLoginErrorsHorusec(w, err)
	case authTypes.Ldap:
		h.checkLoginErrorsLdap(w, err)
	case authTypes.Keycloak:
		httpUtil.StatusInternalServerError(w, err)
	default:
		httpUtil.StatusInternalServerError(w, err)
	}
}

func (h *Handler) checkLoginErrorsHorusec(w http.ResponseWriter, err error) {
	if err == horusecAuthEnums.ErrorWrongEmailOrPassword || err == databaseEnums.ErrorNotFoundRecords {
		httpUtil.StatusForbidden(w, horusecAuthEnums.ErrorWrongEmailOrPassword)
		return
	}

	if err == horusecAuthEnums.ErrorAccountEmailNotConfirmed {
		httpUtil.StatusForbidden(w, err)
		return
	}

	httpUtil.StatusInternalServerError(w, err)
}

func (h *Handler) checkLoginErrorsLdap(w http.ResponseWriter, err error) {
	if err == ldapEnums.ErrorUserDoesNotExist || err == ldapEnums.ErrorLdapUnauthorized {
		httpUtil.StatusForbidden(w, ldapEnums.ErrorWrongUserOrPassword)
		return
	}

	httpUtil.StatusInternalServerError(w, err)
}

func (h *Handler) IsAuthorized(_ context.Context, data *proto.IsAuthorizedData) (*proto.IsAuthorizedResponse, error) {
	isAuthorized, err := h.controller.IsAuthorized(h.useCases.NewAuthorizationDataFromGrpcData(data))

	return h.useCases.NewIsAuthorizedResponse(isAuthorized), err
}

func (h *Handler) GetAccountInfo(_ context.Context, data *proto.GetAccountData) (*proto.GetAccountDataResponse, error) {
	return h.controller.GetAccountInfo(data.Token)
}

func (h *Handler) GetAuthConfig(context.Context, *proto.GetAuthConfigData) (*proto.GetAuthConfigResponse, error) {
	return h.appConfig.ToGetAuthConfigResponse(), nil
}
