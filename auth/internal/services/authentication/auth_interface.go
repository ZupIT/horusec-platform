package authentication

import authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"

type IService interface {
	Login(credentials *authEntities.LoginCredentials) (*authEntities.LoginResponse, error)
	IsAuthorized(data *authEntities.AuthorizationData) (bool, error)
}
