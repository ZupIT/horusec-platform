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
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	emailEntities "github.com/ZupIT/horusec-devkit/pkg/entities/email"
	emailEnums "github.com/ZupIT/horusec-devkit/pkg/enums/email"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
)

func TestNewWorkspaceUseCases(t *testing.T) {
	t.Run("should success create a new use cases", func(t *testing.T) {
		assert.NotNil(t, NewWorkspaceUseCases())
	})
}

func TestWorkspaceDataFromIOReadCloser(t *testing.T) {
	t.Run("should success get workspace data from request body", func(t *testing.T) {
		useCases := NewWorkspaceUseCases()

		data := &workspace.Data{
			AccountID: uuid.New(),
			Name:      "test",
		}

		readCloser, err := parser.ParseEntityToIOReadCloser(data)
		assert.NoError(t, err)

		response, err := useCases.WorkspaceDataFromIOReadCloser(readCloser)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, data.AccountID, response.AccountID)
		assert.Equal(t, data.Name, response.Name)
	})

	t.Run("should return error when failed to parse body to entity", func(t *testing.T) {
		useCases := NewWorkspaceUseCases()

		readCloser, err := parser.ParseEntityToIOReadCloser("")
		assert.NoError(t, err)

		response, err := useCases.WorkspaceDataFromIOReadCloser(readCloser)
		assert.Error(t, err)
		assert.Nil(t, response)
	})
}

func TestFilterAccountWorkspaceByID(t *testing.T) {
	t.Run("should success create a account workspace filter by id", func(t *testing.T) {
		useCases := NewWorkspaceUseCases()
		id := uuid.New()

		filter := useCases.FilterAccountWorkspaceByID(id, id)

		assert.NotPanics(t, func() {
			assert.Equal(t, id, filter["workspace_id"])
			assert.Equal(t, id, filter["account_id"])
		})
	})
}

func TestFilterWorkspaceByID(t *testing.T) {
	t.Run("should success create a account workspace filter by id", func(t *testing.T) {
		useCases := NewWorkspaceUseCases()
		id := uuid.New()

		filter := useCases.FilterWorkspaceByID(id)

		assert.NotPanics(t, func() {
			assert.Equal(t, id, filter["workspace_id"])
		})
	})
}

func TestNewWorkspaceData(t *testing.T) {
	t.Run("should success create a new workspace data", func(t *testing.T) {
		useCases := NewWorkspaceUseCases()
		id := uuid.New()

		accountData := &proto.GetAccountDataResponse{
			AccountID:          id.String(),
			Permissions:        []string{"test"},
			IsApplicationAdmin: true,
		}

		data := useCases.NewWorkspaceData(id, accountData)

		assert.Equal(t, id, data.AccountID)
		assert.Equal(t, id, data.WorkspaceID)
		assert.Equal(t, []string{"test"}, data.Permissions)
		assert.Equal(t, true, data.IsApplicationAdmin)
	})
}

func TestNewOrganizationInviteEmail(t *testing.T) {
	t.Run("should success create a new organization invite email", func(t *testing.T) {
		useCases := NewWorkspaceUseCases()

		emailBytes := useCases.NewOrganizationInviteEmail("test@test.com", "test", "test")
		assert.NotNil(t, emailBytes)
		assert.NotEmpty(t, emailBytes)

		email := &emailEntities.Message{}
		assert.NoError(t, json.Unmarshal(emailBytes, email))

		assert.Equal(t, "test@test.com", email.To)
		assert.Equal(t, emailEnums.OrganizationInvite, email.TemplateName)
		assert.Equal(t, "[Horusec] Organization invite", email.Subject)

		assert.NotPanics(t, func() {
			data := email.Data.(map[string]interface{})

			assert.Equal(t, "test", data["WorkspaceName"])
			assert.Equal(t, "test", data["Username"])
			assert.Equal(t, "http://localhost:8043", data["URL"])
		})
	})
}
