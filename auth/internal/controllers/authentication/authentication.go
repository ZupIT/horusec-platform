package authentication

import (
	authTypes "github.com/ZupIT/horusec-devkit/pkg/enums/auth"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	authEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/horusec"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/keycloak"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/ldap"
)

type IController interface {
	Login(credentials *authEntities.LoginCredentials) (interface{}, error)
}

type Controller struct {
	appConfig    app.IConfig
	horusecAuth  horusec.IService
	keycloakAuth keycloak.IService
	ldapAuth     ldap.IService
}

func NewAuthenticationController(appConfig app.IConfig, authHorusec horusec.IService,
	ldapAuth ldap.IService, keycloakAuth keycloak.IService) IController {
	return &Controller{
		appConfig:    appConfig,
		horusecAuth:  authHorusec,
		ldapAuth:     ldapAuth,
		keycloakAuth: keycloakAuth,
	}
}

func (c *Controller) Login(credentials *authEntities.LoginCredentials) (interface{}, error) {
	switch c.appConfig.GetAuthType() {
	case authTypes.Horusec:
		return c.horusecAuth.Login(credentials)
	case authTypes.Keycloak:
		return c.keycloakAuth.Login(credentials)
	case authTypes.Ldap:
		return c.ldapAuth.Login(credentials)
	}

	return nil, authEnums.ErrorAuthTypeInvalid
}
