// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/ZupIT/horusec-devkit/pkg/utils/env"

	"github.com/ZupIT/horusec-platform/messages/internal/services/mailer/config/enums"
)

const (
	DefaultSMTPPort = 25
)

type IConfig interface {
	Validate() error
	SetUsername(username string)
	GetUsername() string
	SetPassword(password string)
	GetPassword() string
	SetHost(host string)
	GetHost() string
	SetPort(port int)
	GetPort() int
	SetFrom(from string)
	GetFrom() string
}

type Config struct {
	username string
	password string
	host     string
	port     int
	from     string
}

func NewMailerConfig() IConfig {
	config := &Config{}
	config.SetUsername(env.GetEnvOrDefault(enums.EnvSMTPUsername, ""))
	config.SetPassword(env.GetEnvOrDefault(enums.EnvSMTPPassword, ""))
	config.SetHost(env.GetEnvOrDefault(enums.EnvSMTPHost, ""))
	config.SetPort(env.GetEnvOrDefaultInt(enums.EnvSMTPPort, DefaultSMTPPort))
	config.SetFrom(env.GetEnvOrDefault(enums.EnvSMTPEmailFrom, "horusec@zup.com.br"))

	return config
}

func (c *Config) Validate() error {
	validations := []*validation.FieldRules{
		validation.Field(&c.host, validation.Required),
		validation.Field(&c.port, validation.Required),
		validation.Field(&c.username, validation.Required),
		validation.Field(&c.password, validation.Required),
		validation.Field(&c.from, validation.Required),
	}

	return validation.ValidateStruct(c, validations...)
}

func (c *Config) SetHost(host string) {
	c.host = host
}

func (c *Config) GetHost() string {
	return c.host
}

func (c *Config) SetUsername(username string) {
	c.username = username
}

func (c *Config) GetUsername() string {
	return c.username
}

func (c *Config) SetPassword(password string) {
	c.password = password
}

func (c *Config) GetPassword() string {
	return c.password
}

func (c *Config) SetPort(port int) {
	c.port = port
}

func (c *Config) GetPort() int {
	return c.port
}

func (c *Config) SetFrom(from string) {
	c.from = from
}

func (c *Config) GetFrom() string {
	return c.from
}
