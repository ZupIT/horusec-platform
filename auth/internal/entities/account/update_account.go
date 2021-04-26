package account

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type UpdateAccount struct {
	AccountID   uuid.UUID `json:"accountID" swaggerignore:"true"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	IsConfirmed bool      `json:"isConfirmed" swaggerignore:"true"`
}

func (u *UpdateAccount) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, validation.Length(1, 255), is.EmailFormat),
		validation.Field(&u.Username, validation.Length(1, 255), validation.Required),
	)
}

func (u *UpdateAccount) SetAccountIDAndIsConfirmed(accountID uuid.UUID, isConfirmed bool) *UpdateAccount {
	u.AccountID = accountID
	u.IsConfirmed = isConfirmed

	return u
}

func (u *UpdateAccount) HasEmailChange(email string) bool {
	return u.Email != "" && email != u.Email
}
