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
	"context"
	"net/http"
	"time"

	repositoriesToken "github.com/ZupIT/horusec-platform/api/internal/repositories/token"

	"github.com/ZupIT/horusec-platform/api/internal/entities/token"
	tokensEnums "github.com/ZupIT/horusec-platform/api/internal/middelwares/token/enums"

	enumsDatabase "github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares/enums"
	"github.com/ZupIT/horusec-devkit/pkg/utils/crypto"
	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
)

type ITokenAuthz interface {
	IsAuthorized(next http.Handler) http.Handler
}

type Authz struct {
	repoToken repositoriesToken.IToken
}

func NewTokenAuthz(repoToken repositoriesToken.IToken) ITokenAuthz {
	return &Authz{
		repoToken: repoToken,
	}
}

func (a *Authz) IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenValue, err := a.getTokenHashFromAuthorizationHeader(r)
		if err != nil {
			httpUtil.StatusUnauthorized(w, enums.ErrorUnauthorized)
			return
		}
		ctx, err := a.getContextAndValidateIsValidToken(tokenValue, r)
		if err != nil {
			a.verifyValidateTokenErrors(w, err)
			return
		}
		a.showCLIVersion(r)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *Authz) getTokenHashFromAuthorizationHeader(r *http.Request) (string, error) {
	tokenStr := r.Header.Get("X-Horusec-Authorization")
	if tokenStr == "" {
		return "", enums.ErrorUnauthorized
	}

	return crypto.GenerateSHA256(tokenStr), nil
}

func (a *Authz) getContextAndValidateIsValidToken(
	tokenValue string, r *http.Request) (context.Context, error) {
	ctx := r.Context()
	res := a.repoToken.FindTokenByValue(tokenValue)
	if err := res.GetError(); err != nil {
		return nil, err
	}
	if res.GetData() == nil {
		return nil, enumsDatabase.ErrorNotFoundRecords
	}
	tokenFound := res.GetData().(*token.Token)
	ctx = a.bindTokenInCtx(ctx, tokenFound)
	return ctx, a.returnErrorIfTokenIsExpired(tokenFound)
}

func (a *Authz) bindTokenInCtx(ctx context.Context, tokenFound *token.Token) context.Context {
	newCtx := context.WithValue(ctx, tokensEnums.RepositoryID, tokenFound.RepositoryID)
	newCtx = context.WithValue(newCtx, tokensEnums.RepositoryName, tokenFound.RepositoryName)
	newCtx = context.WithValue(newCtx, tokensEnums.WorkspaceID, tokenFound.WorkspaceID)
	return context.WithValue(newCtx, tokensEnums.WorkspaceName, tokenFound.WorkspaceName)
}

func (a *Authz) returnErrorIfTokenIsExpired(tokenFound *token.Token) error {
	if !tokenFound.IsExpirable {
		return nil
	}
	if tokenFound.ExpiresAt.Before(time.Now()) {
		return tokensEnums.ErrorTokenExpired
	}
	return nil
}

func (a *Authz) showCLIVersion(r *http.Request) {
	cliVersion := r.Header.Get("X-Horusec-CLI-Version")
	logger.LogInfo("Current Horusec-CLI version is: " + cliVersion)
}

func (a *Authz) verifyValidateTokenErrors(w http.ResponseWriter, err error) {
	if err == tokensEnums.ErrorTokenExpired {
		httpUtil.StatusUnauthorized(w, tokensEnums.ErrorTokenExpired)
		return
	}
	logger.LogError("{HORUSEC_API}", err)

	httpUtil.StatusUnauthorized(w, enums.ErrorUnauthorized)
}
