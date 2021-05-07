package account

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"

	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"
	_ "github.com/ZupIT/horusec-devkit/pkg/utils/http/entities" // swagger import
	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt/enums"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	accountController "github.com/ZupIT/horusec-platform/auth/internal/controllers/account"
	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
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

func (h *Handler) Options(w http.ResponseWriter, _ *http.Request) {
	httpUtil.StatusNoContent(w)
}

// @Tags Account
// @Description Create account when keycloak auth
// @ID create-account-keycloak
// @Accept  json
// @Produce  json
// @Param AccessToken body accountEntities.AccessToken true "create account with keycloak data"
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /auth/account/create-account-keycloak [post]
// @Security ApiKeyAuth
func (h *Handler) CreateAccountKeycloak(w http.ResponseWriter, r *http.Request) {
	data, err := h.useCases.AccessTokenFromIOReadCloser(r.Body)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	response, err := h.controller.CreateAccountKeycloak(data.AccessToken)
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

// @Tags Account
// @Description Create account when horusec auth
// @ID create-account-horusec
// @Accept  json
// @Produce  json
// @Param AccountData body accountEntities.Data true "create account with horusec data"
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /auth/account/create-account-horusec [post]
// @Security ApiKeyAuth
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

// @Tags Account
// @Description Validate account email
// @ID validate-account-email
// @Accept  json
// @Produce  json
// @Param accountID path string true "ID of the account"
// @Success 304 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /auth/account/validate/{accountID} [get]
// @Security ApiKeyAuth
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

// @Tags Account
// @Description Send a reset password code email
// @ID reset-password-code
// @Accept  json
// @Produce  json
// @Param Email body accountEntities.Email true "email data"
// @Success 204 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /auth/account/send-reset-code [post]
// @Security ApiKeyAuth
func (h *Handler) SendResetPasswordCode(w http.ResponseWriter, r *http.Request) {
	data, err := h.useCases.EmailFromIOReadCloser(r.Body)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	if err := h.controller.SendResetPasswordCode(data.Email); err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusNoContent(w)
}

// @Tags Account
// @Description Check for valid reset password code
// @ID check-reset-password-code
// @Accept  json
// @Produce  json
// @Param ResetCodeData body accountEntities.ResetCodeData true "reset password code"
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 403 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /auth/account/check-reset-code [post]
// @Security ApiKeyAuth
func (h *Handler) CheckResetPasswordCode(w http.ResponseWriter, r *http.Request) {
	data, err := h.useCases.ResetCodeDataFromIOReadCloser(r.Body)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	token, err := h.controller.CheckResetPasswordCode(data)
	if err != nil {
		h.checkResetPasswordCodeError(w, err)
		return
	}

	httpUtil.StatusOK(w, token)
}

func (h *Handler) checkResetPasswordCodeError(w http.ResponseWriter, err error) {
	if err == accountEnums.ErrorIncorrectRetrievePasswordCode {
		httpUtil.StatusForbidden(w, err)
		return
	}

	httpUtil.StatusInternalServerError(w, err)
}

// @Tags Account
// @Description Update account password
// @ID change-password
// @Accept  json
// @Produce  json
// @Param ChangePasswordData body accountEntities.ChangePasswordData true "change password data"
// @Success 204 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /auth/account/change-password [post]
// @Security ApiKeyAuth
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

// @Tags Account
// @Description Refresh access token
// @ID refresh-token
// @Accept  json
// @Produce  json
// @Param RefreshToken body accountEntities.RefreshToken true "refresh token data"
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Router /auth/account/refresh-token [post]
// @Security ApiKeyAuth
func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	data, err := h.useCases.RefreshTokenFromIOReadCloser(r.Body)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	response, err := h.controller.RefreshToken(data.RefreshToken)
	if err != nil {
		logger.LogError(accountEnums.MessageFailedToRefreshToken, err)
		httpUtil.StatusUnauthorized(w, accountEnums.ErrorInvalidOrExpiredToken)
		return
	}

	httpUtil.StatusOK(w, response)
}

// @Tags Account
// @Description Refresh access token
// @ID logout
// @Accept  json
// @Produce  json
// @Param RefreshToken body accountEntities.RefreshToken true "refresh token data"
// @Success 204 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Router /auth/account/logout [post]
// @Security ApiKeyAuth
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	data, err := h.useCases.RefreshTokenFromIOReadCloser(r.Body)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	h.controller.Logout(data.RefreshToken)
	httpUtil.StatusNoContent(w)
}

// @Tags Account
// @Description Verify username or email already in use
// @ID verify-email-username
// @Accept  json
// @Produce  json
// @Param CheckEmailAndUsername body accountEntities.CheckEmailAndUsername true "check already in use email username"
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /auth/account/verify-already-used [post]
// @Security ApiKeyAuth
func (h *Handler) CheckExistingEmailOrUsername(w http.ResponseWriter, r *http.Request) {
	data, err := h.useCases.CheckEmailAndUsernameFromIOReadCloser(r.Body)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	if err := h.controller.CheckExistingEmailOrUsername(data); err != nil {
		h.checkVerifyAlreadyInUseErrors(w, err)
		return
	}

	httpUtil.StatusNoContent(w)
}

func (h *Handler) checkVerifyAlreadyInUseErrors(w http.ResponseWriter, err error) {
	if err == accountEnums.ErrorEmailAlreadyInUse || err == accountEnums.ErrorUsernameAlreadyInUse {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	httpUtil.StatusInternalServerError(w, err)
}

// @Tags Account
// @Description Delete your account
// @ID delete-account
// @Accept  json
// @Produce  json
// @Success 204 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /auth/account/delete [delete]
// @Security ApiKeyAuth
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

// @Tags Account
// @Description Update account data
// @ID update-account
// @Accept  json
// @Produce  json
// @Param UpdateAccount body accountEntities.UpdateAccount true "update account data"
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /auth/account/update [patch]
// @Security ApiKeyAuth
func (h *Handler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	data, err := h.getUpdateAccountData(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	response, err := h.controller.UpdateAccount(data)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusOK(w, response)
}

func (h *Handler) getUpdateAccountData(r *http.Request) (*accountEntities.UpdateAccount, error) {
	accountID, err := h.controller.GetAccountID(r.Header.Get(enums.HorusecJWTHeader))
	if err != nil {
		return nil, err
	}

	data, err := h.useCases.UpdateAccountFromIOReadCloser(r.Body)
	if err != nil {
		return nil, err
	}

	return data.SetAccountIDAndIsConfirmed(accountID, h.appConfig.IsEmailsDisabled()), nil
}
