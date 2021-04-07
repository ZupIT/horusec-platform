package workspace

import (
	"encoding/json"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	workspaceEnums "github.com/ZupIT/horusec-platform/core/internal/enums/workspace"
	"github.com/ZupIT/horusec-platform/core/internal/utils"
)

type Data struct {
	WorkspaceID uuid.UUID `json:"workspaceID" swaggerignore:"true"`
	AccountID   uuid.UUID `json:"accountID" swaggerignore:"true"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	AuthzMember []string  `json:"authzMember"`
	AuthzAdmin  []string  `json:"authzAdmin"`
	Permissions []string  `json:"permissions" swaggerignore:"true"`
}

func (d *Data) Validate() error {
	return validation.ValidateStruct(d,
		validation.Field(&d.Name, validation.Required, validation.Length(1, 255)),
		validation.Field(&d.Description, validation.Length(0, 255)),
		validation.Field(&d.AuthzAdmin, validation.Length(0, 5)),
		validation.Field(&d.AuthzMember, validation.Length(0, 5)),
		validation.Field(&d.AccountID, is.UUID),
		validation.Field(&d.WorkspaceID, is.UUID),
		validation.Field(&d.Permissions, validation.Empty),
	)
}

func (d *Data) ToWorkspace() *Workspace {
	return &Workspace{
		WorkspaceID: uuid.New(),
		Name:        d.Name,
		Description: d.Description,
		AuthzMember: d.AuthzMember,
		AuthzAdmin:  d.AuthzAdmin,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (d *Data) ToBytes() []byte {
	bytes, _ := json.Marshal(d)

	return bytes
}

func (d *Data) CheckLdapGroups(authorizationType auth.AuthorizationType) error {
	if utils.IsInvalidLdapGroups(authorizationType, d.AuthzAdmin, d.Permissions) {
		return workspaceEnums.ErrorInvalidLdapGroup
	}

	return nil
}

func (d *Data) SetWorkspaceID(workspaceID uuid.UUID) *Data {
	d.WorkspaceID = workspaceID

	return d
}

func (d *Data) SetAccountData(accountData *proto.GetAccountDataResponse) *Data {
	d.AccountID = parser.ParseStringToUUID(accountData.AccountID)
	d.Permissions = accountData.Permissions

	return d
}
