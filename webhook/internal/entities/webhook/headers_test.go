package webhook

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderType(t *testing.T) {
	t.Run("Should get value with success", func(t *testing.T) {
		expectedBytes, err := json.Marshal(HeaderType{
			Headers{Key: "key", Value: "value"},
		})
		assert.NoError(t, err)
		headers := HeaderType{
			Headers{Key: "key", Value: "value"},
		}
		valueBytes, err := headers.Value()
		assert.NoError(t, err)
		assert.Equal(t, expectedBytes, valueBytes)
	})
	t.Run("Should scan value with error when data type is wrong", func(t *testing.T) {
		headers := &HeaderType{}
		assert.Error(t, headers.Scan("wrong data type"))
		assert.Empty(t, headers)
	})
	t.Run("Should scan value with success", func(t *testing.T) {
		HeaderTypeBytes, err := json.Marshal(HeaderType{
			Headers{Key: "key", Value: "value"},
		})
		assert.NoError(t, err)
		headers := &HeaderType{}
		assert.NoError(t, headers.Scan(HeaderTypeBytes))
		assert.NotEmpty(t, headers)
	})
	t.Run("Should parse headers to map with success", func(t *testing.T) {
		headers := HeaderType{
			Headers{Key: "key1", Value: "value1"},
			Headers{Key: "key2", Value: "value2"},
		}
		assert.Equal(t, "value1", headers.GetMapHeaders()["key1"])
		assert.Equal(t, "value2", headers.GetMapHeaders()["key2"])
	})
}
