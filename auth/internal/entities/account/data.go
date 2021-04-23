package account

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Data struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (d *Data) Validate() error {
	return validation.ValidateStruct(d,
		validation.Field(&d.Email, validation.Required, validation.Length(1, 255), is.EmailFormat),
		validation.Field(&d.Password, validation.Length(1, 255), validation.Required),
		validation.Field(&d.Username, validation.Length(1, 255), validation.Required),
	)
}

func (d *Data) ToAccount() *Account {
	account := &Account{
		Email:    d.Email,
		Password: d.Password,
		Username: d.Username,
	}

	return account.SetNewAccountData()
}
