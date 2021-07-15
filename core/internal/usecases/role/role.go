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
	"io"

	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/core/internal/entities/role"
)

type IUseCases interface {
	NewRoleData(accountID, workspaceID, repositoryID uuid.UUID) *role.Data
	InviteUserDataFromIOReadCloser(body io.ReadCloser) (*role.UserData, error)
	RoleDataFromIOReadCloser(body io.ReadCloser) (*role.Data, error)
}

type UseCases struct {
}

func NewRoleUseCases() IUseCases {
	return &UseCases{}
}

func (u *UseCases) NewRoleData(accountID, workspaceID, repositoryID uuid.UUID) *role.Data {
	return &role.Data{
		WorkspaceID:  workspaceID,
		RepositoryID: repositoryID,
		AccountID:    accountID,
	}
}

func (u *UseCases) InviteUserDataFromIOReadCloser(body io.ReadCloser) (*role.UserData, error) {
	data := &role.UserData{}

	if err := parser.ParseBodyToEntity(body, data); err != nil {
		return nil, err
	}

	return data, data.Validate()
}

func (u *UseCases) RoleDataFromIOReadCloser(body io.ReadCloser) (*role.Data, error) {
	data := &role.Data{}

	if err := parser.ParseBodyToEntity(body, data); err != nil {
		return nil, err
	}

	return data, data.Validate()
}
