package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
)

func TestIsInvalidLdapGroups(t *testing.T) {
	t.Run("should return false when valid group and ldap auth", func(t *testing.T) {
		assert.False(t, IsInvalidLdapGroups(auth.Ldap, []string{"test"}, []string{"test"}))
	})

	t.Run("should return true when invalid group and ldap auth", func(t *testing.T) {
		assert.True(t, IsInvalidLdapGroups(auth.Ldap, []string{"test"}, []string{"test2"}))
	})

	t.Run("should return true when invalid group and ignore empty", func(t *testing.T) {
		assert.True(t, IsInvalidLdapGroups(auth.Ldap, []string{""}, []string{"", "test"}))
	})

	t.Run("should return false when invalid group and ignore empty", func(t *testing.T) {
		assert.False(t, IsInvalidLdapGroups(auth.Ldap, []string{"", "test"}, []string{"", "test"}))
	})

	t.Run("should return false when not ldap auth", func(t *testing.T) {
		assert.False(t, IsInvalidLdapGroups(auth.Horusec, []string{"test"}, []string{"test"}))
	})
}
