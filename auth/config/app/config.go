// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

package app

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseEnums "github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/env"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"

	"github.com/ZupIT/horusec-platform/auth/config/app/enums"
	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	accountEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/account"
)

type AdminAccount interface {
	CreateOrUpdate(account *accountEntities.Account) error
}

type IConfig interface {
	GetAuthenticationType() auth.AuthenticationType
	ToConfigResponse() map[string]interface{}
	IsApplicationAdmEnabled() bool
	IsEmailsDisabled() bool
	ToGetAuthConfigResponse() *proto.GetAuthConfigResponse
	GetHorusecAuthURL() string
	GetHorusecManagerURL() string
	GetEnableApplicationAdmin() bool
	GetEnableDefaultUser() bool
	GetDefaultUserData() (*accountEntities.Account, error)
	GetApplicationAdminData() (*accountEntities.Account, error)
}

type Config struct {
	HorusecAuthURL         string
	AuthType               auth.AuthenticationType
	DisableEmails          bool
	EnableApplicationAdmin bool
	ApplicationAdminData   string
	EnableDefaultUser      bool
	DefaultUserData        string
	HorusecManagerURL      string
	databaseWrite          database.IDatabaseWrite
	databaseRead           database.IDatabaseRead
}

func NewAuthAppConfig(connection *database.Connection) IConfig {
	config := &Config{
		HorusecAuthURL:         env.GetEnvOrDefault(enums.EnvAuthURL, enums.HorusecAuthLocalhost),
		AuthType:               auth.AuthenticationType(env.GetEnvOrDefault(enums.EnvAuthType, auth.Horusec.ToString())),
		DisableEmails:          env.GetEnvOrDefaultBool(enums.EnvDisableEmails, false),
		EnableApplicationAdmin: env.GetEnvOrDefaultBool(enums.EnvEnableApplicationAdmin, false),
		ApplicationAdminData:   env.GetEnvOrDefault(enums.EnvApplicationAdminData, enums.ApplicationAdminDefaultData),
		EnableDefaultUser:      env.GetEnvOrDefaultBool(enums.EnvEnableDefaultUser, false),
		DefaultUserData:        env.GetEnvOrDefault(enums.EnvDefaultUserData, enums.DefaultUserData),
		HorusecManagerURL:      env.GetEnvOrDefault(enums.EnvHorusecManager, enums.HorusecManagerLocalhost),
		databaseWrite:          connection.Write,
		databaseRead:           connection.Read,
	}

	return config.createDefaultUsers()
}

func (c *Config) GetAuthenticationType() auth.AuthenticationType {
	return c.AuthType
}

func (c *Config) ToConfigResponse() map[string]interface{} {
	return map[string]interface{}{
		"enableApplicationAdmin": c.EnableApplicationAdmin,
		"authType":               c.AuthType,
		"disableEmails":          c.DisableEmails,
	}
}

func (c *Config) IsApplicationAdmEnabled() bool {
	return c.EnableApplicationAdmin
}

func (c *Config) IsEmailsDisabled() bool {
	return c.DisableEmails
}

func (c *Config) ToGetAuthConfigResponse() *proto.GetAuthConfigResponse {
	return &proto.GetAuthConfigResponse{
		EnableApplicationAdmin: c.EnableApplicationAdmin,
		AuthType:               c.AuthType.ToString(),
		DisableEmails:          c.DisableEmails,
	}
}

func (c *Config) GetHorusecAuthURL() string {
	return c.HorusecAuthURL
}

func (c *Config) GetHorusecManagerURL() string {
	return c.HorusecManagerURL
}

func (c *Config) GetEnableApplicationAdmin() bool {
	return c.EnableApplicationAdmin
}

func (c *Config) GetEnableDefaultUser() bool {
	return c.EnableDefaultUser
}

func (c *Config) GetDefaultUserData() (*accountEntities.Account, error) {
	account := &accountEntities.Account{}

	return account, json.Unmarshal([]byte(c.DefaultUserData), &account)
}

func (c *Config) GetApplicationAdminData() (*accountEntities.Account, error) {
	account := &accountEntities.Account{}

	logger.LogWarn(fmt.Sprintf(enums.MessageFailedToFormatAppAdminValue, enums.EnvApplicationAdminData))
	return account.SetApplicationAdminTrue(), json.Unmarshal([]byte(c.ApplicationAdminData), account)
}

func (c *Config) createDefaultUsers() IConfig {
	if c.GetEnableDefaultUser() {
		c.createHorusecDefaultUser()
	}

	if c.GetEnableApplicationAdmin() {
		c.createApplicationAdminUser()
	}

	return c
}

func (c *Config) getDefaultUserData() (*accountEntities.Account, error) {
	account, err := c.GetDefaultUserData()
	if err != nil {
		return nil, err
	}

	return account.SetNewAccountData().SetIsConfirmedTrue(), nil
}

func (c *Config) createHorusecDefaultUser() {
	if c.GetAuthenticationType() != auth.Horusec {
		logger.LogWarn(enums.MessageDefaultUserAuthType)
		return
	}

	account, err := c.getDefaultUserData()
	if err != nil {
		logger.LogPanic(enums.MessageFailedToGetDefaultUserData, err)
		return
	}

	c.createAccount(account)
}

func (c *Config) getApplicationAdminData() (*accountEntities.Account, error) {
	account, err := c.GetApplicationAdminData()
	if err != nil {
		return nil, err
	}

	return account.SetNewAccountData().SetIsConfirmedTrue(), nil
}

func (c *Config) createApplicationAdminUser() {
	account, err := c.getApplicationAdminData()
	if err != nil {
		logger.LogPanic(enums.MessageFailedToGetApplicationAdminData, err)
		return
	}

	if err = c.createOrUpdate(account); err != nil {
		logger.LogPanic(enums.MessageFailedToCreateAccount, err)
	}

	logger.LogInfo(fmt.Sprintf(enums.MessageUserCreateWithSuccess, account.Username, account.Email))
}

func (c *Config) createAccount(account *accountEntities.Account) {
	err := c.databaseWrite.Create(account, accountEnums.DatabaseTableAccount).GetError()
	if err != nil {
		c.checkCreateAccountErrors(err, account)
		return
	}

	logger.LogInfo(fmt.Sprintf(enums.MessageUserCreateWithSuccess, account.Username, account.Email))
}

func (c *Config) checkCreateAccountErrors(err error, account *accountEntities.Account) {
	if strings.Contains(strings.ToLower(err.Error()), enums.DuplicatedAccount) {
		logger.LogInfo(fmt.Sprintf(enums.MessageUserAlreadyExists, account.Username, account.Email))
		return
	}

	logger.LogPanic(enums.MessageFailedToCreateAccount, err)
}

func (c *Config) createOrUpdate(newest *accountEntities.Account) error {
	oldest, err := c.getAccountByEmail(newest.Email)
	if err != nil {
		if err == databaseEnums.ErrorNotFoundRecords {
			return c.deleteAllAndCreateApplicationAdmin(newest)
		}

		return err
	}

	return c.deleteAllAndUpdateApplicationAdmin(oldest, newest)
}

func (c *Config) deleteAllAndUpdateApplicationAdmin(oldest, newest *accountEntities.Account) error {
	transaction := c.databaseWrite.StartTransaction()

	if err := transaction.Delete(c.filterApplicationAdminTrue(),
		accountEnums.DatabaseTableAccount).GetError(); err != nil {
		return errors.Wrap(err, c.checkForNilError(transaction.RollbackTransaction().GetError()))
	}

	if err := transaction.Create(c.mergeApplicationAdminAccounts(oldest, newest),
		accountEnums.DatabaseTableAccount).GetError(); err != nil {
		return errors.Wrap(err, c.checkForNilError(transaction.RollbackTransaction().GetError()))
	}

	return transaction.CommitTransaction().GetError()
}

func (c *Config) mergeApplicationAdminAccounts(oldest, newest *accountEntities.Account) *accountEntities.Account {
	return &accountEntities.Account{
		AccountID:          oldest.AccountID,
		Email:              oldest.Email,
		Password:           newest.Password,
		Username:           newest.Username,
		IsConfirmed:        newest.IsConfirmed,
		IsApplicationAdmin: true,
		CreatedAt:          oldest.CreatedAt,
		UpdatedAt:          newest.UpdatedAt,
	}
}

func (c *Config) deleteAllAndCreateApplicationAdmin(newest *accountEntities.Account) error {
	transaction := c.databaseWrite.StartTransaction()

	if err := transaction.Delete(c.filterApplicationAdminTrue(),
		accountEnums.DatabaseTableAccount).GetError(); err != nil {
		return errors.Wrap(transaction.RollbackTransaction().GetError(), err.Error())
	}

	if err := transaction.Create(newest, accountEnums.DatabaseTableAccount).GetError(); err != nil {
		return errors.Wrap(transaction.RollbackTransaction().GetError(), err.Error())
	}

	return transaction.CommitTransaction().GetError()
}

func (c *Config) filterApplicationAdminTrue() map[string]interface{} {
	return map[string]interface{}{"is_application_admin": true}
}

func (c *Config) getAccountByEmail(email string) (*accountEntities.Account, error) {
	account := &accountEntities.Account{}

	return account, c.databaseRead.Find(account, c.filterAccountByEmail(email),
		accountEnums.DatabaseTableAccount).GetError()
}

func (c *Config) filterAccountByEmail(email string) map[string]interface{} {
	return map[string]interface{}{"email": email}
}

func (c *Config) checkForNilError(err error) string {
	if err != nil {
		return err.Error()
	}

	return ""
}
