package repository

import (
	"io"

	"github.com/google/uuid"

	emailEntities "github.com/ZupIT/horusec-devkit/pkg/entities/email"
	emailEnums "github.com/ZupIT/horusec-devkit/pkg/enums/email"
	databaseEnums "github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/core/internal/entities/repository"
)

type IUseCases interface {
	RepositoryDataFromIOReadCloser(body io.ReadCloser) (*repository.Data, error)
	FilterRepositoryByName(workspaceID uuid.UUID, name string) map[string]interface{}
	IsNotFoundError(err error) bool
	NewRepositoryData(repositoryID, workspaceID uuid.UUID, accountData *proto.GetAccountDataResponse) *repository.Data
	FilterRepositoryByID(repositoryID uuid.UUID) map[string]interface{}
	FilterAccountRepositoryByID(accountID, repositoryID uuid.UUID) map[string]interface{}
	NewRepositoryInviteEmail(email, username, repositoryName string) []byte
}

type UseCases struct {
}

func NewRepositoryUseCases() IUseCases {
	return &UseCases{}
}

func (u *UseCases) RepositoryDataFromIOReadCloser(body io.ReadCloser) (*repository.Data, error) {
	data := &repository.Data{}

	if err := parser.ParseBodyToEntity(body, data); err != nil {
		return nil, err
	}

	return data, data.Validate()
}

func (u *UseCases) FilterRepositoryByName(workspaceID uuid.UUID, name string) map[string]interface{} {
	return map[string]interface{}{"workspace_id": workspaceID, "name": name}
}

func (u *UseCases) IsNotFoundError(err error) bool {
	if err != nil {
		if err == databaseEnums.ErrorNotFoundRecords {
			return true
		}
	}

	return false
}

func (u *UseCases) NewRepositoryData(repositoryID, workspaceID uuid.UUID,
	accountData *proto.GetAccountDataResponse) *repository.Data {
	return &repository.Data{
		RepositoryID: repositoryID,
		WorkspaceID:  workspaceID,
		AccountID:    parser.ParseStringToUUID(accountData.AccountID),
		Permissions:  accountData.Permissions,
	}
}

func (u *UseCases) FilterRepositoryByID(repositoryID uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"repository_id": repositoryID}
}

func (u *UseCases) FilterAccountRepositoryByID(accountID, repositoryID uuid.UUID) map[string]interface{} {
	return map[string]interface{}{"account_id": accountID, "repository_id": repositoryID}
}

func (u *UseCases) NewRepositoryInviteEmail(email, username, repositoryName string) []byte {
	emailMessage := &emailEntities.Message{
		To:           email,
		TemplateName: emailEnums.RepositoryInvite,
		Subject:      "[Horusec] Repository invite",
		Data: map[string]interface{}{
			"repositoryName": repositoryName,
			"username":       username,
		},
	}

	return emailMessage.ToBytes()
}
