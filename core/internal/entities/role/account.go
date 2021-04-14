package role

import "github.com/google/uuid"

type Account struct {
	AccountID uuid.UUID
	Username  string
	Email     string
}
