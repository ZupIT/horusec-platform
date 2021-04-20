package authentication

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/utils/crypto"
)

func TestValidate(t *testing.T) {
	t.Run("should return no error when valid credentials", func(t *testing.T) {
		credentials := &LoginCredentials{
			Username: "test",
			Password: "test",
		}

		assert.NoError(t, credentials.Validate())
	})

	t.Run("should return error when invalid credentials", func(t *testing.T) {
		credentials := &LoginCredentials{
			Username: "",
			Password: "",
		}

		assert.Error(t, credentials.Validate())
	})
}

func TestCheckInvalidPassword(t *testing.T) {
	t.Run("should return false when valid password", func(t *testing.T) {
		credentials := LoginCredentials{
			Username: "test",
			Password: "test",
		}

		hash, _ := crypto.HashPasswordBcrypt("test")

		assert.False(t, credentials.CheckInvalidPassword(hash))
	})

	t.Run("should return true when invalid password", func(t *testing.T) {
		credentials := LoginCredentials{
			Username: "test",
			Password: "test",
		}

		hash, _ := crypto.HashPasswordBcrypt("test2")

		assert.True(t, credentials.CheckInvalidPassword(hash))
	})
}

func TestIsInvalidUsernameEmail(t *testing.T) {
	t.Run("should true when invalid email", func(t *testing.T) {
		credentials := &LoginCredentials{
			Username: "test",
		}

		assert.True(t, credentials.IsInvalidUsernameEmail())
	})

	t.Run("should false when valid email", func(t *testing.T) {
		credentials := &LoginCredentials{
			Username: "test@test.com",
		}

		assert.False(t, credentials.IsInvalidUsernameEmail())
	})
}

func TestToBytes(t *testing.T) {
	t.Run("should success parse to bytes", func(t *testing.T) {
		credentials := &LoginCredentials{
			Username: "test",
			Password: "test",
		}

		assert.NotEmpty(t, credentials.ToBytes())
	})
}
