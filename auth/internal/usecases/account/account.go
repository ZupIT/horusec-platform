package account

import (
	"io"
	"strings"
	"time"

	"github.com/Nerzal/gocloak/v7"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	accountEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/account"
)

type IUseCases interface {
	FilterAccountByID(accountID uuid.UUID) map[string]interface{}
	FilterAccountByEmail(email string) map[string]interface{}
	FilterAccountByUsername(username string) map[string]interface{}
	NewAccountFromKeycloakUserInfo(userInfo *gocloak.UserInfo) *accountEntities.Account
	CheckCreateAccountErrors(err error) error
	AccessTokenFromIOReadCloser(body io.ReadCloser) (string, error)
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

func (u *UseCases) FilterAccountByUsername(username string) map[string]interface{} {
	return map[string]interface{}{"username": username}
}

func (u *UseCases) NewAccountFromKeycloakUserInfo(userInfo *gocloak.UserInfo) *accountEntities.Account {
	username := *userInfo.PreferredUsername
	if username == "" {
		username = *userInfo.Name
	}

	return &accountEntities.Account{
		AccountID:   parser.ParseStringToUUID(*userInfo.Sub),
		Email:       *userInfo.Email,
		Username:    username,
		IsConfirmed: true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (u *UseCases) CheckCreateAccountErrors(err error) error {
	if u.contains(err, accountEnums.DuplicatedConstraintEmail) {
		return accountEnums.ErrorEmailAlreadyInUse
	}

	if u.contains(err, accountEnums.DuplicatedConstraintUsername) {
		return accountEnums.ErrorUsernameAlreadyInUse
	}

	if u.contains(err, accountEnums.DuplicatedConstraintPrimaryKey) {
		return accountEnums.ErrorUsernameAlreadyInUse
	}

	return err
}

func (u *UseCases) contains(err error, check string) bool {
	return strings.Contains(strings.ToLower(err.Error()), check)
}

func (u *UseCases) AccessTokenFromIOReadCloser(body io.ReadCloser) (string, error) {
	var accessToken string

	if err := parser.ParseBodyToEntity(body, accessToken); err != nil {
		return "", err
	}

	return accessToken, nil
}
