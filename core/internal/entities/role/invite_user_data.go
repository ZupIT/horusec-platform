package role

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
)

type InviteUserData struct {
	Role         account.Role `json:"role"`
	Email        string       `json:"email"`
	AccountID    uuid.UUID    `json:"accountID"`
	Username     string       `json:"username"`
	WorkspaceID  uuid.UUID    `json:"workspaceID" swaggerignore:"true"`
	RepositoryID uuid.UUID    `json:"repositoryID" swaggerignore:"true"`
}

func (i *InviteUserData) Validate() error {
	return validation.ValidateStruct(i,
		validation.Field(&i.Role, validation.Required, validation.In(
			account.Admin, account.Supervisor, account.Member)),
		validation.Field(&i.Email, validation.Required, validation.Length(1, 255), is.EmailFormat),
		validation.Field(&i.Username, validation.Required, validation.Length(1, 255)),
		validation.Field(&i.AccountID, validation.Required, is.UUID),
		validation.Field(&i.WorkspaceID, is.UUID),
		validation.Field(&i.RepositoryID, is.UUID),
	)
}

func (i *InviteUserData) SetWorkspaceID(workspaceID uuid.UUID) *InviteUserData {
	i.WorkspaceID = workspaceID

	return i
}

func (i *InviteUserData) ToBytes() []byte {
	bytes, _ := json.Marshal(i)

	return bytes
}
