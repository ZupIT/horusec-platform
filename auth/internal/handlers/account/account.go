package account

import (
	"net/http"

	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"

	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt/enums"

	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"

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

	if err := h.controller.ValidateAccountEmail(accountID); err != nil {
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

	if err := h.controller.SendResetPasswordCode(email); err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusNoContent(w)
}

func (h *Handler) CheckResetPasswordCode(w http.ResponseWriter, r *http.Request) {
	data, err := h.useCases.ResetCodeDataFromIOReadCloser(r.Body)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	token, err := h.controller.CheckResetPasswordCode(data)
	if err != nil {
		h.CheckResetPasswordCodeError(w, err)
		return
	}

	httpUtil.StatusOK(w, token)
}

func (h *Handler) CheckResetPasswordCodeError(w http.ResponseWriter, err error) {
	if err == accountEnums.ErrorIncorrectRetrievePasswordCode {
		httpUtil.StatusForbidden(w, err)
		return
	}

	httpUtil.StatusInternalServerError(w, err)
}

func (h *Handler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	data, err := h.getChangePasswordData(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	if err := h.controller.ChangePassword(data); err != nil {
		h.checkChangePasswordDataErrors(w, err)
		return
	}

	httpUtil.StatusNoContent(w)
}

func (h *Handler) getChangePasswordData(r *http.Request) (*accountEntities.ChangePasswordData, error) {
	accountID, err := h.controller.GetAccountID(r.Header.Get(enums.HorusecJWTHeader))
	if err != nil {
		return nil, err
	}

	data, err := h.useCases.ChangePasswordDataFromIOReadCloser(r.Body)
	if err != nil {
		return nil, err
	}

	return data.SetAccountID(accountID), nil
}

func (h *Handler) checkChangePasswordDataErrors(w http.ResponseWriter, err error) {
	if err == accountEnums.ErrorPasswordEqualPrevious {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	httpUtil.StatusInternalServerError(w, err)
}

func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := h.useCases.RefreshTokenFromIOReadCloser(r.Body)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	response, err := h.controller.RefreshToken(refreshToken)
	if err != nil {
		logger.LogError(accountEnums.MessageFailedToRefreshToken, err)
		httpUtil.StatusUnauthorized(w, accountEnums.ErrorInvalidOrExpiredToken)
		return
	}

	httpUtil.StatusOK(w, response)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := h.useCases.RefreshTokenFromIOReadCloser(r.Body)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	h.controller.Logout(refreshToken)
	httpUtil.StatusNoContent(w)
}

func (h *Handler) VerifyAlreadyInUse(w http.ResponseWriter, r *http.Request) {
	data, err := h.useCases.CheckEmailAndUsernameFromIOReadCloser(r.Body)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	err = h.controller.CheckExistingEmailOrUsername(data)
	if err != nil {
		h.checkVerifyAlreadyInUseErrors(w, err)
		return
	}

	httpUtil.StatusOK(w, "")
}

func (h *Handler) checkVerifyAlreadyInUseErrors(w http.ResponseWriter, err error) {
	if err == accountEnums.ErrorEmailAlreadyInUse || err == accountEnums.ErrorUsernameAlreadyInUse {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	httpUtil.StatusInternalServerError(w, err)
}

func (h *Handler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	accountID, err := h.controller.GetAccountID(r.Header.Get(enums.HorusecJWTHeader))
	if err != nil {
		httpUtil.StatusUnauthorized(w, err)
		return
	}

	if err := h.controller.DeleteAccount(accountID); err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusNoContent(w)
}
