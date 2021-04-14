package account

import (
	"time"

	"github.com/google/uuid"

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
