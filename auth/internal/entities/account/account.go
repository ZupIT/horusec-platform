package account

import (
	"time"

	"github.com/google/uuid"
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
