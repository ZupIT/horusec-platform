package account

import (
	"encoding/json"

	"github.com/go-ozzo/ozzo-validation/v4/is"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/ZupIT/horusec-platform/auth/internal/enums"
)

type Email struct {
	Email string `json:"email"`
}

func (e *Email) Validate() error {
	return validation.ValidateStruct(e,
		validation.Field(&e.Email, validation.Required, is.EmailFormat,
			validation.Length(enums.MinDefaultColumnLength, enums.MaxDefaultColumnLength)),
	)
}

func (e *Email) ToBytes() []byte {
	bytes, _ := json.Marshal(e)

	return bytes
}
