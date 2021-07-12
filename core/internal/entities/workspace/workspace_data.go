// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package workspace

import (
	"encoding/json"
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/enums/ozzovalidation"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"
	utilsValidation "github.com/ZupIT/horusec-devkit/pkg/utils/validation"
)

type Data struct {
	WorkspaceID        uuid.UUID `json:"workspaceID" swaggerignore:"true"`
	AccountID          uuid.UUID `json:"accountID" swaggerignore:"true"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	AuthzMember        []string  `json:"authzMember"`
	AuthzAdmin         []string  `json:"authzAdmin"`
	Permissions        []string  `json:"permissions" swaggerignore:"true"`
	IsApplicationAdmin bool      `json:"isApplicationAdmin" swaggerignore:"true"`
}

func (d *Data) Validate() error {
	return validation.ValidateStruct(d,
		validation.Field(&d.Name, validation.Required,
			validation.Length(ozzovalidation.Length1, ozzovalidation.Length255)),
		validation.Field(&d.Description, validation.Length(ozzovalidation.Length0, ozzovalidation.Length255)),
		validation.Field(&d.AuthzAdmin, validation.Length(ozzovalidation.Length0, ozzovalidation.Length5)),
		validation.Field(&d.AuthzMember, validation.Length(ozzovalidation.Length0, ozzovalidation.Length5)),
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

func (d *Data) CheckLdapGroups(authorizationType auth.AuthenticationType) error {
	return utilsValidation.CheckInvalidLdapGroups(authorizationType, d.AuthzAdmin, d.Permissions)
}

func (d *Data) SetWorkspaceID(workspaceID uuid.UUID) *Data {
	d.WorkspaceID = workspaceID

	return d
}

func (d *Data) SetAccountData(accountData *proto.GetAccountDataResponse) *Data {
	d.AccountID = parser.ParseStringToUUID(accountData.AccountID)
	d.Permissions = accountData.Permissions
	d.IsApplicationAdmin = accountData.IsApplicationAdmin

	return d
}
