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
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-platform/api/internal/entities/token"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
)

func TestToken_FindTokenByValue(t *testing.T) {
	t.Run("Should return token found on database from token workspace", func(t *testing.T) {
		dbMockRead := &database.Mock{}
		dbMockRead.On("Raw").Return(response.NewResponse(1, nil, &token.Token{
			TokenID:        uuid.New(),
			RepositoryID:   uuid.Nil,
			RepositoryName: "",
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.New().String(),
		}))
		connectionMock := &database.Connection{
			Read: dbMockRead,
		}
		res := NewRepositoriesToken(connectionMock).FindTokenByValue(uuid.New().String())
		assert.NoError(t, res.GetError())
		assert.NotEmpty(t, res.GetData())
		assert.NotEqual(t, res.GetData().(*token.Token).TokenID, uuid.Nil)
		assert.NotEqual(t, res.GetData().(*token.Token).WorkspaceID, uuid.Nil)
		assert.Equal(t, res.GetData().(*token.Token).RepositoryID, uuid.Nil)
	})
	t.Run("Should return token found on database from token repository", func(t *testing.T) {
		dbMockRead := &database.Mock{}
		dbMockRead.On("Raw").Return(response.NewResponse(1, nil, &token.Token{
			TokenID:        uuid.New(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.New().String(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.New().String(),
		}))
		connectionMock := &database.Connection{
			Read: dbMockRead,
		}
		res := NewRepositoriesToken(connectionMock).FindTokenByValue(uuid.New().String())
		assert.NoError(t, res.GetError())
		assert.NotEmpty(t, res.GetData())
		assert.NotEqual(t, res.GetData().(*token.Token).TokenID, uuid.Nil)
		assert.NotEqual(t, res.GetData().(*token.Token).WorkspaceID, uuid.Nil)
		assert.NotEqual(t, res.GetData().(*token.Token).RepositoryID, uuid.Nil)
	})
	t.Run("Should return error when not found token on database", func(t *testing.T) {
		dbMockRead := &database.Mock{}
		dbMockRead.On("Raw").Return(response.NewResponse(1, enums.ErrorNotFoundRecords, &token.Token{}))
		connectionMock := &database.Connection{
			Read: dbMockRead,
		}
		res := NewRepositoriesToken(connectionMock).FindTokenByValue(uuid.New().String())
		assert.Error(t, res.GetError())
		assert.Empty(t, res.GetData())
	})
}
