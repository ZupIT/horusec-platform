package authentication

import (
	authTypes "github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-platform/auth/config/app"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
)

type IController interface {
}

type Controller struct {
	appConfig app.IConfig
}

func NewAuthenticationController() IController {
	return &Controller{}
}

func (c *Controller) Login(credentials authEntities.LoginCredentials) (interface{}, error) {
	switch c.appConfig.GetAuthType() {
	case authTypes.Horusec:
		return c.horusAuthService.Authenticate(credentials)
		//case authTypes.Keycloak:
		//	return c.keycloakAuthService.Authenticate(credentials)
		//case authTypes.Ldap:
		//	return c.ldapAuthService.Authenticate(credentials)
	}

	return nil, errors.ErrorUnauthorized
}
