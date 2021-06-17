package account

import (
	"encoding/json"

	"github.com/ZupIT/horusec-platform/auth/internal/enums"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AccessToken struct {
	AccessToken string `json:"accessToken"`
}

func (a *AccessToken) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.AccessToken, validation.Required,
			validation.Length(enums.MinDefaultColumnLength, enums.MaxTokenLength)),
	)
}

func (a *AccessToken) ToBytes() []byte {
	bytes, _ := json.Marshal(a)
	return bytes
}
