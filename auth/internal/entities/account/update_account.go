package account

import (
	"encoding/json"

	"github.com/ZupIT/horusec-devkit/pkg/enums/ozzovalidation"

	"github.com/go-ozzo/ozzo-validation/v4/is"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type UpdateAccount struct {
	AccountID   uuid.UUID `json:"accountID" swaggerignore:"true"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	IsConfirmed bool      `json:"isConfirmed" swaggerignore:"true"`
}

func (u *UpdateAccount) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.EmailFormat,
			validation.Length(ozzovalidation.Length0, ozzovalidation.Length255)),
		validation.Field(&u.Username, validation.Required,
			validation.Length(ozzovalidation.Length0, ozzovalidation.Length255)),
	)
}

func (u *UpdateAccount) SetAccountIDAndIsConfirmed(accountID uuid.UUID, isConfirmed bool) *UpdateAccount {
	u.AccountID = accountID
	u.IsConfirmed = isConfirmed

	return u
}

func (u *UpdateAccount) HasEmailChange(email string) bool {
	return u.Email != "" && email != u.Email
}

func (u *UpdateAccount) ToBytes() []byte {
	bytes, _ := json.Marshal(u)

	return bytes
}
