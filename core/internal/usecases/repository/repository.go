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

package repository

import (
	"io"

	"github.com/google/uuid"

	emailEntities "github.com/ZupIT/horusec-devkit/pkg/entities/email"
	emailEnums "github.com/ZupIT/horusec-devkit/pkg/enums/email"
	databaseEnums "github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	repositoryEntities "github.com/ZupIT/horusec-platform/core/internal/entities/repository"
	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
)

type IUseCases interface {
	RepositoryDataFromIOReadCloser(body io.ReadCloser) (*repositoryEntities.Data, error)
	FilterRepositoryByNameAndWorkspace(workspaceID uuid.UUID, name string) map[string]interface{}
	IsNotFoundError(err error) bool
	NewRepositoryData(repositoryID, workspaceID uuid.UUID,
		accountData *proto.GetAccountDataResponse) *repositoryEntities.Data
	FilterRepositoryByID(repositoryID uuid.UUID) map[string]interface{}
	FilterAccountRepositoryByID(accountID, repositoryID uuid.UUID) map[string]interface{}
	NewRepositoryInviteEmail(email, username, repositoryName string) []byte
	InheritWorkspaceGroups(repository *repositoryEntities.Repository,
		workspace *workspaceEntities.Workspace) *repositoryEntities.Repository
}

type UseCases struct {
}

func NewRepositoryUseCases() IUseCases {
	return &UseCases{}
}

func (u *UseCases) RepositoryDataFromIOReadCloser(body io.ReadCloser) (*repositoryEntities.Data, error) {
	data := &repositoryEntities.Data{}

	if err := parser.ParseBodyToEntity(body, data); err != nil {
		return nil, err
	}

	return data, data.Validate()
}

func (u *UseCases) FilterRepositoryByNameAndWorkspace(workspaceID uuid.UUID, name string) map[string]interface{} {
	return map[string]interface{}{"workspace_id": workspaceID, "name": name}
}

func (u *UseCases) IsNotFoundError(err error) bool {
	if err != nil {
		if err == databaseEnums.ErrorNotFoundRecords {
			return true
		}
	}

	return false
}

func (u *UseCases) NewRepositoryData(repositoryID, workspaceID uuid.UUID,
	accountData *proto.GetAccountDataResponse) *repositoryEntities.Data {
	return &repositoryEntities.Data{
		RepositoryID:       repositoryID,
		WorkspaceID:        workspaceID,
		AccountID:          parser.ParseStringToUUID(accountData.AccountID),
		Permissions:        accountData.Permissions,
		IsApplicationAdmin: accountData.IsApplicationAdmin,
	}
}

func (u *UseCases) FilterRepositoryByID(repositoryID uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"repository_id": repositoryID}
}

func (u *UseCases) FilterAccountRepositoryByID(accountID, repositoryID uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"account_id": accountID, "repository_id": repositoryID}
}

func (u *UseCases) NewRepositoryInviteEmail(email, username, repositoryName string) []byte {
	emailMessage := &emailEntities.Message{
		To:           email,
		TemplateName: emailEnums.RepositoryInvite,
		Subject:      "[Horusec] Repository invite",
		Data: map[string]interface{}{
			"repositoryName": repositoryName,
			"username":       username,
		},
	}

	return emailMessage.ToBytes()
}

func (u *UseCases) InheritWorkspaceGroups(repository *repositoryEntities.Repository,
	workspace *workspaceEntities.Workspace) *repositoryEntities.Repository {
	if !repository.ContainsAllAuthzGroups() {
		repository.AuthzAdmin = u.replaceGroupsIfEmpty(repository.AuthzAdmin, workspace.AuthzAdmin)
		repository.AuthzSupervisor = u.replaceGroupsIfEmpty(repository.AuthzSupervisor, workspace.AuthzAdmin)
		repository.AuthzMember = u.replaceGroupsIfEmpty(repository.AuthzMember, workspace.AuthzMember)
	}

	return repository
}

func (u *UseCases) replaceGroupsIfEmpty(repositoryGroups, workspaceGroups []string) []string {
	if len(repositoryGroups) == 0 {
		return workspaceGroups
	}

	return repositoryGroups
}
