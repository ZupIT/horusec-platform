package app

import (
	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/env"

	"github.com/ZupIT/horusec-platform/auth/config/app/enums"
)

type IConfig interface {
	GetAuthType() auth.AuthenticationType
	ToConfigResponse() map[string]interface{}
	IsApplicationAdminEnabled() bool
	IsDisableBroker() bool
	ToGetAuthConfigResponse() *proto.GetAuthConfigResponse
	GetHorusecAuthURL() string
	GetHorusecManagerURL() string
}

type Config struct {
	HorusecAuthURL         string
	AuthType               auth.AuthenticationType
	DisableBroker          bool
	EnableApplicationAdmin bool
	ApplicationAdminData   string
	EnableDefaultUser      bool
	DefaultUserData        string
	HorusecManagerURL      string
}

func NewAuthAppConfig() IConfig {
	return &Config{
		HorusecAuthURL:         env.GetEnvOrDefault(enums.EnvAuthURL, "http://localhost:8006"),
		AuthType:               auth.AuthenticationType(env.GetEnvOrDefault(enums.EnvAuthType, auth.Horusec.ToString())),
		DisableBroker:          env.GetEnvOrDefaultBool(enums.EnvDisableBroker, false),
		EnableApplicationAdmin: env.GetEnvOrDefaultBool(enums.EnvEnableApplicationAdmin, false),
		ApplicationAdminData:   env.GetEnvOrDefault(enums.EnvApplicationAdminData, enums.ApplicationAdminDefaultData),
		EnableDefaultUser:      env.GetEnvOrDefaultBool(enums.EnvEnableDefaultUser, true),
		DefaultUserData:        env.GetEnvOrDefault(enums.EnvDefaultUserData, enums.DefaultUserData),
		HorusecManagerURL:      env.GetEnvOrDefault(enums.EnvHorusecManager, "http://localhost:8043"),
	}
}

func (c *Config) GetAuthType() auth.AuthenticationType {
	return c.AuthType
}

func (c *Config) ToConfigResponse() map[string]interface{} {
	return map[string]interface{}{
		"enableApplicationAdmin": c.EnableApplicationAdmin,
		"authType":               c.AuthType,
		"disableBroker":          c.DisableBroker,
	}
}

func (c *Config) IsApplicationAdminEnabled() bool {
	return c.EnableApplicationAdmin
}

func (c *Config) IsDisableBroker() bool {
	return c.DisableBroker
}

func (c *Config) ToGetAuthConfigResponse() *proto.GetAuthConfigResponse {
	return &proto.GetAuthConfigResponse{
		EnableApplicationAdmin: c.EnableApplicationAdmin,
		AuthType:               c.AuthType.ToString(),
		DisableBroker:          c.DisableBroker,
	}
}

func (c *Config) GetHorusecAuthURL() string {
	return c.HorusecAuthURL
}

func (c *Config) GetHorusecManagerURL() string {
	return c.HorusecManagerURL
}
