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
)

func TestValidateRoleData(t *testing.T) {
	t.Run("should return no error when valid data", func(t *testing.T) {
		data := &Data{}

		data.Role = account.Admin
		assert.NoError(t, data.Validate())

		data.Role = account.Supervisor
		assert.NoError(t, data.Validate())

		data.Role = account.Member
		assert.NoError(t, data.Validate())
	})

	t.Run("should return when invalid role value", func(t *testing.T) {
		data := &Data{
			Role: "test",
		}

		assert.Error(t, data.Validate())
	})
}

func TestSetDataIDs(t *testing.T) {
	t.Run("should success set account and workspace id", func(t *testing.T) {
		data := &Data{}
		id := uuid.New()

		_ = data.SetDataIDs(id, id.String(), id.String())
		assert.Equal(t, id, data.WorkspaceID)
		assert.Equal(t, id, data.RepositoryID)
		assert.Equal(t, id, data.AccountID)
	})
}

func TestToBytesRoleData(t *testing.T) {
	t.Run("should success parse to bytes", func(t *testing.T) {
		data := Data{}

		assert.NotEmpty(t, data.ToBytes())
	})
}
