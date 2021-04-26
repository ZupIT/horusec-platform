package account

import (
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/enums/queues"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/cache"
	"github.com/ZupIT/horusec-devkit/pkg/utils/crypto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	accountEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/account"
	authEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication"
	accountRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/account"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/keycloak"
	accountUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/account"
)

type IController interface {
	CreateAccountKeycloak(token string) (*accountEntities.Response, error)
	CreateAccountHorusec(data *accountEntities.Data) (*accountEntities.Response, error)
	ValidateAccountEmail(accountID uuid.UUID) error
	SendResetPasswordCode(email string) error
	CheckResetPasswordCode(data *accountEntities.ResetCodeData) (string, error)
	ChangePassword(data *accountEntities.ChangePasswordData) error
	RefreshToken(refreshToken string) (*authEntities.LoginResponse, error)
	Logout(refreshToken string)
	CheckExistingEmailOrUsername(data *accountEntities.CheckEmailAndUsername) error
	DeleteAccount(accountID uuid.UUID) error
	GetAccountID(token string) (uuid.UUID, error)
	UpdateAccount(data *accountEntities.UpdateAccount) (*accountEntities.Response, error)
}

type Controller struct {
	keycloakAuth      keycloak.IService
	accountRepository accountRepository.IRepository
	accountUseCases   accountUseCases.IUseCases
	appConfig         app.IConfig
	broker            broker.IBroker
	cache             cache.ICache
}

func NewAccountController(repositoryAccount accountRepository.IRepository, keycloakAuth keycloak.IService,
	useCasesAccount accountUseCases.IUseCases, appConfig app.IConfig, brokerLib broker.IBroker,
	cacheLib cache.ICache) IController {
	return &Controller{
		accountRepository: repositoryAccount,
		keycloakAuth:      keycloakAuth,
		appConfig:         appConfig,
		accountUseCases:   useCasesAccount,
		broker:            brokerLib,
		cache:             cacheLib,
	}
}

func (c *Controller) CreateAccountKeycloak(token string) (*accountEntities.Response, error) {
	userInfo, err := c.keycloakAuth.GetUserInfo(token)
	if err != nil {
		return nil, err
	}

	account, err := c.accountRepository.CreateAccount(c.accountUseCases.NewAccountFromKeycloakUserInfo(userInfo))
	if err != nil {
		return c.accountUseCases.NewAccountFromKeycloakUserInfo(userInfo).ToResponse(),
			c.accountUseCases.CheckCreateAccountErrors(err)
	}

	return account.ToResponse(), nil
}

func (c *Controller) CreateAccountHorusec(data *accountEntities.Data) (*accountEntities.Response, error) {
	account := data.ToAccount()
	if c.appConfig.IsBrokerDisabled() {
		_ = account.SetIsConfirmedTrue()
	}

	if _, err := c.accountRepository.CreateAccount(account); err != nil {
		return nil, c.accountUseCases.CheckCreateAccountErrors(err)
	}

	return account.ToResponse(), c.sendValidateAccountEmail(account)
}

func (c *Controller) sendValidateAccountEmail(account *accountEntities.Account) error {
	if c.appConfig.IsBrokerDisabled() {
		return nil
	}

	return c.broker.Publish(queues.HorusecEmail.ToString(), "", "",
		c.accountUseCases.NewAccountValidationEmail(account))
}

func (c *Controller) ValidateAccountEmail(accountID uuid.UUID) error {
	account, err := c.accountRepository.GetAccount(accountID)
	if err != nil {
		return err
	}

	_, err = c.accountRepository.Update(account.SetIsConfirmedTrue())
	return err
}

func (c *Controller) SendResetPasswordCode(email string) error {
	account, err := c.accountRepository.GetAccountByEmail(email)
	if err != nil {
		return err
	}

	code := c.accountUseCases.GenerateResetPasswordCode()

	c.cache.Set(account.Email, code, accountEnums.ResetPasswordCodeDuration)

	return c.sendResetPasswordCodeEmail(account, code)
}

func (c *Controller) sendResetPasswordCodeEmail(account *accountEntities.Account, code string) error {
	if c.appConfig.IsBrokerDisabled() {
		return nil
	}

	return c.broker.Publish(queues.HorusecEmail.ToString(), "", "",
		c.accountUseCases.NewResetPasswordCodeEmail(account, code))
}

func (c *Controller) CheckResetPasswordCode(data *accountEntities.ResetCodeData) (string, error) {
	correctCode, err := c.cache.GetString(data.Email)
	if err != nil {
		return "", err
	}

	if data.Code != correctCode {
		return "", accountEnums.ErrorIncorrectRetrievePasswordCode
	}

	return c.createAccessTokenIfCorrectCode(data)
}

func (c *Controller) createAccessTokenIfCorrectCode(data *accountEntities.ResetCodeData) (string, error) {
	account, err := c.accountRepository.GetAccountByEmail(data.Email)
	if err != nil {
		return "", err
	}

	token, _, err := jwt.CreateToken(account.ToTokenData(), nil)
	if err != nil {
		return "", err
	}

	c.cache.Delete(data.Email)
	return token, nil
}

func (c *Controller) ChangePassword(data *accountEntities.ChangePasswordData) error {
	account, err := c.accountRepository.GetAccount(data.AccountID)
	if err != nil {
		return err
	}

	if crypto.CheckPasswordHashBcrypt(data.Password, account.Password) {
		return accountEnums.ErrorPasswordEqualPrevious
	}

	_, err = c.accountRepository.Update(account.SetNewPassword(data.Password))
	return err
}

func (c *Controller) RefreshToken(refreshToken string) (*authEntities.LoginResponse, error) {
	accountID, err := c.cache.GetString(refreshToken)
	if err != nil {
		return nil, err
	}

	account, err := c.accountRepository.GetAccount(parser.ParseStringToUUID(accountID))
	if err != nil {
		return nil, err
	}

	c.cache.Delete(refreshToken)
	return c.createNewTokens(account)
}

func (c *Controller) createNewTokens(account *accountEntities.Account) (*authEntities.LoginResponse, error) {
	accessToken, expiresAt, err := jwt.CreateToken(account.ToTokenData(), nil)
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.CreateRefreshToken()
	c.cache.Set(refreshToken, account.AccountID, authEnums.TokenDuration)
	return account.ToLoginResponse(accessToken, refreshToken, expiresAt), nil
}

func (c *Controller) Logout(refreshToken string) {
	c.cache.Delete(refreshToken)
}

func (c *Controller) CheckExistingEmailOrUsername(data *accountEntities.CheckEmailAndUsername) error {
	validateEmail, _ := c.accountRepository.GetAccountByEmail(data.Email)
	if validateEmail != nil && validateEmail.Email != "" {
		return accountEnums.ErrorEmailAlreadyInUse
	}

	validateUsername, _ := c.accountRepository.GetAccountByUsername(data.Username)
	if validateUsername != nil && validateUsername.Username != "" {
		return accountEnums.ErrorUsernameAlreadyInUse
	}

	return nil
}

func (c *Controller) DeleteAccount(accountID uuid.UUID) error {
	return c.accountRepository.Delete(accountID)
}

func (c *Controller) GetAccountID(token string) (uuid.UUID, error) {
	switch c.appConfig.GetAuthenticationType() {
	case auth.Horusec:
		return jwt.GetAccountIDByJWTToken(token)
	case auth.Keycloak:
		return c.getAccountIDKeycloak(token)
	case auth.Ldap:
		return jwt.GetAccountIDByJWTToken(token)
	}

	return uuid.Nil, authEnums.ErrorAuthTypeInvalid
}

func (c *Controller) getAccountIDKeycloak(token string) (uuid.UUID, error) {
	data, err := c.keycloakAuth.GetAccountDataFromToken(token)
	if err != nil {
		return uuid.Nil, err
	}

	return parser.ParseStringToUUID(data.AccountID), nil
}

func (c *Controller) UpdateAccount(data *accountEntities.UpdateAccount) (*accountEntities.Response, error) {
	account, err := c.accountRepository.GetAccount(data.AccountID)
	if err != nil {
		return nil, err
	}

	if errMail := c.checkForEmailChange(data, account); errMail != nil {
		return nil, errMail
	}

	account, err = c.accountRepository.Update(account)
	if err != nil {
		return nil, err
	}

	return account.ToResponse(), err
}

func (c *Controller) checkForEmailChange(data *accountEntities.UpdateAccount, account *accountEntities.Account) error {
	if data.HasEmailChange(account.Email) {
		account.UpdateFromUpdateAccountData(data)
		return c.sendValidateAccountEmail(account)
	}

	account.UpdateFromUpdateAccountData(data)
	return nil
}
