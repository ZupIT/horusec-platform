package account

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type ResetCodeData struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

//nolint // valid magic number
func (r *ResetCodeData) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Email, validation.Required, validation.Length(1, 255), is.EmailFormat),
		validation.Field(&r.Code, validation.Required, validation.Length(6, 6)),
	)
}
