package account

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"

	utils "github.com/ZupIT/horusec-devkit/pkg/utils/validation"
)

type ChangePasswordData struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	AccountID uuid.UUID `json:"accountID"`
}

func (c *ChangePasswordData) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Email, validation.Required, validation.Length(1, 255), is.EmailFormat),
		validation.Field(&c.Password, utils.PasswordValidationRules()...),
	)
}

func (c *ChangePasswordData) SetAccountID(accountID uuid.UUID) *ChangePasswordData {
	c.AccountID = accountID

	return c
}
