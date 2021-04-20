package account

import (
	keycloakEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication/keycloak"
	accountRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/account"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/keycloak"
	accountUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/account"
)

type IController interface {
	CreateAccountKeycloak(token string) error
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

func (c *Controller) CreateAccountKeycloak(token string) error {
	userInfo, err := c.keycloakAuth.GetUserInfo(token)
	if err != nil {
		return err
	}

	if userInfo.Email == nil || userInfo.Sub == nil {
		return keycloakEnums.ErrorKeycloakMissingUsernameOrSub
	}

	_, err = c.accountRepository.CreateAccount(c.accountUseCases.NewAccountFromKeycloakUserInfo(userInfo))
	return c.accountUseCases.CheckCreateAccountErrors(err)
}
