package horusec

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	accountEnums "github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/cache"
	"github.com/ZupIT/horusec-devkit/pkg/utils/crypto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	horusecAuthEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication/horusec"
	accountRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/account"
	authRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/authentication"
	"github.com/ZupIT/horusec-platform/auth/internal/usecases/authentication"
)

func TestNewHorusecAuthenticationService(t *testing.T) {
	t.Run("should success create a new service", func(t *testing.T) {
		assert.NotNil(t, NewHorusecAuthenticationService(
			nil, nil, nil, nil, nil))
	})
}

func TestLogin(t *testing.T) {
	t.Run("should success login without errors", func(t *testing.T) {
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		passwordHash, _ := crypto.HashPasswordBcrypt("test")
		account := &accountEntities.Account{
			Password:           passwordHash,
			IsConfirmed:        true,
			Email:              "test@test.com",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByEmail").Return(account, nil)

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		credentials := &authEntities.LoginCredentials{
			Username: "test@test.com",
			Password: "test",
		}

		result, err := service.Login(credentials)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotEmpty(t, result.AccessToken)
		assert.NotEmpty(t, result.RefreshToken)
		assert.True(t, result.IsApplicationAdmin)
		assert.Equal(t, "test", result.Username)
		assert.Equal(t, "test@test.com", result.Email)
	})

	t.Run("should return error when account not confirmed", func(t *testing.T) {
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		passwordHash, _ := crypto.HashPasswordBcrypt("test")
		account := &accountEntities.Account{
			Password:           passwordHash,
			IsConfirmed:        false,
			Email:              "test@test.com",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByEmail").Return(account, nil)

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		credentials := &authEntities.LoginCredentials{
			Username: "test@test.com",
			Password: "test",
		}

		result, err := service.Login(credentials)
		assert.Error(t, err)
		assert.Equal(t, horusecAuthEnums.ErrorAccountEmailNotConfirmed, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when invalid email", func(t *testing.T) {
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		passwordHash, _ := crypto.HashPasswordBcrypt("test")
		account := &accountEntities.Account{
			Password:           passwordHash,
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByEmail").Return(account, nil)

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		credentials := &authEntities.LoginCredentials{
			Username: "test",
			Password: "test",
		}

		result, err := service.Login(credentials)
		assert.Error(t, err)
		assert.Equal(t, horusecAuthEnums.ErrorWrongEmailOrPassword, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when wrong password", func(t *testing.T) {
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		passwordHash, _ := crypto.HashPasswordBcrypt("test")
		account := &accountEntities.Account{
			Password:           passwordHash,
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByEmail").Return(account, nil)

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		credentials := &authEntities.LoginCredentials{
			Username: "test@test.com",
			Password: "test2",
		}

		result, err := service.Login(credentials)
		assert.Error(t, err)
		assert.Equal(t, horusecAuthEnums.ErrorWrongEmailOrPassword, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when failed to get account by email", func(t *testing.T) {
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		passwordHash, _ := crypto.HashPasswordBcrypt("test")
		account := &accountEntities.Account{
			Password:           passwordHash,
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByEmail").Return(account, errors.New("test"))

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		credentials := &authEntities.LoginCredentials{
			Username: "test@test.com",
			Password: "test2",
		}

		result, err := service.Login(credentials)
		assert.Error(t, err)
		assert.Equal(t, horusecAuthEnums.ErrorWrongEmailOrPassword, err)
		assert.Nil(t, result)
	})
}

func TestIsAuthorizedApplicationAdmin(t *testing.T) {
	t.Run("should return true and no error when application admin", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(account, nil)

		appConfig := &app.Config{EnableApplicationAdmin: true}

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.ApplicationAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("should return error when failed to get account", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(account, errors.New("test"))

		appConfig := &app.Config{EnableApplicationAdmin: true}

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.ApplicationAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should return error when application admin not enabled", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(account, errors.New("test"))

		appConfig := &app.Config{EnableApplicationAdmin: false}

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.ApplicationAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should return error when failed to get account id from token", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}
		accountRepositoryMock := &accountRepository.Mock{}

		appConfig := &app.Config{EnableApplicationAdmin: true}

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		data := &authEntities.AuthorizationData{
			Token:        "test",
			Type:         auth.ApplicationAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})
}

func TestIsAuthorizedWorkspaceAdmin(t *testing.T) {
	t.Run("should return true and no error when workspace admin", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: false,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Admin, nil)

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.WorkspaceAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("should return error when failed to get workspace role", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: false,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Admin, errors.New("test"))

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.WorkspaceAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should return error when failed to get account id from token", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		data := &authEntities.AuthorizationData{
			Token:        "test",
			Type:         auth.WorkspaceAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})
}

func TestIsAuthorizedWorkspaceMember(t *testing.T) {
	t.Run("should return true and no error when workspace member", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: false,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Member, nil)

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.WorkspaceMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("should return error when failed to get workspace role", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: false,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Member, errors.New("test"))

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.WorkspaceMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should return error when failed to get account id from token", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		data := &authEntities.AuthorizationData{
			Token:        "test",
			Type:         auth.WorkspaceMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})
}

func TestIsAuthorizedRepositoryMember(t *testing.T) {
	t.Run("should return true and no error when repository member", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: false,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetRepositoryRole").Return(accountEnums.Member, nil)
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Member, nil)

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.RepositoryMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("should return error when failed to get repository role", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: false,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetRepositoryRole").Return(accountEnums.Member, errors.New("test"))
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Member, errors.New("test"))

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.RepositoryMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should return error when failed to get account id from token", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		data := &authEntities.AuthorizationData{
			Token:        "test",
			Type:         auth.RepositoryMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})
}

func TestIsAuthorizedRepositorySupervisor(t *testing.T) {
	t.Run("should return true and no error when repository supervisor", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: false,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetRepositoryRole").Return(accountEnums.Supervisor, nil)
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Supervisor, nil)

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.RepositorySupervisor,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("should return error when failed to get repository role", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: false,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetRepositoryRole").Return(accountEnums.Supervisor, errors.New("test"))
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Supervisor, errors.New("test"))

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.RepositorySupervisor,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should return error when failed to get account id from token", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		data := &authEntities.AuthorizationData{
			Token:        "test",
			Type:         auth.RepositorySupervisor,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})
}

func TestIsAuthorizedRepositoryAdmin(t *testing.T) {
	t.Run("should return true and no error when repository admin", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: false,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetRepositoryRole").Return(accountEnums.Admin, nil)
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Admin, nil)

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.RepositoryAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("should return error when failed to get repository role", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: false,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetRepositoryRole").Return(accountEnums.Admin, errors.New("test"))
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Member, nil)

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.RepositoryAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should return error when failed to get account id from token", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		data := &authEntities.AuthorizationData{
			Token:        "test",
			Type:         auth.RepositoryAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})
}

func TestGetAccountDataFromToken(t *testing.T) {
	t.Run("should return account data without errors", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			AccountID:          uuid.New(),
			IsConfirmed:        true,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(account, nil)

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		result, err := service.GetAccountDataFromToken(token)
		assert.NoError(t, err)
		assert.Equal(t, account.AccountID.String(), result.AccountID)
		assert.Equal(t, account.IsApplicationAdmin, result.IsApplicationAdmin)
		assert.Equal(t, []string{"test"}, result.Permissions)
	})

	t.Run("should return error when failed to get account", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			AccountID:          uuid.New(),
			IsConfirmed:        true,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(account, errors.New("test"))

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		result, err := service.GetAccountDataFromToken(token)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when failed to get account", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			AccountID:          uuid.New(),
			IsConfirmed:        true,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(account, errors.New("test"))

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		result, err := service.GetAccountDataFromToken(token)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when failed to decode token", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}
		appConfig := &app.Config{}
		accountRepositoryMock := &accountRepository.Mock{}

		service := NewHorusecAuthenticationService(accountRepositoryMock, appConfig,
			authentication.NewAuthenticationUseCases(), authRepositoryMock, cache.NewCache())

		result, err := service.GetAccountDataFromToken("")
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestCheckRepositoryRequestForWorkspaceAdmin(t *testing.T) {
	t.Run("should return true for workspace admin", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			AccountID:          uuid.New(),
			IsConfirmed:        true,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Admin, nil)

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			cache:             cache.NewCache(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})
		data := &authEntities.AuthorizationData{
			Token: token,
		}

		result, err := service.checkRepositoryRequestForWorkspaceAdmin(data, enums.ErrorNotFoundRecords)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("should return false for workspace admin", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			AccountID:          uuid.New(),
			IsConfirmed:        true,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Member, nil)

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			cache:             cache.NewCache(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})
		data := &authEntities.AuthorizationData{
			Token: token,
		}

		result, err := service.checkRepositoryRequestForWorkspaceAdmin(data, enums.ErrorNotFoundRecords)
		assert.NoError(t, err)
		assert.False(t, result)
	})

	t.Run("should return false and error when failed to check for workspace admin", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			AccountID:          uuid.New(),
			IsConfirmed:        true,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Member, errors.New("test"))

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			cache:             cache.NewCache(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})
		data := &authEntities.AuthorizationData{
			Token: token,
		}

		result, err := service.checkRepositoryRequestForWorkspaceAdmin(data, enums.ErrorNotFoundRecords)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should return false and error when error different than not found records", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		authRepositoryMock := &authRepository.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			AccountID:          uuid.New(),
			IsConfirmed:        true,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			cache:             cache.NewCache(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})
		data := &authEntities.AuthorizationData{
			Token: token,
		}

		result, err := service.checkRepositoryRequestForWorkspaceAdmin(data, errors.New("test"))
		assert.Error(t, err)
		assert.False(t, result)
	})
}
