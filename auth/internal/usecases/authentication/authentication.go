package authentication

import (
	"io"

	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	authEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication"
)

type IUseCases interface {
	CheckLoginData(credentials *authEntities.LoginCredentials, account *accountEntities.Account) error
	LoginCredentialsFromIOReadCloser(body io.ReadCloser) (*authEntities.LoginCredentials, error)
}

type UseCases struct {
}

func NewAuthenticationUseCases() IUseCases {
	return &UseCases{}
}

func (u *UseCases) CheckLoginData(credentials *authEntities.LoginCredentials, account *accountEntities.Account) error {
	if credentials.CheckInvalidPassword(credentials.Password, account.Password) ||
		credentials.IsInvalidUsernameEmail() {
		return authEnums.ErrorWrongEmailOrPassword
	}

	if account.IsNotConfirmed() {
		return authEnums.ErrorAccountEmailNotConfirmed
	}

	return nil
}

func (u *UseCases) LoginCredentialsFromIOReadCloser(body io.ReadCloser) (*authEntities.LoginCredentials, error) {
	credentials := &authEntities.LoginCredentials{}

	if err := parser.ParseBodyToEntity(body, credentials); err != nil {
		return nil, err
	}

	return credentials, credentials.Validate()
}
