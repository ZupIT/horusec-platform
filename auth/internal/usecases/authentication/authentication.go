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
	"io"

	"github.com/google/uuid"

	authorization "github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	horusecAuthEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication/horusec"
)

type IUseCases interface {
	CheckLoginData(credentials *authEntities.LoginCredentials, account *accountEntities.Account) error
	LoginCredentialsFromIOReadCloser(body io.ReadCloser) (*authEntities.LoginCredentials, error)
	SetLdapAccountData(userData map[string]string) *accountEntities.Account
	FilterWorkspaceByID(workspaceID uuid.UUID) map[string]interface{}
	FilterRepositoryByID(repository uuid.UUID) map[string]interface{}
	FilterAccountWorkspaceByID(accountID, workspaceID uuid.UUID) map[string]interface{}
	FilterAccountRepositoryByID(accountID, repository uuid.UUID) map[string]interface{}
	NewAuthorizationDataFromGrpcData(data *proto.IsAuthorizedData) *authEntities.AuthorizationData
	NewIsAuthorizedResponse(isAuthorized bool) *proto.IsAuthorizedResponse
}

type UseCases struct {
}

func NewAuthenticationUseCases() IUseCases {
	return &UseCases{}
}

func (u *UseCases) CheckLoginData(credentials *authEntities.LoginCredentials, account *accountEntities.Account) error {
	if credentials.CheckInvalidPassword(account.Password) ||
		credentials.IsInvalidUsernameEmail() {
		return horusecAuthEnums.ErrorWrongEmailOrPassword
	}

	if account.IsNotConfirmed() {
		return horusecAuthEnums.ErrorAccountEmailNotConfirmed
	}

	return nil
}

func (u *UseCases) LoginCredentialsFromIOReadCloser(body io.ReadCloser) (*authEntities.LoginCredentials, error) {
	credentials := &authEntities.LoginCredentials{}

	if err := parser.ParseBodyToEntity(body, credentials); err != nil {
		return nil, err
	}

	return credentials, credentials.Validate()
}

func (u *UseCases) SetLdapAccountData(userData map[string]string) *accountEntities.Account {
	account := &accountEntities.Account{
		Username: userData["sAMAccountName"],
		Password: uuid.NewString(),
	}

	if userData["mail"] == "" {
		account.Email = userData["sAMAccountName"]
	} else {
		account.Email = userData["mail"]
	}

	return account.SetNewAccountData()
}

func (u *UseCases) FilterWorkspaceByID(workspaceID uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"workspace_id": workspaceID}
}

func (u *UseCases) FilterRepositoryByID(repository uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"repository_id": repository}
}

func (u *UseCases) FilterAccountWorkspaceByID(accountID, workspaceID uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"account_id": accountID, "workspace_id": workspaceID}
}

func (u *UseCases) FilterAccountRepositoryByID(accountID, repository uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"account_id": accountID, "repository_id": repository}
}

func (u *UseCases) NewAuthorizationDataFromGrpcData(data *proto.IsAuthorizedData) *authEntities.AuthorizationData {
	return &authEntities.AuthorizationData{
		Token:        data.Token,
		Type:         authorization.AuthorizationType(data.Type),
		WorkspaceID:  parser.ParseStringToUUID(data.WorkspaceID),
		RepositoryID: parser.ParseStringToUUID(data.RepositoryID),
	}
}

func (u *UseCases) NewIsAuthorizedResponse(isAuthorized bool) *proto.IsAuthorizedResponse {
	return &proto.IsAuthorizedResponse{
		IsAuthorized: isAuthorized,
	}
}
