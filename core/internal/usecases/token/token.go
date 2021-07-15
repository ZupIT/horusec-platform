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

package token

import (
	"io"

	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	tokenEntities "github.com/ZupIT/horusec-platform/core/internal/entities/token"
)

type IUseCases interface {
	TokenDataFromIOReadCloser(body io.ReadCloser) (*tokenEntities.Data, error)
	FilterWorkspaceTokenByID(tokenID, workspaceID uuid.UUID) map[string]interface{}
	FilterRepositoryTokenByID(tokenID, workspaceID, repositoryID uuid.UUID) map[string]interface{}
	FilterListWorkspaceTokens(workspaceID uuid.UUID) map[string]interface{}
	FilterListRepositoryTokens(workspaceID, repositoryID uuid.UUID) map[string]interface{}
	NewTokenData(tokenID uuid.UUID, workspaceID, repositoryID string) *tokenEntities.Data
}

type UseCases struct {
}

func NewTokenUseCases() IUseCases {
	return &UseCases{}
}

func (u *UseCases) TokenDataFromIOReadCloser(body io.ReadCloser) (*tokenEntities.Data, error) {
	data := &tokenEntities.Data{}

	if err := parser.ParseBodyToEntity(body, data); err != nil {
		return nil, err
	}

	return data, data.Validate()
}

func (u *UseCases) FilterWorkspaceTokenByID(tokenID, workspaceID uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"token_id": tokenID, "workspace_id": workspaceID, "repository_id": nil}
}

func (u *UseCases) FilterRepositoryTokenByID(tokenID, workspaceID, repositoryID uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"token_id": tokenID, "workspace_id": workspaceID, "repository_id": repositoryID}
}

func (u *UseCases) FilterListWorkspaceTokens(workspaceID uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"workspace_id": workspaceID, "repository_id": nil}
}

func (u *UseCases) FilterListRepositoryTokens(workspaceID, repositoryID uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"workspace_id": workspaceID, "repository_id": repositoryID}
}

func (u *UseCases) NewTokenData(tokenID uuid.UUID, workspaceID, repositoryID string) *tokenEntities.Data {
	return &tokenEntities.Data{
		RepositoryID: parser.ParseStringToUUID(repositoryID),
		WorkspaceID:  parser.ParseStringToUUID(workspaceID),
		TokenID:      tokenID,
	}
}
