package authentication

import (
	authTypes "github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	authEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication"
)

type IController interface {
	Login(credentials *authEntities.LoginCredentials) (interface{}, error)
	IsAuthorized(data *authEntities.AuthorizationData) (bool, error)
	GetAccountInfo(token string) (*proto.GetAccountDataResponse, error)
}

type Controller struct {
	appConfig    app.IConfig
	horusecAuth  authentication.IService
	keycloakAuth authentication.IService
	ldapAuth     authentication.IService
}

func NewAuthenticationController(appConfig app.IConfig, authHorusec authentication.IService,
	ldapAuth authentication.IService, keycloakAuth authentication.IService) IController {
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

func (c *Controller) IsAuthorized(data *authEntities.AuthorizationData) (bool, error) {
	switch c.appConfig.GetAuthType() {
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
	switch c.appConfig.GetAuthType() {
	case authTypes.Horusec:
		return c.horusecAuth.GetAccountFromToken(token)
	case authTypes.Keycloak:
		return c.keycloakAuth.GetAccountFromToken(token)
	case authTypes.Ldap:
		return c.ldapAuth.GetAccountFromToken(token)
	}

	return nil, authEnums.ErrorAuthTypeInvalid
}
