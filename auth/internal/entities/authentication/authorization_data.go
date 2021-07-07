// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
