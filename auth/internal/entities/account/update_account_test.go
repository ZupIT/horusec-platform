package account

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidateUpdateAccount(t *testing.T) {
	t.Run("should return no error when valid data", func(t *testing.T) {
		data := &UpdateAccount{
			Email:    "test@test.com",
			Username: "test",
		}

		assert.NoError(t, data.Validate())
	})

	t.Run("should return error when invalid email", func(t *testing.T) {
		data := &UpdateAccount{
			Email:    "test",
			Username: "test",
		}

		assert.Error(t, data.Validate())
	})

	t.Run("should return error when invalid username", func(t *testing.T) {
		data := &UpdateAccount{
			Email:    "test@test.com",
			Username: "",
		}

		assert.Error(t, data.Validate())
	})
}

func TestSetAccountIDAndIsConfirmed(t *testing.T) {
	t.Run("should success set account id and is confirmed", func(t *testing.T) {
		id := uuid.New()
		data := &UpdateAccount{}

		data.SetAccountIDAndIsConfirmed(id, true)
		assert.Equal(t, id, data.AccountID)
		assert.Equal(t, true, data.IsConfirmed)
	})
}

func TestHasEmailChange(t *testing.T) {
	t.Run("should true when contains email changes", func(t *testing.T) {
		data := &UpdateAccount{
			Email: "test@test.com",
		}

		assert.True(t, data.HasEmailChange("test@test1.com"))
	})

	t.Run("should false when not contains email changes", func(t *testing.T) {
		data := &UpdateAccount{
			Email: "test@test.com",
		}

		assert.False(t, data.HasEmailChange("test@test.com"))
	})
}
