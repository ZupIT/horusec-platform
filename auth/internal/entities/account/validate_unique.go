package account

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CheckEmailAndUsername struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (c *CheckEmailAndUsername) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Email, validation.Required, validation.Length(1, 255), is.EmailFormat),
		validation.Field(&c.Username, validation.Required, validation.Length(1, 255)),
	)
}

func (c *CheckEmailAndUsername) ToBytes() []byte {
	bytes, _ := json.Marshal(c)

	return bytes
}
