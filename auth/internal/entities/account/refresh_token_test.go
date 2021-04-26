package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateRefreshToken(t *testing.T) {
	t.Run("should return no error when valid token", func(t *testing.T) {
		data := &RefreshToken{
			RefreshToken: "test",
		}

		assert.NoError(t, data.Validate())
	})
}
