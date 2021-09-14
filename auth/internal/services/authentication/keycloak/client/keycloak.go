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

package keycloak

import (
	"context"
	"strings"

	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"

	"github.com/pkg/errors"

	"github.com/Nerzal/gocloak/v7"
	"github.com/ZupIT/horusec-devkit/pkg/utils/env"
	"github.com/google/uuid"

	keycloakEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication/keycloak"
)

type IClient interface {
	Authenticate(username, password string) (*gocloak.JWT, error)
	IsActiveToken(token string) (bool, error)
	GetAccountIDByJWTToken(token string) (uuid.UUID, error)
	GetUserInfo(accessToken string) (*gocloak.UserInfo, error)
}

type Client struct {
	ctx          context.Context
	client       gocloak.GoCloak
	clientID     string
	clientSecret string
	realm        string
	totp         bool
}

func NewKeycloakClient() IClient {
	return &Client{
		ctx:          context.Background(),
		client:       gocloak.NewClient(env.GetEnvOrDefault(keycloakEnums.EnvHorusecKeycloakBasePath, "")),
		clientID:     env.GetEnvOrDefault(keycloakEnums.EnvHorusecKeycloakClientID, ""),
		clientSecret: env.GetEnvOrDefault(keycloakEnums.EnvHorusecKeycloakClientSecret, ""),
		realm:        env.GetEnvOrDefault(keycloakEnums.EnvHorusecKeycloakRealm, ""),
		totp:         env.GetEnvOrDefaultBool(keycloakEnums.EnvHorusecKeycloakTOPT, false),
	}
}

func (c *Client) Authenticate(username, password string) (*gocloak.JWT, error) {
	return c.client.LoginOtp(c.ctx, c.clientID, c.clientSecret, c.realm, username, password, "")
}

func (c *Client) IsActiveToken(token string) (bool, error) {
	result, err := c.client.RetrospectToken(c.ctx, c.removeBearer(token), c.clientID, c.clientSecret, c.realm)
	if err != nil {
		return false, errors.Wrap(err, keycloakEnums.MessageFailedToCheckIfTokenIsActive)
	}

	return *result.Active, nil
}

func (c *Client) GetAccountIDByJWTToken(token string) (uuid.UUID, error) {
	userInfo, err := c.GetUserInfo(c.removeBearer(token))
	if err != nil {
		return uuid.Nil, errors.Wrap(err, keycloakEnums.MessageFailedToGetAccountIDFromKeycloakToken)
	}

	return uuid.Parse(*userInfo.Sub)
}

func (c *Client) GetUserInfo(accessToken string) (*gocloak.UserInfo, error) {
	userInfo, err := c.client.GetUserInfo(c.ctx, c.removeBearer(accessToken), c.realm)
	if err != nil {
		logger.LogError("ERROR -> ", err)
		return nil, errors.Wrap(err, keycloakEnums.MessageFailedToGetUserInfo)
	}

	return userInfo, nil
}

func (c *Client) removeBearer(accessToken string) string {
	return strings.ReplaceAll(accessToken, "Bearer ", "")
}
