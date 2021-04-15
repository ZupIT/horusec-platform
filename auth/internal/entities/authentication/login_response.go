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
	ExpiresIn          int       `json:"expiresIn"`
	RefreshExpiresIn   int       `json:"refreshExpiresIn"`
	IsApplicationAdmin bool      `json:"isApplicationAdmin"`
}
