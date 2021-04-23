package account

import validation "github.com/go-ozzo/ozzo-validation/v4"

type RefreshToken struct {
	RefreshToken string `json:"refreshToken"`
}

func (r *RefreshToken) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.RefreshToken, validation.Required, validation.Length(1, 500)),
	)
}
