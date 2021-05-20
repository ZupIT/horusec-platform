package account

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	accountController "github.com/ZupIT/horusec-platform/auth/internal/controllers/account"
	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	"github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	accountEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/account"
	accountUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/account"
)

func getAppConfig() app.IConfig {
	databaseMock := &database.Mock{}
	databaseMock.On("Create").Return(&response.Response{})

	return app.NewAuthAppConfig(&database.Connection{Read: databaseMock, Write: databaseMock})
}

func TestNewAccountHandler(t *testing.T) {
	t.Run("should success create a new handler", func(t *testing.T) {
		assert.NotNil(t, NewAccountHandler(nil, nil, nil))
	})
}

func TestCreateAccountKeycloak(t *testing.T) {
	t.Run("should return 200 when success create account", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("CreateAccountKeycloak").Return(&accountEntities.Response{}, nil)

		data := &accountEntities.AccessToken{AccessToken: "test"}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.CreateAccountKeycloak(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 200 when account already exists", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("CreateAccountKeycloak").Return(
			&accountEntities.Response{}, accountEnums.ErrorEmailAlreadyInUse)

		data := &accountEntities.AccessToken{AccessToken: "test"}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.CreateAccountKeycloak(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("CreateAccountKeycloak").Return(
			&accountEntities.Response{}, errors.New("test"))

		data := &accountEntities.AccessToken{AccessToken: "test"}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.CreateAccountKeycloak(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid request body", func(t *testing.T) {
		appConfig := getAppConfig()
		controllerMock := &accountController.Mock{}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader([]byte("")))
		w := httptest.NewRecorder()

		handler.CreateAccountKeycloak(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestCreateAccountHorusec(t *testing.T) {
	t.Run("should return 200 when success create account", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("CreateAccountHorusec").Return(&accountEntities.Response{}, nil)

		data := &accountEntities.Data{
			Email:    "test@test.com",
			Password: "Test@123",
			Username: "test",
		}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.CreateAccountHorusec(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 400 when username or email already in use", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("CreateAccountHorusec").Return(
			&accountEntities.Response{}, accountEnums.ErrorEmailAlreadyInUse)

		data := &accountEntities.Data{
			Email:    "test@test.com",
			Password: "Test@123",
			Username: "test",
		}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.CreateAccountHorusec(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("CreateAccountHorusec").Return(
			&accountEntities.Response{}, errors.New("test"))

		data := &accountEntities.Data{
			Email:    "test@test.com",
			Password: "Test@123",
			Username: "test",
		}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.CreateAccountHorusec(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid request body", func(t *testing.T) {
		appConfig := getAppConfig()
		controllerMock := &accountController.Mock{}

		data := &accountEntities.Data{}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.CreateAccountHorusec(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestValidateAccountEmail(t *testing.T) {
	t.Run("should return 303 when success validate account email", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("ValidateAccountEmail").Return(nil)

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("accountID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.ValidateAccountEmail(w, r)

		assert.Equal(t, http.StatusSeeOther, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("ValidateAccountEmail").Return(errors.New("test"))

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("accountID", uuid.NewString())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.ValidateAccountEmail(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid account id", func(t *testing.T) {
		appConfig := getAppConfig()
		controllerMock := &accountController.Mock{}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		handler.ValidateAccountEmail(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestSendResetPasswordCode(t *testing.T) {
	t.Run("should return 204 when success sent code", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("SendResetPasswordCode").Return(nil)

		data := &accountEntities.Email{
			Email: "test@test.com",
		}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.SendResetPasswordCode(w, r)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("SendResetPasswordCode").Return(errors.New("test"))

		data := &accountEntities.Email{
			Email: "test@test.com",
		}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.SendResetPasswordCode(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid request body", func(t *testing.T) {
		appConfig := getAppConfig()
		controllerMock := &accountController.Mock{}

		data := &accountEntities.Email{}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.SendResetPasswordCode(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestCheckResetPasswordCode(t *testing.T) {
	t.Run("should return 200 when valid code", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("CheckResetPasswordCode").Return("test", nil)

		data := &accountEntities.ResetCodeData{
			Email: "test@test.com",
			Code:  "123456",
		}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.CheckResetPasswordCode(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 403 when wrong code", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("CheckResetPasswordCode").Return(
			"", accountEnums.ErrorIncorrectRetrievePasswordCode)

		data := &accountEntities.ResetCodeData{
			Email: "test@test.com",
			Code:  "123456",
		}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.CheckResetPasswordCode(w, r)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("CheckResetPasswordCode").Return("", errors.New("test"))

		data := &accountEntities.ResetCodeData{
			Email: "test@test.com",
			Code:  "123456",
		}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.CheckResetPasswordCode(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid request body", func(t *testing.T) {
		appConfig := getAppConfig()
		controllerMock := &accountController.Mock{}

		data := &accountEntities.ResetCodeData{}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.CheckResetPasswordCode(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestChangePassword(t *testing.T) {
	t.Run("should return 204 when success change password", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("ChangePassword").Return(nil)
		controllerMock.On("GetAccountID").Return(uuid.New(), nil)

		data := &accountEntities.ChangePasswordData{
			Password: "Test@123",
		}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.ChangePassword(w, r)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("should return 400 when password equal previous one", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("ChangePassword").Return(accountEnums.ErrorPasswordEqualPrevious)
		controllerMock.On("GetAccountID").Return(uuid.New(), nil)

		data := &accountEntities.ChangePasswordData{
			Password: "Test@123",
		}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.ChangePassword(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("ChangePassword").Return(errors.New("test"))
		controllerMock.On("GetAccountID").Return(uuid.New(), nil)

		data := &accountEntities.ChangePasswordData{
			Password: "Test@123",
		}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.ChangePassword(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid request body", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("GetAccountID").Return(uuid.New(), nil)

		data := &accountEntities.ChangePasswordData{}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.ChangePassword(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when failed to get account id", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("GetAccountID").Return(uuid.New(), errors.New("test"))

		data := &accountEntities.ChangePasswordData{}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.ChangePassword(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestRefreshToken(t *testing.T) {
	t.Run("should return 200 when success set refresh token", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("RefreshToken").Return(&authentication.LoginResponse{}, nil)

		data := &accountEntities.RefreshToken{
			RefreshToken: "test",
		}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.RefreshToken(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 401 when failed to refresh token", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("RefreshToken").Return(&authentication.LoginResponse{}, errors.New("test"))

		data := &accountEntities.RefreshToken{
			RefreshToken: "test",
		}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.RefreshToken(w, r)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("should return 400 when invalid request body", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("RefreshToken").Return(&authentication.LoginResponse{}, errors.New("test"))

		data := &accountEntities.RefreshToken{}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.RefreshToken(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestLogout(t *testing.T) {
	t.Run("should return 204 when success logout", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("Logout")

		data := &accountEntities.RefreshToken{
			RefreshToken: "test",
		}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.Logout(w, r)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("should return 400 when invalid request body", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("Logout")

		data := &accountEntities.RefreshToken{}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.Logout(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestCheckExistingEmailOrUsername(t *testing.T) {
	t.Run("should return 204 when not in use", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("CheckExistingEmailOrUsername").Return(nil)

		data := &accountEntities.CheckEmailAndUsername{
			Email:    "test@test.com",
			Username: "test",
		}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.CheckExistingEmailOrUsername(w, r)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("should return 400 when email or username already in use", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("CheckExistingEmailOrUsername").Return(accountEnums.ErrorEmailAlreadyInUse)

		data := &accountEntities.CheckEmailAndUsername{
			Email:    "test@test.com",
			Username: "test",
		}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.CheckExistingEmailOrUsername(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("CheckExistingEmailOrUsername").Return(errors.New("test"))

		data := &accountEntities.CheckEmailAndUsername{
			Email:    "test@test.com",
			Username: "test",
		}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.CheckExistingEmailOrUsername(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid request body", func(t *testing.T) {
		appConfig := getAppConfig()
		controllerMock := &accountController.Mock{}

		data := &accountEntities.CheckEmailAndUsername{}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.CheckExistingEmailOrUsername(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestDeleteAccount(t *testing.T) {
	t.Run("should return 204 success delete account", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("GetAccountID").Return(uuid.New(), nil)
		controllerMock.On("DeleteAccount").Return(nil)

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodDelete, "test", nil)
		w := httptest.NewRecorder()

		handler.DeleteAccount(w, r)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("GetAccountID").Return(uuid.New(), nil)
		controllerMock.On("DeleteAccount").Return(errors.New("test"))

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodDelete, "test", nil)
		w := httptest.NewRecorder()

		handler.DeleteAccount(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 401 when failed to get account id", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("GetAccountID").Return(uuid.New(), errors.New("test"))

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodDelete, "test", nil)
		w := httptest.NewRecorder()

		handler.DeleteAccount(w, r)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

func TestUpdateAccount(t *testing.T) {
	t.Run("should return 200 when success update account", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("UpdateAccount").Return(&accountEntities.Response{}, nil)
		controllerMock.On("GetAccountID").Return(uuid.New(), nil)

		data := &accountEntities.UpdateAccount{
			Email:    "test@test.com",
			Username: "test",
		}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.UpdateAccount(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when something went wrong", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("UpdateAccount").Return(&accountEntities.Response{}, errors.New("test"))
		controllerMock.On("GetAccountID").Return(uuid.New(), nil)

		data := &accountEntities.UpdateAccount{
			Email:    "test@test.com",
			Username: "test",
		}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.UpdateAccount(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid request body", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("GetAccountID").Return(uuid.New(), nil)

		data := &accountEntities.UpdateAccount{}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.UpdateAccount(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when failed to get account id", func(t *testing.T) {
		appConfig := getAppConfig()

		controllerMock := &accountController.Mock{}
		controllerMock.On("GetAccountID").Return(uuid.New(), errors.New("test"))

		data := &accountEntities.UpdateAccount{}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodPatch, "test", bytes.NewReader(data.ToBytes()))
		w := httptest.NewRecorder()

		handler.UpdateAccount(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestOptions(t *testing.T) {
	t.Run("should return 204 when options", func(t *testing.T) {
		appConfig := getAppConfig()
		controllerMock := &accountController.Mock{}

		handler := NewAccountHandler(accountUseCases.NewAccountUseCases(appConfig), controllerMock, appConfig)

		r, _ := http.NewRequest(http.MethodOptions, "test", nil)
		w := httptest.NewRecorder()

		handler.Options(w, r)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}
