package account

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AccessToken struct {
	AccessToken string `json:"accessToken"`
}

func (a *AccessToken) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.AccessToken, validation.Required, validation.Length(1, 500)),
	)
}
