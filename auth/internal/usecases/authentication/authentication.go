package authentication

import (
	"io"

	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	horusecAuthEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication/horusec"
)

type IUseCases interface {
	CheckLoginData(credentials *authEntities.LoginCredentials, account *accountEntities.Account) error
	LoginCredentialsFromIOReadCloser(body io.ReadCloser) (*authEntities.LoginCredentials, error)
	SetLdapAccountData(userData map[string]string) *accountEntities.Account
	FilterWorkspaceByID(workspaceID uuid.UUID) map[string]interface{}
	FilterRepositoryByID(repository uuid.UUID) map[string]interface{}
	FilterAccountWorkspaceByID(accountID, workspaceID uuid.UUID) map[string]interface{}
	FilterAccountRepositoryByID(accountID, repository uuid.UUID) map[string]interface{}
}

type UseCases struct {
}

func NewAuthenticationUseCases() IUseCases {
	return &UseCases{}
}

func (u *UseCases) CheckLoginData(credentials *authEntities.LoginCredentials, account *accountEntities.Account) error {
	if credentials.CheckInvalidPassword(credentials.Password, account.Password) ||
		credentials.IsInvalidUsernameEmail() {
		return horusecAuthEnums.ErrorWrongEmailOrPassword
	}

	if account.IsNotConfirmed() {
		return horusecAuthEnums.ErrorAccountEmailNotConfirmed
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

func (u *UseCases) SetLdapAccountData(userData map[string]string) *accountEntities.Account {
	account := &accountEntities.Account{
		Username: userData["sAMAccountName"],
		Password: uuid.NewString(),
	}

	if userData["mail"] == "" {
		account.Email = userData["sAMAccountName"]
	} else {
		account.Email = userData["mail"]
	}

	return account.SetNewAccountData()
}

func (u *UseCases) FilterWorkspaceByID(workspaceID uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"workspace_id": workspaceID}
}

func (u *UseCases) FilterRepositoryByID(repository uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"repository_id": repository}
}

func (u *UseCases) FilterAccountWorkspaceByID(accountID, workspaceID uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"account_id": accountID, "workspace_id": workspaceID}
}

func (u *UseCases) FilterAccountRepositoryByID(accountID, repository uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"account_id": accountID, "repository_id": repository}
}
