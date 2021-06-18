package account

import (
	"encoding/json"

	"github.com/ZupIT/horusec-devkit/pkg/enums/ozzovalidation"

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
		validation.Field(&r.Email, validation.Required, validation.Length(ozzovalidation.Length1, ozzovalidation.Length255), is.EmailFormat),
		validation.Field(&r.Code, validation.Required, validation.Length(ozzovalidation.Length6, ozzovalidation.Length6)),
	)
}

func (r *ResetCodeData) ToBytes() []byte {
	bytes, _ := json.Marshal(r)

	return bytes
}
