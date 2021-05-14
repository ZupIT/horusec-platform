package cors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCorsConfig(t *testing.T) {
	t.Run("should create instance of cors check if is not empty", func(t *testing.T) {
		assert.NotEmpty(t, NewCorsConfig())
	})
}
