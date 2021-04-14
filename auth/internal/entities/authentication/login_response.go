package authentication

import (
	"time"

	"github.com/google/uuid"
)

type LoginResponse struct {
	AccountID          uuid.UUID `json:"accountID"`
	AccessToken        string    `json:"accessToken"`
	RefreshToken       string    `json:"refreshToken"`
	Username           string    `json:"username"`
	Email              string    `json:"email"`
	ExpiresAt          time.Time `json:"expiresAt"`
	IsApplicationAdmin bool      `json:"isApplicationAdmin"`
}
