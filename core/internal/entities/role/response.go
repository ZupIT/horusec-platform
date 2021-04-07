package role

import (
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
)

type Response struct {
	AccountID uuid.UUID    `json:"accountID,omitempty"`
	Email     string       `json:"email,omitempty"`
	Username  string       `json:"username,omitempty"`
	Role      account.Role `json:"role,omitempty"`
}
