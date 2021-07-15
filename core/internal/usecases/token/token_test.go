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
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/core/internal/entities/token"
)

func TestNewTokenUseCases(t *testing.T) {
	t.Run("should success create a new use cases", func(t *testing.T) {
		assert.NotNil(t, NewTokenUseCases())
	})
}

func TestWorkspaceDataFromIOReadCloser(t *testing.T) {
	t.Run("should success get workspace data from request body", func(t *testing.T) {
		useCases := NewTokenUseCases()

		data := &token.Data{
			Description: "test",
			IsExpirable: true,
			ExpiresAt:   time.Date(9999, 1, 1, 1, 1, 1, 1, time.UTC),
		}

		readCloser, err := parser.ParseEntityToIOReadCloser(data)
		assert.NoError(t, err)

		response, err := useCases.TokenDataFromIOReadCloser(readCloser)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, data.Description, response.Description)
		assert.Equal(t, data.IsExpirable, response.IsExpirable)
		assert.Equal(t, data.ExpiresAt, response.ExpiresAt)
	})

	t.Run("should return error when failed to parse body to entity", func(t *testing.T) {
		useCases := NewTokenUseCases()

		readCloser, err := parser.ParseEntityToIOReadCloser("")
		assert.NoError(t, err)

		response, err := useCases.TokenDataFromIOReadCloser(readCloser)
		assert.Error(t, err)
		assert.Nil(t, response)
	})
}

func TestFilterWorkspaceTokenByID(t *testing.T) {
	t.Run("should success create a token workspace filter by id", func(t *testing.T) {
		useCases := NewTokenUseCases()
		id := uuid.New()

		filter := useCases.FilterWorkspaceTokenByID(id, id)

		assert.NotPanics(t, func() {
			assert.Equal(t, id, filter["workspace_id"])
			assert.Equal(t, id, filter["token_id"])
			assert.Equal(t, nil, filter["repository_id"])
		})
	})
}

func TestFilterRepositoryTokenByID(t *testing.T) {
	t.Run("should success create a repository token filter by id", func(t *testing.T) {
		useCases := NewTokenUseCases()
		id := uuid.New()

		filter := useCases.FilterRepositoryTokenByID(id, id, id)

		assert.NotPanics(t, func() {
			assert.Equal(t, id, filter["workspace_id"])
			assert.Equal(t, id, filter["repository_id"])
			assert.Equal(t, id, filter["token_id"])
		})
	})
}

func TestFilterListWorkspaceTokens(t *testing.T) {
	t.Run("should success create a repository token filter by id", func(t *testing.T) {
		useCases := NewTokenUseCases()
		id := uuid.New()

		filter := useCases.FilterListWorkspaceTokens(id)

		assert.NotPanics(t, func() {
			assert.Equal(t, id, filter["workspace_id"])
			assert.Equal(t, nil, filter["repository_id"])
		})
	})
}

func TestFilterListRepositoryTokens(t *testing.T) {
	t.Run("should success create a repository token filter by id", func(t *testing.T) {
		useCases := NewTokenUseCases()
		id := uuid.New()

		filter := useCases.FilterListRepositoryTokens(id, id)

		assert.NotPanics(t, func() {
			assert.Equal(t, id, filter["workspace_id"])
			assert.Equal(t, id, filter["repository_id"])
		})
	})
}

func TestNewTokenData(t *testing.T) {
	t.Run("should success create a new token data", func(t *testing.T) {
		useCases := NewTokenUseCases()
		id := uuid.New()

		data := useCases.NewTokenData(id, id.String(), id.String())

		assert.Equal(t, id, data.TokenID)
		assert.Equal(t, id, data.RepositoryID)
		assert.Equal(t, id, data.WorkspaceID)
	})
}
