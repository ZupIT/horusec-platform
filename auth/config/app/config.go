package app

import (
	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/utils/env"

	"github.com/ZupIT/horusec-platform/auth/config/app/enums"
)

type IConfig interface {
	GetAuthType() auth.AuthorizationType
}

type Config struct {
	HorusecAuthURL         string
	AuthType               auth.AuthorizationType
	DisableBroker          bool
	EnableApplicationAdmin bool
	ApplicationAdminData   string
	EnableDefaultUser      bool
	DefaultUserData        string
}

func NewAuthAppConfig() IConfig {
	return &Config{
		HorusecAuthURL:         env.GetEnvOrDefault(enums.EnvAuthURL, "http://localhost:8006"),
		AuthType:               auth.AuthorizationType(env.GetEnvOrDefault(enums.EnvAuthType, auth.Horusec.ToString())),
		DisableBroker:          env.GetEnvOrDefaultBool(enums.EnvDisableBroker, false),
		EnableApplicationAdmin: env.GetEnvOrDefaultBool(enums.EnvEnableApplicationAdmin, false),
		ApplicationAdminData:   env.GetEnvOrDefault(enums.EnvApplicationAdminData, enums.ApplicationAdminDefaultData),
		EnableDefaultUser:      env.GetEnvOrDefaultBool(enums.EnvEnableDefaultUser, true),
		DefaultUserData:        env.GetEnvOrDefault(enums.EnvDefaultUserData, enums.DefaultUserData),
	}
}

func (c *Config) GetAuthType() auth.AuthorizationType {
	return c.AuthType
}
