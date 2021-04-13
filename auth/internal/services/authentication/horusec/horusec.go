package horusec

import (
	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	accountRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/account"
)

type IService interface {
}

type Service struct {
	accountRepository accountRepository.IRepository
}

func NewHorusecAuthenticationService() IService {
	return &Service{}
}

func (s *Service) Login(credentials *authEntities.LoginCredentials) (*authEntities.LoginResponse, error) {
	account, err := s.accountRepository.GetAccountByEmail(credentials.Username)
	if err != nil {
		return nil, err
	}

	if account.IsNotConfirmed() {

	}
	credentials.IsInvalidUsernameEmail()

}

func (s *Service) checkLoginData(credentials *authEntities.LoginCredentials, account *accountEntities.Account) (*authEntities.LoginResponse, error) {
	if !credentials.CheckPassword(credentials.Password, account.Password) {

	}

	if  != nil {

	}


}
