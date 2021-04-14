package authentication

import (
	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/google/uuid"
)

type AuthorizationData struct {
	Token           string                 `json:"token"`
	Type            auth.AuthorizationType `json:"type"`
	WorkspaceID     uuid.UUID              `json:"workspaceID"`
	RepositoryID    uuid.UUID              `json:"repositoryID"`
	AuthzMember     []string               `json:"authzMember"`
	AuthzAdmin      []string               `json:"authzAdmin"`
	AuthzSupervisor []string               `json:"authzSupervisor"`
}

func (a *AuthorizationData) SetGroups(authzGroups *AuthzGroups) *AuthorizationData {
	a.AuthzMember = authzGroups.AuthzMember
	a.AuthzSupervisor = authzGroups.AuthzSupervisor
	a.AuthzAdmin = authzGroups.AuthzAdmin

	return a
}
