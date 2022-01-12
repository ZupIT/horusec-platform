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
	"errors"
	"testing"

	"github.com/Nerzal/gocloak/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewKeycloakClient(t *testing.T) {
	t.Run("should return a new horusec keycloak client", func(t *testing.T) {
		assert.IsType(t, NewKeycloakClient(), &Client{})
	})
}

func TestAuthenticate(t *testing.T) {
	t.Run("should login with success in keycloak", func(t *testing.T) {
		token := &gocloak.JWT{
			AccessToken:      "access_token",
			IDToken:          uuid.New().String(),
			ExpiresIn:        15,
			RefreshExpiresIn: 15,
			RefreshToken:     "refresh_token",
			TokenType:        "unique",
		}

		goCloakMock := &GoCloakMock{}
		goCloakMock.On("LoginOtp").Return(token, nil)

		service := &Client{
			ctx:    context.Background(),
			client: goCloakMock,
		}

		resp, err := service.Authenticate("root", "root")
		assert.NoError(t, err)
		assert.NotNil(t, resp.AccessToken)
		assert.Equal(t, "access_token", resp.AccessToken)
	})
}

func TestGetAccountIDByJWTToken(t *testing.T) {
	t.Run("should success get account id without errors", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI4NDc3ZDdmYy0wOTFlLTQwZWEtYjJkMC04ZTg0YWM0Y2Q5ZDQiLCJuYW1lIjoiVGVzdGUiLCJpYXQiOjE1MTYyMzkwMjJ9.HbLKk9hkWw_nGPNwststdFrEjqbQQpDdpQb42KKSVLM"

		goCloakMock := &GoCloakMock{}

		service := &Client{
			ctx:    context.Background(),
			client: goCloakMock,
		}

		userID, err := service.GetAccountIDByJWTToken(token)
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, userID)
	})

	t.Run("should return error when failed to get user info", func(t *testing.T) {
		valid := true
		sub := uuid.New().String()

		goCloakMock := &GoCloakMock{}
		goCloakMock.On("RetrospectToken").Return(&gocloak.RetrospecTokenResult{Active: &valid}, nil)
		goCloakMock.On("GetUserInfo").Return(&gocloak.UserInfo{Sub: &sub}, errors.New("some error"))

		service := &Client{
			ctx:    context.Background(),
			client: goCloakMock,
		}

		userID, err := service.GetAccountIDByJWTToken("")
		assert.Error(t, err)
		assert.Equal(t, uuid.Nil, userID)
	})
}

func TestIsActiveToken(t *testing.T) {
	t.Run("should return true and no errors when success", func(t *testing.T) {
		active := true

		goCloakMock := &GoCloakMock{}
		goCloakMock.On("RetrospectToken").Return(&gocloak.RetrospecTokenResult{Active: &active}, nil)

		service := &Client{
			ctx:    context.Background(),
			client: goCloakMock,
		}

		isActive, err := service.IsActiveToken("")
		assert.NoError(t, err)
		assert.True(t, isActive)
	})

	t.Run("should return false when inactive token", func(t *testing.T) {
		active := false

		goCloakMock := &GoCloakMock{}
		goCloakMock.On("RetrospectToken").Return(&gocloak.RetrospecTokenResult{Active: &active}, nil)

		service := &Client{
			ctx:    context.Background(),
			client: goCloakMock,
		}

		isActive, err := service.IsActiveToken("")
		assert.NoError(t, err)
		assert.False(t, isActive)
	})

	t.Run("should return error when failed to verify if token is valid", func(t *testing.T) {
		goCloakMock := &GoCloakMock{}
		goCloakMock.On("RetrospectToken").Return(&gocloak.RetrospecTokenResult{}, errors.New("error"))

		service := &Client{
			ctx:    context.Background(),
			client: goCloakMock,
		}

		_, err := service.IsActiveToken("")
		assert.Error(t, err)
	})
}

func TestGetUserInfo(t *testing.T) {
	t.Run("should success get user info", func(t *testing.T) {
		email := "test@horusec.com"

		goCloakMock := &GoCloakMock{}
		goCloakMock.On("GetUserInfo").Return(&gocloak.UserInfo{Email: &email}, nil)

		service := &Client{
			ctx:    context.Background(),
			client: goCloakMock,
		}

		user, err := service.GetUserInfo("access_token")
		assert.NoError(t, err)
		assert.Equal(t, email, *user.Email)
	})

	t.Run("should return error when failed to get user info", func(t *testing.T) {
		goCloakMock := &GoCloakMock{}
		goCloakMock.On("GetUserInfo").Return(&gocloak.UserInfo{}, errors.New("error"))

		service := &Client{
			ctx:    context.Background(),
			client: goCloakMock,
		}

		_, err := service.GetUserInfo("access_token")
		assert.Error(t, err)
	})
}
