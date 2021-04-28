package app

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
)

func TestNewAuthAppConfig(t *testing.T) {
	t.Run("should success create a new config", func(t *testing.T) {
		assert.NotNil(t, NewAuthAppConfig())
	})
}

func TestGetAuthType(t *testing.T) {
	t.Run("should success create a new config", func(t *testing.T) {
		appConfig := NewAuthAppConfig()

		assert.Equal(t, auth.Horusec, appConfig.GetAuthenticationType())
	})
}

func TestToConfigResponse(t *testing.T) {
	t.Run("should success parse config to response", func(t *testing.T) {
		appConfig := NewAuthAppConfig()

		result := appConfig.ToConfigResponse()
		assert.NotPanics(t, func() {
			assert.Equal(t, false, result["enableApplicationAdmin"])
			assert.Equal(t, auth.Horusec, result["authType"])
			assert.Equal(t, false, result["disableBroker"])
		})
	})
}

func TestIsApplicationAdminEnabled(t *testing.T) {
	t.Run("should return false when not active", func(t *testing.T) {
		appConfig := NewAuthAppConfig()

		assert.False(t, appConfig.IsApplicationAdmEnabled())
	})
}

func TestIsDisableBroker(t *testing.T) {
	t.Run("should return false when not active", func(t *testing.T) {
		appConfig := NewAuthAppConfig()

		assert.False(t, appConfig.IsBrokerDisabled())
	})
}

func TestToGetAuthConfigResponse(t *testing.T) {
	t.Run("should return false when not active", func(t *testing.T) {
		appConfig := NewAuthAppConfig()

		result := appConfig.ToGetAuthConfigResponse()
		assert.Equal(t, false, result.EnableApplicationAdmin)
		assert.Equal(t, false, result.DisableBroker)
		assert.Equal(t, auth.Horusec.ToString(), result.AuthType)
	})
}

func TestGetHorusecAuthURL(t *testing.T) {
	t.Run("should success get auth url", func(t *testing.T) {
		appConfig := NewAuthAppConfig()

		assert.Equal(t, "http://localhost:8006", appConfig.GetHorusecAuthURL())
	})
}

func TestGetHorusecManagerURL(t *testing.T) {
	t.Run("should success get manager url", func(t *testing.T) {
		appConfig := NewAuthAppConfig()

		assert.Equal(t, "http://localhost:8043", appConfig.GetHorusecManagerURL())
	})
}
