package role

import (
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
)

type Response struct {
	AccountID uuid.UUID    `json:"accountID"`
	Email     string       `json:"email"`
	Username  string       `json:"username"`
	Role      account.Role `json:"role"`
}
