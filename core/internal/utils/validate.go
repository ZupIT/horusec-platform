package utils

import (
	"strings"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
)

func IsInvalidLdapGroups(authType auth.AuthorizationType, groups, permissions []string) bool {
	if authType == auth.Ldap && isInvalidGroup(groups, permissions) {
		return true
	}

	return false
}

func isInvalidGroup(groups, permissions []string) bool {
	for _, group := range groups {
		if group == "" {
			continue
		}

		for _, permission := range permissions {
			if strings.TrimSpace(group) == permission {
				return false
			}
		}
	}

	return true
}
