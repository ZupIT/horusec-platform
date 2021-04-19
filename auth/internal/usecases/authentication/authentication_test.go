package authentication

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"

	authorization "github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"
	"github.com/ZupIT/horusec-devkit/pkg/utils/crypto"

	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	horusecAuthEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication/horusec"
)

func TestNewAuthenticationUseCases(t *testing.T) {
	t.Run("should success create a new use cases", func(t *testing.T) {
		assert.NotNil(t, NewAuthenticationUseCases())
	})
}

func TestCheckLoginData(t *testing.T) {
	t.Run("should no error when everything is ok", func(t *testing.T) {
		useCases := NewAuthenticationUseCases()

		credentials := &authEntities.LoginCredentials{
			Username: "test@test.com",
			Password: "test",
		}

		hash, _ := crypto.HashPasswordBcrypt("test")
		account := &accountEntities.Account{
			AccountID:          uuid.New(),
			Email:              "test@test.com",
			Password:           hash,
			Username:           "test",
			IsConfirmed:        true,
			IsApplicationAdmin: true,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}

		assert.NoError(t, useCases.CheckLoginData(credentials, account))
	})

	t.Run("should no error when account not confirmed", func(t *testing.T) {
		useCases := NewAuthenticationUseCases()

		credentials := &authEntities.LoginCredentials{
			Username: "test@test.com",
			Password: "test",
		}

		hash, _ := crypto.HashPasswordBcrypt("test")
		account := &accountEntities.Account{
			AccountID:          uuid.New(),
			Email:              "test@test.com",
			Password:           hash,
			Username:           "test",
			IsConfirmed:        false,
			IsApplicationAdmin: true,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}

		err := useCases.CheckLoginData(credentials, account)
		assert.Error(t, err)
		assert.Equal(t, horusecAuthEnums.ErrorAccountEmailNotConfirmed, err)
	})

	t.Run("should no error when invalid email", func(t *testing.T) {
		useCases := NewAuthenticationUseCases()

		credentials := &authEntities.LoginCredentials{
			Username: "test",
			Password: "test",
		}

		hash, _ := crypto.HashPasswordBcrypt("test")
		account := &accountEntities.Account{
			AccountID:          uuid.New(),
			Email:              "test@test.com",
			Password:           hash,
			Username:           "test",
			IsConfirmed:        false,
			IsApplicationAdmin: true,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}

		err := useCases.CheckLoginData(credentials, account)
		assert.Error(t, err)
		assert.Equal(t, horusecAuthEnums.ErrorWrongEmailOrPassword, err)
	})

	t.Run("should no error when wrong password", func(t *testing.T) {
		useCases := NewAuthenticationUseCases()

		credentials := &authEntities.LoginCredentials{
			Username: "test@test.com",
			Password: "test",
		}

		hash, _ := crypto.HashPasswordBcrypt("123")
		account := &accountEntities.Account{
			AccountID:          uuid.New(),
			Email:              "test@test.com",
			Password:           hash,
			Username:           "test",
			IsConfirmed:        false,
			IsApplicationAdmin: true,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}

		err := useCases.CheckLoginData(credentials, account)
		assert.Error(t, err)
		assert.Equal(t, horusecAuthEnums.ErrorWrongEmailOrPassword, err)
	})
}

func TestLoginCredentialsFromIOReadCloser(t *testing.T) {
	t.Run("should success get login data from request body", func(t *testing.T) {
		useCases := NewAuthenticationUseCases()

		data := &authEntities.LoginCredentials{
			Username: "test",
			Password: "test",
		}

		readCloser, err := parser.ParseEntityToIOReadCloser(data)
		assert.NoError(t, err)

		response, err := useCases.LoginCredentialsFromIOReadCloser(readCloser)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, data.Username, response.Username)
		assert.Equal(t, data.Password, response.Password)
	})

	t.Run("should return error when failed to parse body to entity", func(t *testing.T) {
		useCases := NewAuthenticationUseCases()

		readCloser, err := parser.ParseEntityToIOReadCloser("")
		assert.NoError(t, err)

		response, err := useCases.LoginCredentialsFromIOReadCloser(readCloser)
		assert.Error(t, err)
		assert.Nil(t, response)
	})
}

func TestSetLdapAccountData(t *testing.T) {
	t.Run("should success set ldap account data with mail", func(t *testing.T) {
		useCases := NewAuthenticationUseCases()

		userData := map[string]string{
			"mail":           "test@test.com",
			"sAMAccountName": "test",
		}

		account := useCases.SetLdapAccountData(userData)
		assert.NotNil(t, account)
		assert.Equal(t, "test", account.Username)
		assert.Equal(t, "test@test.com", account.Email)
		assert.NotEmpty(t, account.Password)
	})

	t.Run("should success set ldap account data with account name", func(t *testing.T) {
		useCases := NewAuthenticationUseCases()

		userData := map[string]string{
			"mail":           "",
			"sAMAccountName": "test",
		}

		account := useCases.SetLdapAccountData(userData)
		assert.NotNil(t, account)
		assert.Equal(t, "test", account.Username)
		assert.Equal(t, "test", account.Email)
		assert.NotEmpty(t, account.Password)
	})
}

func TestFilterWorkspaceByID(t *testing.T) {
	t.Run("should success create a filter by workspace id", func(t *testing.T) {
		useCases := NewAuthenticationUseCases()

		id := uuid.New()

		filter := useCases.FilterWorkspaceByID(id)
		assert.NotPanics(t, func() {
			assert.Equal(t, id, filter["workspace_id"])
		})
	})
}

func TestFilterRepositoryByID(t *testing.T) {
	t.Run("should success create a filter by repository id", func(t *testing.T) {
		useCases := NewAuthenticationUseCases()

		id := uuid.New()

		filter := useCases.FilterRepositoryByID(id)
		assert.NotPanics(t, func() {
			assert.Equal(t, id, filter["repository_id"])
		})
	})
}

func TestFilterAccountWorkspaceByID(t *testing.T) {
	t.Run("should success create a filter by account workspace id", func(t *testing.T) {
		useCases := NewAuthenticationUseCases()

		id := uuid.New()

		filter := useCases.FilterAccountWorkspaceByID(id, id)
		assert.NotPanics(t, func() {
			assert.Equal(t, id, filter["account_id"])
			assert.Equal(t, id, filter["workspace_id"])
		})
	})
}

func TestFilterAccountRepositoryByID(t *testing.T) {
	t.Run("should success create a filter by account repository id", func(t *testing.T) {
		useCases := NewAuthenticationUseCases()

		id := uuid.New()

		filter := useCases.FilterAccountRepositoryByID(id, id)
		assert.NotPanics(t, func() {
			assert.Equal(t, id, filter["account_id"])
			assert.Equal(t, id, filter["repository_id"])
		})
	})
}

func TestNewAuthorizationDataFromGrpcData(t *testing.T) {
	t.Run("should success create a filter by account repository id", func(t *testing.T) {
		useCases := NewAuthenticationUseCases()

		isAuthorizedData := &proto.IsAuthorizedData{
			Token:        "test",
			Type:         "admin",
			WorkspaceID:  uuid.NewString(),
			RepositoryID: uuid.NewString(),
		}

		data := useCases.NewAuthorizationDataFromGrpcData(isAuthorizedData)
		assert.Equal(t, isAuthorizedData.Token, data.Token)
		assert.Equal(t, authorization.AuthorizationType(isAuthorizedData.Type), data.Type)
		assert.Equal(t, parser.ParseStringToUUID(isAuthorizedData.WorkspaceID), data.WorkspaceID)
		assert.Equal(t, parser.ParseStringToUUID(isAuthorizedData.RepositoryID), data.RepositoryID)
	})
}

func TestNewIsAuthorizedResponse(t *testing.T) {
	t.Run("should success create a new is authorized response", func(t *testing.T) {
		useCases := NewAuthenticationUseCases()

		response := useCases.NewIsAuthorizedResponse(true)
		assert.True(t, response.IsAuthorized)
	})
}
