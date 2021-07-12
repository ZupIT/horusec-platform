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

package ldap

const (
	EnvLdapHost               = "HORUSEC_LDAP_HOST"
	EnvLdapPort               = "HORUSEC_LDAP_PORT"
	EnvLdapBase               = "HORUSEC_LDAP_BASE"
	EnvLdapBindDn             = "HORUSEC_LDAP_BINDDN"
	EnvLdapBindPassword       = "HORUSEC_LDAP_BINDPASSWORD" //nolint:gosec // false positive
	EnvLdapUseSSL             = "HORUSEC_LDAP_USESSL"
	EnvLdapSkipTLS            = "HORUSEC_LDAP_SKIP_TLS"
	EnvLdapInsecureSkipVerify = "HORUSEC_LDAP_INSECURE_SKIP_VERIFY"
	EnvLdapUserFilter         = "HORUSEC_LDAP_USERFILTER"
	EnvLdapAdminGroup         = "HORUSEC_LDAP_ADMIN_GROUP"
	DefaultLdapUserFilter     = "(sAMAccountName=%s)"
)
