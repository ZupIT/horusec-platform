package account

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"

	utils "github.com/ZupIT/horusec-devkit/pkg/utils/validation"
)

type ChangePasswordData struct {
	Password  string    `json:"password"`
	AccountID uuid.UUID `json:"accountID" swaggerignore:"true"`
}

func (c *ChangePasswordData) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Password, utils.PasswordValidationRules()...),
	)
}

func (c *ChangePasswordData) SetAccountID(accountID uuid.UUID) *ChangePasswordData {
	c.AccountID = accountID

	return c
}
