package account

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"

	"github.com/ZupIT/horusec-platform/auth/config/app"

	emailEnums "github.com/ZupIT/horusec-devkit/pkg/enums/email"

	"github.com/ZupIT/horusec-devkit/pkg/entities/email"

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
	AccountDataFromIOReadCloser(body io.ReadCloser) (*accountEntities.Data, error)
	NewAccountValidationEmail(account *accountEntities.Account) []byte
	EmailFromIOReadCloser(body io.ReadCloser) (string, error)
	GenerateResetPasswordCode() string
	NewResetPasswordCodeEmail(account *accountEntities.Account, code string) []byte
}

type UseCases struct {
	appConfig app.IConfig
}

func NewAccountUseCases(appConfig app.IConfig) IUseCases {
	return &UseCases{
		appConfig: appConfig,
	}
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
	data := &map[string]string{"accessToken": ""}

	if err := parser.ParseBodyToEntity(body, data); err != nil {
		return "", err
	}

	result := *data
	return result["accessToken"], nil
}

func (u *UseCases) AccountDataFromIOReadCloser(body io.ReadCloser) (*accountEntities.Data, error) {
	data := &accountEntities.Data{}

	if err := parser.ParseBodyToEntity(body, data); err != nil {
		return nil, err
	}

	return data, data.Validate()
}

func (u *UseCases) NewAccountValidationEmail(account *accountEntities.Account) []byte {
	message := &email.Message{
		To:           account.Email,
		Subject:      "[Horusec] Account Confirmation Email",
		TemplateName: emailEnums.AccountConfirmation,
		Data: map[string]interface{}{"Username": account.Username,
			"URL": u.getAccountValidationEmailURL(account.AccountID)},
	}

	return message.ToBytes()
}

func (u *UseCases) getAccountValidationEmailURL(accountID uuid.UUID) string {
	return fmt.Sprintf("%s/auth/account/validate/%s", u.appConfig.GetHorusecAuthURL(), accountID)
}

func (u *UseCases) EmailFromIOReadCloser(body io.ReadCloser) (string, error) {
	data := &map[string]string{"email": ""}

	if err := parser.ParseBodyToEntity(body, data); err != nil {
		return "", err
	}

	result := *data
	return result["email"], nil
}

func (u *UseCases) GenerateResetPasswordCode() string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	code := make([]byte, 6)
	for i := range code {
		code[i] = accountEnums.ResetPasswordCharset[seededRand.Intn(len(accountEnums.ResetPasswordCharset))]
	}

	return string(code)
}

func (u *UseCases) NewResetPasswordCodeEmail(account *accountEntities.Account, code string) []byte {
	message := &email.Message{
		To:           account.Email,
		Subject:      "[Horusec] Reset Password",
		TemplateName: emailEnums.ResetPassword,
		Data: map[string]interface{}{"Username": account.Username, "Code": code,
			"URL": u.getResetPasswordCodeEmailURL(account.Email, code)},
	}

	return message.ToBytes()
}

func (u *UseCases) getResetPasswordCodeEmailURL(email, code string) string {
	return fmt.Sprintf("%s/auth/recovery-password/check-code?email=%s&code=%s",
		u.appConfig.GetHorusecManagerURL(), email, code)
}
