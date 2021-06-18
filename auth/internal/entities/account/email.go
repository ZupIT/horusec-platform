package account

import (
	"encoding/json"

	"github.com/ZupIT/horusec-devkit/pkg/enums/ozzovalidation"

	"github.com/go-ozzo/ozzo-validation/v4/is"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Email struct {
	Email string `json:"email"`
}

func (e *Email) Validate() error {
	return validation.ValidateStruct(e,
		validation.Field(&e.Email, validation.Required, is.EmailFormat,
			validation.Length(ozzovalidation.Length0, ozzovalidation.Length255)),
	)
}

func (e *Email) ToBytes() []byte {
	bytes, _ := json.Marshal(e)

	return bytes
}
