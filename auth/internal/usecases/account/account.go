package account

import (
	"github.com/google/uuid"
)

type IUseCases interface {
	FilterAccountByID(accountID uuid.UUID) map[string]interface{}
	FilterAccountByEmail(email string) map[string]interface{}
	FilterAccountByUsername(username string) map[string]interface{}
}

type UseCases struct {
}

func NewAccountUseCases() IUseCases {
	return &UseCases{}
}

func (u *UseCases) FilterAccountByID(accountID uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"account_id": accountID}
}

func (u *UseCases) FilterAccountByEmail(email string) map[string]interface{} {
	return map[string]interface{}{"email": email}
}

func (u *UseCases) FilterAccountByUsername(username string) map[string]interface{} {
	return map[string]interface{}{"username": username}
}
