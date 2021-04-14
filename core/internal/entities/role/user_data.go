package role

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"
)

type UserData struct {
	Role         account.Role `json:"role"`
	Email        string       `json:"email"`
	AccountID    uuid.UUID    `json:"accountID"`
	Username     string       `json:"username"`
	WorkspaceID  uuid.UUID    `json:"workspaceID" swaggerignore:"true"`
	RepositoryID uuid.UUID    `json:"repositoryID" swaggerignore:"true"`
}

func (u *UserData) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Role, validation.Required, validation.In(
			account.Admin, account.Supervisor, account.Member)),
		validation.Field(&u.Email, validation.Required, validation.Length(1, 255), is.EmailFormat),
		validation.Field(&u.Username, validation.Required, validation.Length(1, 255)),
		validation.Field(&u.AccountID, validation.Required, is.UUID),
		validation.Field(&u.WorkspaceID, is.UUID),
		validation.Field(&u.RepositoryID, is.UUID),
	)
}

func (u *UserData) SetIDs(workspaceID, repositoryID string) *UserData {
	u.WorkspaceID = parser.ParseStringToUUID(workspaceID)
	u.RepositoryID = parser.ParseStringToUUID(repositoryID)

	return u
}

func (u *UserData) ToBytes() []byte {
	bytes, _ := json.Marshal(u)

	return bytes
}
