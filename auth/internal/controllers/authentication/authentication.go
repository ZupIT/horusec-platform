package authentication

import (
	authTypes "github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	authEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/horusec"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/keycloak"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/ldap"
)

type IController interface {
	Login(credentials *authEntities.LoginCredentials) (*authEntities.LoginResponse, error)
	IsAuthorized(data *authEntities.AuthorizationData) (bool, error)
	GetAccountInfo(token string) (*proto.GetAccountDataResponse, error)
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

func (c *Controller) Login(credentials *authEntities.LoginCredentials) (*authEntities.LoginResponse, error) {
	switch c.appConfig.GetAuthenticationType() {
	case authTypes.Horusec:
		return c.horusecAuth.Login(credentials)
	case authTypes.Keycloak:
		return c.keycloakAuth.Login(credentials)
	case authTypes.Ldap:
		return c.ldapAuth.Login(credentials)
	}

	return nil, authEnums.ErrorAuthTypeInvalid
}

func (c *Controller) IsAuthorized(data *authEntities.AuthorizationData) (bool, error) {
	switch c.appConfig.GetAuthenticationType() {
	case authTypes.Horusec:
		return c.horusecAuth.IsAuthorized(data)
	case authTypes.Keycloak:
		return c.keycloakAuth.IsAuthorized(data)
	case authTypes.Ldap:
		return c.ldapAuth.IsAuthorized(data)
	}

	return false, authEnums.ErrorAuthTypeInvalid
}

func (c *Controller) GetAccountInfo(token string) (*proto.GetAccountDataResponse, error) {
	switch c.appConfig.GetAuthenticationType() {
	case authTypes.Horusec:
		return c.horusecAuth.GetAccountDataFromToken(token)
	case authTypes.Keycloak:
		return c.keycloakAuth.GetAccountDataFromToken(token)
	case authTypes.Ldap:
		return c.ldapAuth.GetAccountDataFromToken(token)
	}

	return nil, authEnums.ErrorAuthTypeInvalid
}
