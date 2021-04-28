package account

import (
	"time"

	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/crypto"
	tokenEntities "github.com/ZupIT/horusec-devkit/pkg/utils/jwt/entities"

	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
)

type Account struct {
	AccountID          uuid.UUID `json:"accountID" gorm:"primary_key"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	Username           string    `json:"username"`
	IsConfirmed        bool      `json:"isConfirmed"`
	IsApplicationAdmin bool      `json:"isApplicationAdmin"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}

func (a *Account) ToResponse() *Response {
	return &Response{
		AccountID:          a.AccountID,
		Email:              a.Email,
		Username:           a.Username,
		IsConfirmed:        a.IsConfirmed,
		IsApplicationAdmin: a.IsApplicationAdmin,
		CreatedAt:          a.CreatedAt,
		UpdatedAt:          a.UpdatedAt,
	}
}

func (a *Account) IsNotConfirmed() bool {
	return !a.IsConfirmed
}

func (a *Account) ToTokenData() *tokenEntities.TokenData {
	return &tokenEntities.TokenData{
		Email:     a.Email,
		Username:  a.Username,
		AccountID: a.AccountID,
	}
}

func (a *Account) ToLoginResponse(accessToken, refreshToken string, expiresAt time.Time) *authEntities.LoginResponse {
	return &authEntities.LoginResponse{
		AccessToken:        accessToken,
		RefreshToken:       refreshToken,
		ExpiresAt:          expiresAt,
		AccountID:          a.AccountID,
		Username:           a.Username,
		Email:              a.Email,
		IsApplicationAdmin: a.IsApplicationAdmin,
	}
}

func (a *Account) HashPassword() {
	hash, _ := crypto.HashPasswordBcrypt(a.Password)
	a.Password = hash
}

func (a *Account) SetNewAccountData() *Account {
	a.HashPassword()
	a.AccountID = uuid.New()
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()

	return a
}

func (a *Account) ToGetAccountDataResponse(permissions []string) *proto.GetAccountDataResponse {
	return &proto.GetAccountDataResponse{
		AccountID:          a.AccountID.String(),
		IsApplicationAdmin: a.IsApplicationAdmin,
		Permissions:        permissions,
	}
}

func (a *Account) SetIsConfirmedTrue() *Account {
	a.Update()
	a.IsConfirmed = true

	return a
}

func (a *Account) Update() *Account {
	a.UpdatedAt = time.Now()

	return a
}

func (a *Account) SetNewPassword(password string) *Account {
	a.Password = password
	a.HashPassword()

	return a.Update()
}

func (a *Account) UpdateFromUpdateAccountData(data *UpdateAccount) {
	a.Update()
	a.Email = data.Email
	a.Username = data.Username
	a.IsConfirmed = data.IsConfirmed
}
