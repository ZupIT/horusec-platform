package account

import (
	"encoding/json"

	"github.com/ZupIT/horusec-devkit/pkg/enums/ozzovalidation"

	"github.com/go-ozzo/ozzo-validation/v4/is"

	utils "github.com/ZupIT/horusec-devkit/pkg/utils/validation"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Data struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (u *Data) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.EmailFormat,
			validation.Length(ozzovalidation.Length0, ozzovalidation.Length255)),
		validation.Field(&u.Password, utils.PasswordValidationRules()...),
		validation.Field(&u.Username, validation.Required,
			validation.Length(ozzovalidation.Length0, ozzovalidation.Length255)),
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

func (u *Data) ToBytes() []byte {
	bytes, _ := json.Marshal(u)
	return bytes
}
