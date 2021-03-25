package workspace

import (
	"time"

	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
)

type AccountWorkspace struct {
	WorkspaceID uuid.UUID    `json:"companyID"`
	AccountID   uuid.UUID    `json:"accountID"`
	Role        account.Role `json:"role"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}
