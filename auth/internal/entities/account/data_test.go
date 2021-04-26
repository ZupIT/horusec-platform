package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("should return no error when valid data", func(t *testing.T) {
		data := &Data{
			Email:    "test@test.com",
			Password: "Test@123",
			Username: "test",
		}

		assert.NoError(t, data.Validate())
	})

	t.Run("should return error when invalid data email", func(t *testing.T) {
		data := &Data{
			Email:    "test",
			Password: "Test@123",
			Username: "test",
		}

		assert.Error(t, data.Validate())
	})

	t.Run("should return error when invalid data password", func(t *testing.T) {
		data := &Data{
			Email:    "test@test.com",
			Password: "test",
			Username: "test",
		}

		assert.Error(t, data.Validate())
	})
}

func TestToAccount(t *testing.T) {
	t.Run("should success parse data to account", func(t *testing.T) {
		data := &Data{
			Email:    "test@test.com",
			Password: "Test@123",
			Username: "test",
		}

		account := data.ToAccount()
		assert.Equal(t, data.Email, account.Email)
		assert.NotEmpty(t, account.Password)
		assert.Equal(t, data.Username, account.Username)
	})
}
