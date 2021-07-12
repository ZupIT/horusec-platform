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
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/utils/crypto"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	entityToken "github.com/ZupIT/horusec-platform/api/internal/entities/token"
	"github.com/ZupIT/horusec-platform/api/internal/repositories/token"

	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
)

func testHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func TestCrypto(t *testing.T) {
	textStr := crypto.GenerateSHA256("bbf6c937-c985-4874-9981-c37c2889b745")
	assert.Equal(t, "25398a3476bb0250be7f367e96cc54f0907ced1549fadda2cca110ee5c89ae5d", textStr)
}

func TestAuthz_IsAuthorized(t *testing.T) {
	t.Run("Should return success when check if token is authorized", func(t *testing.T) {
		repoTokenMock := &token.Mock{}
		data := &entityToken.Token{
			TokenID:        uuid.New(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.NewString(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.NewString(),
			IsExpirable:    false,
		}

		repoTokenMock.On("FindTokenByValue").Return(response.NewResponse(1, nil, data))

		middleware := NewTokenAuthz(repoTokenMock)

		handler := middleware.IsAuthorized(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", uuid.New().String())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("Should return error unauthorized when not exist token on header", func(t *testing.T) {
		repoTokenMock := &token.Mock{}

		middleware := NewTokenAuthz(repoTokenMock)

		handler := middleware.IsAuthorized(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
	t.Run("Should return error unauthorized when return error on find token in database", func(t *testing.T) {
		repoTokenMock := &token.Mock{}
		data := &entityToken.Token{}
		err := errors.New("unexpected error")

		repoTokenMock.On("FindTokenByValue").Return(response.NewResponse(1, err, data))

		middleware := NewTokenAuthz(repoTokenMock)

		handler := middleware.IsAuthorized(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", uuid.New().String())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
	t.Run("Should return error unauthorized when return not found records in database", func(t *testing.T) {
		repoTokenMock := &token.Mock{}

		repoTokenMock.On("FindTokenByValue").Return(response.NewResponse(1, nil, nil))

		middleware := NewTokenAuthz(repoTokenMock)

		handler := middleware.IsAuthorized(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", uuid.New().String())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
	t.Run("Should return error when token was expired", func(t *testing.T) {
		repoTokenMock := &token.Mock{}
		data := &entityToken.Token{
			TokenID:        uuid.New(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.NewString(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.NewString(),
			IsExpirable:    true,
			ExpiresAt:      time.Now().Add(-(time.Duration(24) * time.Hour)),
		}

		repoTokenMock.On("FindTokenByValue").Return(response.NewResponse(1, nil, data))

		middleware := NewTokenAuthz(repoTokenMock)

		handler := middleware.IsAuthorized(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", uuid.New().String())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
	t.Run("Should return success when token is was expired", func(t *testing.T) {
		repoTokenMock := &token.Mock{}
		data := &entityToken.Token{
			TokenID:        uuid.New(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.NewString(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.NewString(),
			IsExpirable:    true,
			ExpiresAt:      time.Now().Add(time.Duration(24) * time.Hour),
		}

		repoTokenMock.On("FindTokenByValue").Return(response.NewResponse(1, nil, data))

		middleware := NewTokenAuthz(repoTokenMock)

		handler := middleware.IsAuthorized(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", uuid.New().String())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
