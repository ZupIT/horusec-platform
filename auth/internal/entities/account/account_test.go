package account

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestToResponse(t *testing.T) {
	t.Run("should success parse to account response", func(t *testing.T) {
		account := &Account{
			AccountID:          uuid.New(),
			Email:              "test@test.com",
			Username:           "test",
			IsConfirmed:        true,
			IsApplicationAdmin: true,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}

		response := account.ToResponse()
		assert.Equal(t, account.AccountID, response.AccountID)
		assert.Equal(t, account.Email, response.Email)
		assert.Equal(t, account.Username, response.Username)
		assert.Equal(t, account.IsConfirmed, response.IsConfirmed)
		assert.Equal(t, account.IsApplicationAdmin, response.IsApplicationAdmin)
		assert.Equal(t, account.CreatedAt, response.CreatedAt)
		assert.Equal(t, account.UpdatedAt, response.UpdatedAt)
	})
}

func TestIsNotConfirmed(t *testing.T) {
	t.Run("should true when account is not confirmed", func(t *testing.T) {
		account := &Account{
			IsConfirmed: false,
		}

		assert.True(t, account.IsNotConfirmed())
	})

	t.Run("should false when account is confirmed", func(t *testing.T) {
		account := &Account{
			IsConfirmed: true,
		}

		assert.False(t, account.IsNotConfirmed())
	})
}

func TestToTokenData(t *testing.T) {
	t.Run("should success parse to token data", func(t *testing.T) {
		account := &Account{
			AccountID:          uuid.New(),
			Email:              "test@test.com",
			Username:           "test",
			IsConfirmed:        true,
			IsApplicationAdmin: true,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}

		response := account.ToTokenData()
		assert.Equal(t, account.AccountID, response.AccountID)
		assert.Equal(t, account.Email, response.Email)
		assert.Equal(t, account.Username, response.Username)
	})
}

func TestToLoginResponse(t *testing.T) {
	t.Run("should success parse to account response", func(t *testing.T) {
		expiresAt := time.Now()

		account := &Account{
			AccountID:          uuid.New(),
			Email:              "test@test.com",
			Username:           "test",
			IsConfirmed:        true,
			IsApplicationAdmin: true,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}

		response := account.ToLoginResponse("test", "test", expiresAt)
		assert.Equal(t, "test", response.RefreshToken)
		assert.Equal(t, "test", response.AccessToken)
		assert.Equal(t, expiresAt, response.ExpiresAt)
		assert.Equal(t, account.AccountID, response.AccountID)
		assert.Equal(t, account.Username, response.Username)
		assert.Equal(t, account.Email, response.Email)
		assert.Equal(t, account.IsApplicationAdmin, response.IsApplicationAdmin)

	})
}

func TestHashPassword(t *testing.T) {
	t.Run("should success parse to account response", func(t *testing.T) {
		account := &Account{
			Password: "test",
		}

		account.HashPassword()
		assert.NotEqual(t, "test", account.Password)
	})
}

func TestSetNewAccountData(t *testing.T) {
	t.Run("should success set new account default data", func(t *testing.T) {
		account := &Account{}

		_ = account.SetNewAccountData()
		assert.NotEqual(t, uuid.UUID{}, account.AccountID)
		assert.NotEmpty(t, account.Password)
		assert.NotEqual(t, time.Time{}, account.UpdatedAt)
		assert.NotEqual(t, time.Time{}, account.CreatedAt)
	})
}

func TestToGetAccountDataResponse(t *testing.T) {
	t.Run("should success parse to grpc struct get account data response", func(t *testing.T) {
		account := &Account{
			AccountID:          uuid.New(),
			IsApplicationAdmin: true,
		}

		response := account.ToGetAccountDataResponse([]string{"test"})
		assert.Equal(t, account.AccountID.String(), response.AccountID)
		assert.Equal(t, account.IsApplicationAdmin, response.IsApplicationAdmin)
		assert.Equal(t, []string{"test"}, response.Permissions)
	})
}
