package role

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
)

type Data struct {
	Role         account.Role `json:"role"`
	Email        string       `json:"email"`
	AccountID    uuid.UUID    `json:"accountID"`
	Username     string       `json:"username"`
	WorkspaceID  uuid.UUID    `json:"workspaceID"`
	RepositoryID uuid.UUID    `json:"repositoryID"`
}

func (r *Data) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Role, validation.Required, validation.In(r.Role.Values())),
		validation.Field(&r.Email, validation.Length(0, 255), is.EmailFormat),
		validation.Field(&r.Username, validation.Length(0, 255)),
		validation.Field(&r.AccountID, is.UUID),
		validation.Field(&r.WorkspaceID, is.UUID),
		validation.Field(&r.RepositoryID, is.UUID),
	)
}

func (r *Data) SetAccountAndWorkspaceID(accountID, workspaceID uuid.UUID) *Data {
	r.AccountID = accountID
	r.SetWorkspaceID(workspaceID)

	return r
}

func (r *Data) SetWorkspaceID(workspaceID uuid.UUID) *Data {
	r.WorkspaceID = workspaceID

	return r
}

func (r *Data) SetAccountID(accountID uuid.UUID) *Data {
	r.AccountID = accountID

	return r
}
