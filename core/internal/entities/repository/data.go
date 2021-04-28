package repository

import (
	"encoding/json"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"
	utilsValidation "github.com/ZupIT/horusec-devkit/pkg/utils/validation"
)

type Data struct {
	WorkspaceID     uuid.UUID `json:"workspaceID" swaggerignore:"true"`
	RepositoryID    uuid.UUID `json:"repositoryID" swaggerignore:"true"`
	AccountID       uuid.UUID `json:"accountID" swaggerignore:"true"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	AuthzMember     []string  `json:"authzMember"`
	AuthzAdmin      []string  `json:"authzAdmin"`
	AuthzSupervisor []string  `json:"authzSupervisor"`
	Permissions     []string  `json:"permissions" swaggerignore:"true"`
}

func (d *Data) Validate() error {
	return validation.ValidateStruct(d,
		validation.Field(&d.Name, validation.Required, validation.Length(1, 255)),
		validation.Field(&d.Description, validation.Length(0, 255)),
		validation.Field(&d.AuthzAdmin, validation.Length(0, 5)),
		validation.Field(&d.AuthzMember, validation.Length(0, 5)),
		validation.Field(&d.AuthzSupervisor, validation.Length(0, 5)),
		validation.Field(&d.AccountID, is.UUID),
		validation.Field(&d.WorkspaceID, is.UUID),
		validation.Field(&d.RepositoryID, is.UUID),
		validation.Field(&d.Permissions, validation.Empty),
	)
}

func (d *Data) CheckLdapGroups(authorizationType auth.AuthenticationType) error {
	return utilsValidation.CheckInvalidLdapGroups(authorizationType, d.AuthzAdmin, d.Permissions)
}

func (d *Data) SetWorkspaceIDAndAccountData(workspaceID uuid.UUID, accountData *proto.GetAccountDataResponse) *Data {
	d.AccountID = parser.ParseStringToUUID(accountData.AccountID)
	d.Permissions = accountData.Permissions
	d.WorkspaceID = workspaceID

	return d
}

func (d *Data) ToRepository() *Repository {
	return &Repository{
		RepositoryID:    uuid.New(),
		WorkspaceID:     d.WorkspaceID,
		Name:            d.Name,
		Description:     d.Description,
		AuthzMember:     d.AuthzMember,
		AuthzAdmin:      d.AuthzAdmin,
		AuthzSupervisor: d.AuthzSupervisor,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func (d *Data) ToBytes() []byte {
	bytes, _ := json.Marshal(d)

	return bytes
}

func (d *Data) SetWorkspaceAndRepositoryID(workspaceID, repositoryID uuid.UUID) *Data {
	d.RepositoryID = repositoryID
	d.WorkspaceID = workspaceID

	return d
}
