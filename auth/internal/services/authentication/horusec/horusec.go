package horusec

import (
	"github.com/patrickmn/go-cache"

	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt"

	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	authEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication"
	accountRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/account"
	authUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/authentication"
)

type IService interface {
	Login(credentials *authEntities.LoginCredentials) (*authEntities.LoginResponse, error)
}

type Service struct {
	accountRepository accountRepository.IRepository
	authUseCases      authUseCases.IUseCases
	cache             *cache.Cache
}

func NewHorusecAuthenticationService(repositoryAccount accountRepository.IRepository,
	useCasesAuth authUseCases.IUseCases) IService {
	return &Service{
		cache:             cache.New(authEnums.TokenDuration, authEnums.TokenCheckExpiredDuration),
		authUseCases:      useCasesAuth,
		accountRepository: repositoryAccount,
	}
}

func (s *Service) Login(credentials *authEntities.LoginCredentials) (*authEntities.LoginResponse, error) {
	account, err := s.accountRepository.GetAccountByEmail(credentials.Username)
	if err != nil {
		return nil, authEnums.ErrorWrongEmailOrPassword
	}

	if err := s.authUseCases.CheckLoginData(credentials, account); err != nil {
		return nil, err
	}

	return s.setTokensAndResponse(account)
}

func (s *Service) setTokensAndResponse(account *accountEntities.Account) (*authEntities.LoginResponse, error) {
	refreshToken := jwt.CreateRefreshToken()

	accessToken, expireAt, err := jwt.CreateToken(account.ToTokenData(), nil)
	if err != nil {
		return nil, err
	}

	if err = s.setRefreshTokenCache(account.AccountID.String(), refreshToken); err != nil {
		return nil, err
	}

	return account.ToLoginResponse(accessToken, refreshToken, expireAt), nil
}

func (s *Service) setRefreshTokenCache(accountID, refreshToken string) error {
	s.cache.Delete(accountID)
	return s.cache.Add(accountID, refreshToken, authEnums.TokenDuration)
}
