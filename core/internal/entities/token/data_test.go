package token

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("should return no error when valid data", func(t *testing.T) {
		data := &Data{
			Description: "test",
			IsExpirable: true,
			ExpiresAt:   time.Date(9999, 1, 1, 1, 1, 1, 1, time.UTC),
		}

		assert.NoError(t, data.Validate())
	})

	t.Run("should return error when invalid expires at", func(t *testing.T) {
		data := &Data{
			Description: "test",
			IsExpirable: true,
			ExpiresAt:   time.Date(1997, 1, 1, 1, 1, 1, 1, time.UTC),
		}

		assert.Error(t, data.Validate())
	})

	t.Run("should return error missing description", func(t *testing.T) {
		data := &Data{
			IsExpirable: true,
			ExpiresAt:   time.Date(9999, 1, 1, 1, 1, 1, 1, time.UTC),
		}

		assert.Error(t, data.Validate())
	})
}

func TestSetWorkspaceID(t *testing.T) {
	t.Run("should success set workspace id", func(t *testing.T) {
		data := &Data{}
		id := uuid.New()

		_ = data.SetWorkspaceID(id)
		assert.Equal(t, id, data.WorkspaceID)
	})
}

func TestSetIDs(t *testing.T) {
	t.Run("should success set workspace, repository and token id", func(t *testing.T) {
		data := &Data{}
		id := uuid.New()

		_ = data.SetIDs(id, id, id)
		assert.Equal(t, id, data.WorkspaceID)
		assert.Equal(t, id, data.RepositoryID)
		assert.Equal(t, id, data.TokenID)
	})
}

func TestToToken(t *testing.T) {
	t.Run("should success parse token data to token", func(t *testing.T) {
		data := &Data{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			Description:  "test",
			IsExpirable:  true,
			ExpiresAt:    time.Date(9999, 1, 1, 1, 1, 1, 1, time.UTC),
		}

		token, tokenString := data.ToToken()

		assert.NotEmpty(t, tokenString)
		assert.NotEmpty(t, token.Value)
		assert.NotEmpty(t, token.SuffixValue)
		assert.Equal(t, &data.RepositoryID, token.RepositoryID)
		assert.Equal(t, data.WorkspaceID, token.WorkspaceID)
		assert.Equal(t, data.Description, token.Description)
		assert.Equal(t, data.IsExpirable, token.IsExpirable)
		assert.Equal(t, data.ExpiresAt, token.ExpiresAt)
	})

	t.Run("should success parse token data to token with nil repository id", func(t *testing.T) {
		data := &Data{
			RepositoryID: uuid.UUID{},
			WorkspaceID:  uuid.New(),
			Description:  "test",
			IsExpirable:  true,
			ExpiresAt:    time.Date(9999, 1, 1, 1, 1, 1, 1, time.UTC),
		}

		token, tokenString := data.ToToken()

		assert.NotEmpty(t, tokenString)
		assert.NotEmpty(t, token.Value)
		assert.NotEmpty(t, token.SuffixValue)
		assert.Nil(t, token.RepositoryID)
		assert.Equal(t, data.WorkspaceID, token.WorkspaceID)
		assert.Equal(t, data.Description, token.Description)
		assert.Equal(t, data.IsExpirable, token.IsExpirable)
		assert.Equal(t, data.ExpiresAt, token.ExpiresAt)
	})
}

func TestToBytes(t *testing.T) {
	t.Run("should parse data to bytes", func(t *testing.T) {
		data := &Data{Description: "test"}

		assert.NotEmpty(t, data.ToByes())
	})
}
