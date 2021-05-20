package cors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCorsConfig(t *testing.T) {
	t.Run("should success create a new cors config", func(t *testing.T) {
		config := NewCorsConfig()
		assert.Equal(t, []string{"*"}, config.AllowedOrigins)
		assert.Equal(t, []string{"GET", "POST", "DELETE", "OPTIONS", "PATCH"}, config.AllowedMethods)
		assert.Equal(t, []string{"Accept", "headers", "X-Horusec-Authorization", "Content-Type"}, config.AllowedHeaders)
		assert.True(t, config.AllowCredentials)
		assert.Equal(t, 300, config.MaxAge)
	})
}
