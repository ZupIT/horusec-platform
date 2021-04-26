package account

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	utils "github.com/ZupIT/horusec-devkit/pkg/utils/validation"
)

type Data struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (u *Data) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, validation.Length(1, 255), is.EmailFormat),
		validation.Field(&u.Password, utils.PasswordValidationRules()...),
		validation.Field(&u.Username, validation.Length(1, 255), validation.Required),
	)
}

func (u *Data) ToAccount() *Account {
	account := &Account{
		Email:    u.Email,
		Password: u.Password,
		Username: u.Username,
	}

	return account.SetNewAccountData()
}
