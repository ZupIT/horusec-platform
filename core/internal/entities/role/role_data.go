package role

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
)

type Data struct {
	Role         account.Role `json:"role"`
	AccountID    uuid.UUID    `json:"accountID" swaggerignore:"true"`
	WorkspaceID  uuid.UUID    `json:"workspaceID" swaggerignore:"true"`
	RepositoryID uuid.UUID    `json:"repositoryID" swaggerignore:"true"`
}

func (d *Data) Validate() error {
	return validation.ValidateStruct(d,
		validation.Field(&d.Role, validation.Required, validation.In(
			account.Admin, account.Supervisor, account.Member)),
		validation.Field(&d.AccountID, is.UUID),
		validation.Field(&d.WorkspaceID, is.UUID),
		validation.Field(&d.RepositoryID, is.UUID),
	)
}

func (d *Data) SetAccountAndWorkspaceID(accountID, workspaceID uuid.UUID) *Data {
	d.AccountID = accountID
	d.WorkspaceID = workspaceID

	return d
}

func (d *Data) ToBytes() []byte {
	bytes, _ := json.Marshal(d)

	return bytes
}
