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
	"io"

	"github.com/google/uuid"

	emailEntities "github.com/ZupIT/horusec-devkit/pkg/entities/email"
	emailEnums "github.com/ZupIT/horusec-devkit/pkg/enums/email"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	envUtils "github.com/ZupIT/horusec-devkit/pkg/utils/env"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
)

type IUseCases interface {
	WorkspaceDataFromIOReadCloser(body io.ReadCloser) (data *workspace.Data, err error)
	FilterAccountWorkspaceByID(accountID, workspaceID uuid.UUID) map[string]interface{}
	FilterWorkspaceByID(workspaceID uuid.UUID) map[string]interface{}
	NewWorkspaceData(workspaceID uuid.UUID, accountData *proto.GetAccountDataResponse) *workspace.Data
	NewOrganizationInviteEmail(email, username, workspaceName string) []byte
}

type UseCases struct {
}

func NewWorkspaceUseCases() IUseCases {
	return &UseCases{}
}

func (u *UseCases) WorkspaceDataFromIOReadCloser(body io.ReadCloser) (*workspace.Data, error) {
	data := &workspace.Data{}

	if err := parser.ParseBodyToEntity(body, data); err != nil {
		return nil, err
	}

	return data, data.Validate()
}

func (u *UseCases) FilterAccountWorkspaceByID(accountID, workspaceID uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"account_id": accountID, "workspace_id": workspaceID}
}

func (u *UseCases) FilterWorkspaceByID(workspaceID uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"workspace_id": workspaceID}
}

func (u *UseCases) NewWorkspaceData(workspaceID uuid.UUID, accountData *proto.GetAccountDataResponse) *workspace.Data {
	data := &workspace.Data{
		WorkspaceID: workspaceID,
	}

	return data.SetAccountData(accountData)
}

func (u *UseCases) NewOrganizationInviteEmail(email, username, workspaceName string) []byte {
	emailMessage := &emailEntities.Message{
		To:           email,
		TemplateName: emailEnums.OrganizationInvite,
		Subject:      "[Horusec] Organization invite",
		Data: map[string]interface{}{
			"WorkspaceName": workspaceName,
			"Username":      username,
			"URL":           envUtils.GetHorusecManagerURL()},
	}

	return emailMessage.ToBytes()
}
