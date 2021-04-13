package authentication

import (
	"net/http"

	"github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"

	authTypes "github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	databaseEnums "github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	authController "github.com/ZupIT/horusec-platform/auth/internal/controllers/authentication"
	authEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication"
	accountUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/account"
)

type Handler struct {
	useCases   accountUseCases.IUseCases
	appConfig  app.IConfig
	controller authController.IController
}

func NewAuthenticationHandler() *Handler {
	return &Handler{}
}

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
		httpUtil.StatusInternalServerError(w, err)
	case authTypes.Keycloak:
		httpUtil.StatusInternalServerError(w, err)
	default:
		httpUtil.StatusInternalServerError(w, err)
	}
}

func (h *Handler) checkLoginErrorsHorusec(w http.ResponseWriter, err error) {
	if err == authEnums.ErrorWrongEmailOrPassword || err == databaseEnums.ErrorNotFoundRecords {
		httpUtil.StatusForbidden(w, authEnums.ErrorWrongEmailOrPassword)
		return
	}

	if err == authEnums.ErrorAccountEmailNotConfirmed {
		httpUtil.StatusForbidden(w, err)
		return
	}

	httpUtil.StatusInternalServerError(w, err)
}
