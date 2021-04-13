package authentication

import "time"

type LoginResponse struct {
	AccessToken        string    `json:"accessToken omitempty"`
	RefreshToken       string    `json:"refreshToken omitempty"`
	Username           string    `json:"username omitempty"`
	Email              string    `json:"email omitempty"`
	ExpiresAt          time.Time `json:"expiresAt omitempty"`
	IsApplicationAdmin bool      `json:"isApplicationAdmin omitempty"`
}
