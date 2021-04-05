package workspace

import (
	"encoding/json"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"

	workspaceEnums "github.com/ZupIT/horusec-platform/core/internal/enums/workspace"
	"github.com/ZupIT/horusec-platform/core/internal/utils"
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
		validation.Field(&c.AuthzMember, validation.Length(0, 5)),
		validation.Field(&c.AccountID, is.UUID),
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

func (c *CreateWorkspaceData) SetAccountData(accountData *proto.GetAccountDataResponse) {
	accountID, _ := uuid.Parse(accountData.AccountID)

	c.AccountID = accountID
	c.Permissions = accountData.Permissions
}

func (c *CreateWorkspaceData) ToBytes() []byte {
	bytes, _ := json.Marshal(c)
	return bytes
}

func (c *CreateWorkspaceData) CheckLdapGroups(authorizationType auth.AuthorizationType) error {
	if utils.IsInvalidLdapGroups(authorizationType, c.AuthzAdmin, c.Permissions) {
		return workspaceEnums.ErrorInvalidLdapGroup
	}

	return nil
}
