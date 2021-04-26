package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateResetCodeData(t *testing.T) {
	t.Run("should return no error when valid reset code data", func(t *testing.T) {
		data := &ResetCodeData{
			Email: "test@test.com",
			Code:  "123456",
		}

		assert.NoError(t, data.Validate())
	})

	t.Run("should return error when invalid code", func(t *testing.T) {
		data := &ResetCodeData{
			Email: "test@test.com",
			Code:  "test",
		}

		assert.Error(t, data.Validate())
	})

	t.Run("should return error when invalid email", func(t *testing.T) {
		data := &ResetCodeData{
			Email: "test",
			Code:  "123456",
		}

		assert.Error(t, data.Validate())
	})
}
