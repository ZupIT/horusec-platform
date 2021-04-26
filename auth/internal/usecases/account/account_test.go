package account

import (
	"testing"

	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-platform/auth/config/app"
)

func TestNewAccountUseCases(t *testing.T) {
	t.Run("should success create a new use cases", func(t *testing.T) {
		assert.NotNil(t, NewAccountUseCases(app.NewAuthAppConfig()))
	})
}

func TestFilterAccountByID(t *testing.T) {
	t.Run("should success create a filter by account id", func(t *testing.T) {
		useCases := NewAccountUseCases(app.NewAuthAppConfig())

		id := uuid.New()

		filter := useCases.FilterAccountByID(id)
		assert.NotPanics(t, func() {
			assert.Equal(t, id, filter["account_id"])
		})
	})
}

func TestFilterAccountByEmail(t *testing.T) {
	t.Run("should success create a filter by account id", func(t *testing.T) {
		useCases := NewAccountUseCases(app.NewAuthAppConfig())

		filter := useCases.FilterAccountByEmail("test@test.com")
		assert.NotPanics(t, func() {
			assert.Equal(t, "test@test.com", filter["email"])
		})
	})
}

func TestFilterAccountByUsername(t *testing.T) {
	t.Run("should success create a filter by account id", func(t *testing.T) {
		useCases := NewAccountUseCases(app.NewAuthAppConfig())

		filter := useCases.FilterAccountByUsername("test")
		assert.NotPanics(t, func() {
			assert.Equal(t, "test", filter["username"])
		})
	})
}

func TestLoginCredentialsFromIOReadCloser(t *testing.T) {
	t.Run("should success get data from request body", func(t *testing.T) {
		useCases := NewAccountUseCases(app.NewAuthAppConfig())

		data := map[string]string{"accessToken": "test"}

		readCloser, err := parser.ParseEntityToIOReadCloser(data)
		assert.NoError(t, err)

		response, err := useCases.AccessTokenFromIOReadCloser(readCloser)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, "test", response.AccessToken)
	})

	t.Run("should return error when failed to parse body to entity", func(t *testing.T) {
		useCases := NewAccountUseCases(app.NewAuthAppConfig())

		readCloser, err := parser.ParseEntityToIOReadCloser("")
		assert.NoError(t, err)

		response, err := useCases.AccessTokenFromIOReadCloser(readCloser)
		assert.Error(t, err)
		assert.Empty(t, response)
	})
}
