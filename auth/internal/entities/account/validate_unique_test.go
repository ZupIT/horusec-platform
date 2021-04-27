package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCheckEmailAndUsername(t *testing.T) {
	t.Run("should return no error when valid data", func(t *testing.T) {
		data := &CheckEmailAndUsername{
			Email:    "test@test.com",
			Username: "test",
		}

		assert.NoError(t, data.Validate())
	})

	t.Run("should return error when invalid email", func(t *testing.T) {
		data := &CheckEmailAndUsername{
			Email:    "test",
			Username: "test",
		}

		assert.Error(t, data.Validate())
	})

	t.Run("should return error when invalid username", func(t *testing.T) {
		data := &CheckEmailAndUsername{
			Email:    "test@test.com",
			Username: "",
		}

		assert.Error(t, data.Validate())
	})
}

func TestToBytesCheckEmailAndUsername(t *testing.T) {
	t.Run("should success parse to bytes", func(t *testing.T) {
		data := &CheckEmailAndUsername{
			Email:    "test@test.com",
			Username: "test",
		}

		assert.NotEmpty(t, data.ToBytes())
	})
}
