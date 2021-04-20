package account

import (
	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	accountRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/account"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/keycloak"
	accountUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/account"
)

type IController interface {
	CreateAccountKeycloak(token string) (*accountEntities.Response, error)
}

type Controller struct {
	keycloakAuth      keycloak.IService
	accountRepository accountRepository.IRepository
	accountUseCases   accountUseCases.IUseCases
}

func NewAccountController(repositoryAccount accountRepository.IRepository) IController {
	return &Controller{
		accountRepository: repositoryAccount,
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
