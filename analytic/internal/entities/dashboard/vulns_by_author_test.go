package dashboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToResponseByAuthor(t *testing.T) {
	t.Run("should success parse", func(t *testing.T) {
		assert.NotNil(t, (&VulnerabilitiesByAuthor{}).ToResponseByAuthor())
	})
}
