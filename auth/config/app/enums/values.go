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

package enums

const (
	EnvAuthURL                  = "HORUSEC_AUTH_URL"
	EnvAuthType                 = "HORUSEC_AUTH_TYPE"
	EnvDisableEmails            = "HORUSEC_DISABLE_EMAILS"
	EnvEnableApplicationAdmin   = "HORUSEC_ENABLE_APPLICATION_ADMIN"
	EnvApplicationAdminData     = "HORUSEC_APPLICATION_ADMIN_DATA"
	EnvEnableDefaultUser        = "HORUSEC_ENABLE_DEFAULT_USER"
	EnvDefaultUserData          = "HORUSEC_DEFAULT_USER_DATA"
	EnvHorusecManager           = "HORUSEC_MANAGER_URL"
	DuplicatedAccount           = "duplicate key value violates unique constraint"
	DefaultUserData             = "{\"username\": \"dev\", \"email\":\"dev@example.com\", \"password\":\"Devpass0*\"}"
	ApplicationAdminDefaultData = "{\"username\": \"horusec-admin\", \"email\":\"horusec-admin@example.com\"," +
		" \"password\":\"Devpass0*\"}"
)
