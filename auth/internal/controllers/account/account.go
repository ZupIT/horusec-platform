package account

import (
	"github.com/ZupIT/horusec-devkit/pkg/enums/queues"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/cache"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	accountEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/account"
	accountRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/account"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/keycloak"
	accountUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/account"
)

type IController interface {
	CreateAccountKeycloak(token string) (*accountEntities.Response, error)
	CreateAccountHorusec(data *accountEntities.Data) (*accountEntities.Response, error)
	ValidateAccountEmail(accountID uuid.UUID) error
	SendResetPasswordCode(email string) error
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
	if c.appConfig.IsDisableBroker() {
		account.SetIsConfirmedTrue()
	}

	if _, err := c.accountRepository.CreateAccount(account); err != nil {
		return nil, c.accountUseCases.CheckCreateAccountErrors(err)
	}

	return account.ToResponse(), c.sendValidateAccountEmail(account)
}

func (c *Controller) sendValidateAccountEmail(account *accountEntities.Account) error {
	if c.appConfig.IsDisableBroker() {
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

	_, err = c.accountRepository.Update(account.SetNewAccountData().Update())
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
	if c.appConfig.IsDisableBroker() {
		return nil
	}

	return c.broker.Publish(queues.HorusecEmail.ToString(), "", "",
		c.accountUseCases.NewResetPasswordCodeEmail(account, code))
}
