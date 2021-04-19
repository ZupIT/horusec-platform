package authentication

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"

	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
)

type IService interface {
	Login(credentials *authEntities.LoginCredentials) (*authEntities.LoginResponse, error)
	IsAuthorized(data *authEntities.AuthorizationData) (bool, error)
	GetAccountDataFromToken(token string) (*proto.GetAccountDataResponse, error)
}
