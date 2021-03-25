package workspace

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type CreateWorkspaceData struct {
	AccountID   uuid.UUID `json:"accountID"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	AuthzMember []string  `json:"authzMember"`
	AuthzAdmin  []string  `json:"authzAdmin"`
	Permissions []string  `json:"permissions"`
}

func (c *CreateWorkspaceData) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Name, validation.Required, validation.Length(1, 255)),
		validation.Field(&c.Description, validation.Length(0, 255)),
		validation.Field(&c.AuthzAdmin, validation.Length(0, 5)),
		validation.Field(&c.Description, validation.Length(0, 5)),
		//todo add account id
		validation.Field(&c.Permissions, validation.Empty),
	)
}

func (c *CreateWorkspaceData) ToWorkspace() *Workspace {
	return &Workspace{
		WorkspaceID: uuid.New(),
		Name:        c.Name,
		Description: c.Description,
		AuthzMember: c.AuthzMember,
		AuthzAdmin:  c.AuthzAdmin,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}
}
