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

package token

import (
	"github.com/ZupIT/horusec-platform/api/internal/entities/token"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
)

type IToken interface {
	FindTokenByValue(tokenValue string) response.IResponse
}

type Token struct {
	databaseRead database.IDatabaseRead
}

func NewRepositoriesToken(connection *database.Connection) IToken {
	return &Token{
		databaseRead: connection.Read,
	}
}

func (a *Token) FindTokenByValue(tokenValue string) response.IResponse {
	rawSQL := `
		SELECT tokens.token_id as token_id, tokens.repository_id as repository_id,
			repositories.name as repository_name, tokens.workspace_id as workspace_id,
			workspaces.name as workspace_name, tokens.expires_at as expires_at, tokens.is_expirable as is_expirable
		FROM tokens
		INNER JOIN workspaces ON tokens.workspace_id = workspaces.workspace_id  
		LEFT JOIN repositories ON tokens.repository_id = repositories.repository_id
		WHERE value = ?
	`
	return a.databaseRead.Raw(rawSQL, &token.Token{}, tokenValue)
}
