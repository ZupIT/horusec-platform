package account

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Email struct {
	Email string `json:"email"`
}

func (e *Email) Validate() error {
	return validation.ValidateStruct(e,
		validation.Field(&e.Email, validation.Required, validation.Length(1, 255), is.EmailFormat),
	)
}
