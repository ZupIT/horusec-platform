package account

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/ZupIT/horusec-platform/auth/internal/enums"
)

type CheckEmailAndUsername struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (c *CheckEmailAndUsername) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Email, validation.Required, is.EmailFormat,
			validation.Length(enums.MinDefaultColumnLength, enums.MaxDefaultColumnLength)),
		validation.Field(&c.Username, validation.Required,
			validation.Length(enums.MinDefaultColumnLength, enums.MaxDefaultColumnLength)),
	)
}

func (c *CheckEmailAndUsername) ToBytes() []byte {
	bytes, _ := json.Marshal(c)

	return bytes
}
