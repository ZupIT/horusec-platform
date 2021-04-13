package account

import (
	"io"

	"github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"

	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"
	"github.com/google/uuid"
)

type IUseCases interface {
	FilterAccountByID(accountID uuid.UUID) map[string]interface{}
	LoginCredentialsFromIOReadCloser(body io.ReadCloser) (*authentication.LoginCredentials, error)
	FilterAccountByEmail(email string) map[string]interface{}
}

type UseCases struct {
}

func NewAccountUseCases() IUseCases {
	return &UseCases{}
}

func (u *UseCases) FilterAccountByID(accountID uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"account_id": accountID}
}

func (u *UseCases) FilterAccountByEmail(email string) map[string]interface{} {
	return map[string]interface{}{"email": email}
}

func (u *UseCases) LoginCredentialsFromIOReadCloser(body io.ReadCloser) (*authentication.LoginCredentials, error) {
	credentials := &authentication.LoginCredentials{}

	if err := parser.ParseBodyToEntity(body, credentials); err != nil {
		return nil, err
	}

	return credentials, credentials.Validate()
}
