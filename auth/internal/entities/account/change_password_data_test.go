package account

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidateChangePasswordData(t *testing.T) {
	t.Run("should return no error when valid data", func(t *testing.T) {
		data := &ChangePasswordData{
			Password: "Test@213",
		}

		assert.NoError(t, data.Validate())
	})

	t.Run("should return error when invalid data", func(t *testing.T) {
		data := &ChangePasswordData{
			Password: "test",
		}

		assert.Error(t, data.Validate())
	})
}

func TestSetAccountID(t *testing.T) {
	t.Run("should success set account id", func(t *testing.T) {
		id := uuid.New()

		data := &ChangePasswordData{}

		_ = data.SetAccountID(id)
		assert.Equal(t, id, data.AccountID)
	})
}
