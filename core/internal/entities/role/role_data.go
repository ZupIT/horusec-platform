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

package role

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"
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

func (d *Data) SetDataIDs(accountID uuid.UUID, workspaceID, repositoryID string) *Data {
	d.AccountID = accountID
	d.WorkspaceID = parser.ParseStringToUUID(workspaceID)
	d.RepositoryID = parser.ParseStringToUUID(repositoryID)

	return d
}

func (d *Data) ToBytes() []byte {
	bytes, _ := json.Marshal(d)

	return bytes
}
