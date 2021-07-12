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
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/core/internal/entities/role"
)

func TestNewRoleData(t *testing.T) {
	t.Run("should success create a new workspace data", func(t *testing.T) {
		useCases := NewRoleUseCases()
		id := uuid.New()

		data := useCases.NewRoleData(id, id, id)

		assert.Equal(t, id, data.AccountID)
		assert.Equal(t, id, data.WorkspaceID)
		assert.Equal(t, id, data.RepositoryID)
	})
}

func TestInviteUserDataFromIOReadCloser(t *testing.T) {
	t.Run("should success invite user data from request body", func(t *testing.T) {
		useCases := NewRoleUseCases()
		id := uuid.New()

		data := &role.UserData{
			Role:         "admin",
			Email:        "test@test.com",
			AccountID:    id,
			WorkspaceID:  id,
			RepositoryID: id,
			Username:     "test",
		}

		readCloser, err := parser.ParseEntityToIOReadCloser(data)
		assert.NoError(t, err)

		data, err = useCases.InviteUserDataFromIOReadCloser(readCloser)

		assert.NoError(t, err)
		assert.Equal(t, account.Admin, data.Role)
		assert.Equal(t, "test@test.com", data.Email)
		assert.Equal(t, "test", data.Username)
		assert.Equal(t, id, data.AccountID)
		assert.Equal(t, id, data.WorkspaceID)
		assert.Equal(t, id, data.RepositoryID)
	})

	t.Run("should return error when failed to parse body to entity", func(t *testing.T) {
		useCases := NewRoleUseCases()

		readCloser, err := parser.ParseEntityToIOReadCloser("")
		assert.NoError(t, err)

		response, err := useCases.InviteUserDataFromIOReadCloser(readCloser)
		assert.Error(t, err)
		assert.Nil(t, response)
	})
}

func TestRoleDataFromIOReadCloser(t *testing.T) {
	t.Run("should success get role data from request body", func(t *testing.T) {
		useCases := NewRoleUseCases()
		id := uuid.New()

		data := &role.Data{
			AccountID:    id,
			WorkspaceID:  id,
			RepositoryID: id,
			Role:         "admin",
		}

		readCloser, err := parser.ParseEntityToIOReadCloser(data)
		assert.NoError(t, err)

		response, err := useCases.RoleDataFromIOReadCloser(readCloser)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, data.AccountID, response.AccountID)
		assert.Equal(t, data.WorkspaceID, response.WorkspaceID)
		assert.Equal(t, data.RepositoryID, response.RepositoryID)
		assert.Equal(t, data.Role, response.Role)
	})

	t.Run("should return error when failed to parse body to entity", func(t *testing.T) {
		useCases := NewRoleUseCases()

		readCloser, err := parser.ParseEntityToIOReadCloser("")
		assert.NoError(t, err)

		response, err := useCases.RoleDataFromIOReadCloser(readCloser)
		assert.Error(t, err)
		assert.Nil(t, response)
	})
}
