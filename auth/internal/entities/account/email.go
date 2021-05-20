package account

import (
	"encoding/json"

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

func (e *Email) ToBytes() []byte {
	bytes, _ := json.Marshal(e)

	return bytes
}
