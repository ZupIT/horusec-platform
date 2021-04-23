package account

import (
	"net/http"

	"github.com/ZupIT/horusec-platform/auth/config/app"

	"github.com/google/uuid"

	"github.com/go-chi/chi"

	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"

	accountController "github.com/ZupIT/horusec-platform/auth/internal/controllers/account"
	accountEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/account"
	accountUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/account"
)

type Handler struct {
	useCases   accountUseCases.IUseCases
	controller accountController.IController
	appConfig  app.IConfig
}

func NewAccountHandler(useCases accountUseCases.IUseCases, controller accountController.IController,
	appConfig app.IConfig) *Handler {
	return &Handler{
		useCases:   useCases,
		controller: controller,
		appConfig:  appConfig,
	}
}

func (h *Handler) CreateAccountKeycloak(w http.ResponseWriter, r *http.Request) {
	keyCloakToken, err := h.useCases.AccessTokenFromIOReadCloser(r.Body)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	response, err := h.controller.CreateAccountKeycloak(keyCloakToken)
	if err != nil {
		h.checkCreateAccountKeycloakErrors(w, err, response)
		return
	}

	httpUtil.StatusOK(w, response)
}

func (h *Handler) checkCreateAccountKeycloakErrors(w http.ResponseWriter, err error, response interface{}) {
	if err == accountEnums.ErrorEmailAlreadyInUse || err == accountEnums.ErrorUsernameAlreadyInUse {
		httpUtil.StatusOK(w, response)
		return
	}

	httpUtil.StatusInternalServerError(w, err)
}

func (h *Handler) CreateAccountHorusec(w http.ResponseWriter, r *http.Request) {
	data, err := h.useCases.AccountDataFromIOReadCloser(r.Body)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	response, err := h.controller.CreateAccountHorusec(data)
	if err != nil {
		h.checkCreateAccountHorusecErrors(w, err)
		return
	}

	httpUtil.StatusOK(w, response)
}

func (h *Handler) checkCreateAccountHorusecErrors(w http.ResponseWriter, err error) {
	if err == accountEnums.ErrorEmailAlreadyInUse || err == accountEnums.ErrorUsernameAlreadyInUse {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	httpUtil.StatusInternalServerError(w, err)
}

func (h *Handler) ValidateAccountEmail(w http.ResponseWriter, r *http.Request) {
	accountID, err := uuid.Parse(chi.URLParam(r, accountEnums.ID))
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	if h.controller.ValidateAccountEmail(accountID) != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	http.Redirect(w, r, h.appConfig.GetHorusecManagerURL(), http.StatusSeeOther)
}

func (h *Handler) SendResetPasswordCode(w http.ResponseWriter, r *http.Request) {
	email, err := h.useCases.EmailFromIOReadCloser(r.Body)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	err = h.controller.SendResetPasswordCode(email)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusNoContent(w)
}
