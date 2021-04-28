package authentication

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHorusecAuthenticationService(t *testing.T) {
	t.Run("should success set authz groups", func(t *testing.T) {
		data := AuthorizationData{}

		authzGroups := &AuthzGroups{
			AuthzMember:     []string{"test"},
			AuthzAdmin:      []string{"test"},
			AuthzSupervisor: []string{"test"},
		}

		_ = data.SetGroups(authzGroups)

		assert.Equal(t, authzGroups.AuthzAdmin, data.AuthzAdmin)
		assert.Equal(t, authzGroups.AuthzSupervisor, data.AuthzSupervisor)
		assert.Equal(t, authzGroups.AuthzMember, data.AuthzMember)
	})
}
