package account

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewAccountUseCases(t *testing.T) {
	t.Run("should success create a new use cases", func(t *testing.T) {
		assert.NotNil(t, NewAccountUseCases())
	})
}

func TestFilterAccountByID(t *testing.T) {
	t.Run("should success create a filter by account id", func(t *testing.T) {
		useCases := NewAccountUseCases()

		id := uuid.New()

		filter := useCases.FilterAccountByID(id)
		assert.NotPanics(t, func() {
			assert.Equal(t, id, filter["account_id"])
		})
	})
}

func TestFilterAccountByEmail(t *testing.T) {
	t.Run("should success create a filter by account id", func(t *testing.T) {
		useCases := NewAccountUseCases()

		filter := useCases.FilterAccountByEmail("test@test.com")
		assert.NotPanics(t, func() {
			assert.Equal(t, "test@test.com", filter["email"])
		})
	})
}

func TestFilterAccountByUsername(t *testing.T) {
	t.Run("should success create a filter by account id", func(t *testing.T) {
		useCases := NewAccountUseCases()

		filter := useCases.FilterAccountByUsername("test")
		assert.NotPanics(t, func() {
			assert.Equal(t, "test", filter["username"])
		})
	})
}
