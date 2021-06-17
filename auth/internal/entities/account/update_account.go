package account

import (
	"encoding/json"

	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/ZupIT/horusec-platform/auth/internal/enums"

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
			validation.Length(enums.MinDefaultColumnLength, enums.MaxDefaultColumnLength)),
		validation.Field(&u.Username, validation.Required,
			validation.Length(enums.MinDefaultColumnLength, enums.MaxDefaultColumnLength)),
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
