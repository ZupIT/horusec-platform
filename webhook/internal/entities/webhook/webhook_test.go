package webhook

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestWebhook(t *testing.T) {
	t.Run("Should get table of webhook", func(t *testing.T) {
		wh := &Webhook{}
		assert.Equal(t, "webhooks", wh.GetTable())
	})
	t.Run("Should generate webhook ID", func(t *testing.T) {
		wh := &Webhook{}
		assert.NotEqual(t, uuid.Nil, wh.GenerateID())
	})
	t.Run("Should generate createdAt", func(t *testing.T) {
		wh := &Webhook{}
		assert.NotEqual(t, time.Time{}, wh.GenerateCreateAt())
	})
	t.Run("Should generate updatedAt", func(t *testing.T) {
		wh := &Webhook{}
		assert.NotEqual(t, time.Time{}, wh.GenerateUpdatedAt())
	})
}
