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

	"github.com/ZupIT/horusec-devkit/pkg/enums/ozzovalidation"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
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
		validation.Field(&u.Email, validation.Required, is.EmailFormat,
			validation.Length(ozzovalidation.Length1, ozzovalidation.Length255)),
		validation.Field(&u.Username, validation.Length(ozzovalidation.Length0, ozzovalidation.Length255)),
		validation.Field(&u.AccountID, is.UUID),
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

func (u *UserData) SetWorkspaceIDAndAccountData(workspaceID string, data *proto.GetAccountDataResponse) *UserData {
	u.WorkspaceID = parser.ParseStringToUUID(workspaceID)
	u.AccountID = parser.ParseStringToUUID(data.AccountID)
	u.Email = data.Email
	u.Username = data.Username

	return u
}
