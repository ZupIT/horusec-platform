package token

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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
