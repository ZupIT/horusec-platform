package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmail(t *testing.T) {
	t.Run("should return no error when valid data", func(t *testing.T) {
		data := &Email{
			Email: "test@test.com",
		}

		assert.NoError(t, data.Validate())
	})

	t.Run("should return error when invalid data", func(t *testing.T) {
		data := &Email{
			Email: "test",
		}

		assert.Error(t, data.Validate())
	})
}

func TestToBytesEmail(t *testing.T) {
	t.Run("should success parse to bytes", func(t *testing.T) {
		data := &Email{
			Email: "test@test.com",
		}

		assert.NotEmpty(t, data.ToBytes())
	})
}
