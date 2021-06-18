package account

import (
	"encoding/json"

	"github.com/ZupIT/horusec-devkit/pkg/enums/ozzovalidation"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type RefreshToken struct {
	RefreshToken string `json:"refreshToken"`
}

func (r *RefreshToken) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.RefreshToken, validation.Required,
			validation.Length(ozzovalidation.Length0, ozzovalidation.Length500)),
	)
}

func (r *RefreshToken) ToBytes() []byte {
	bytes, _ := json.Marshal(r)

	return bytes
}
