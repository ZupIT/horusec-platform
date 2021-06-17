package account

import (
	"encoding/json"

	"github.com/ZupIT/horusec-platform/auth/internal/enums"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type RefreshToken struct {
	RefreshToken string `json:"refreshToken"`
}

func (r *RefreshToken) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.RefreshToken, validation.Required,
			validation.Length(enums.MinDefaultColumnLength, enums.MaxTokenLength)),
	)
}

func (r *RefreshToken) ToBytes() []byte {
	bytes, _ := json.Marshal(r)

	return bytes
}
