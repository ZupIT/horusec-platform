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

package client

import (
	"crypto/tls"
	"errors"
	"testing"

	"github.com/go-ldap/ldap/v3"
	"github.com/stretchr/testify/assert"

	ldapEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication/ldap"
)

func TestNewLDAPClient(t *testing.T) {
	t.Run("should create a new ldap client instance", func(t *testing.T) {
		ldapClient := NewLdapClient()
		assert.NotNil(t, ldapClient)
	})
}

func TestConnect(t *testing.T) {
	t.Run("should return error when connecting without ssl", func(t *testing.T) {
		service := &LdapClient{}

		assert.Error(t, service.Connect())
	})

	t.Run("should return error when connecting with ssl", func(t *testing.T) {
		service := &LdapClient{
			UseSSL:             true,
			ClientCertificates: []tls.Certificate{{}},
		}

		assert.Error(t, service.Connect())
	})
}

func TestAuthenticate(t *testing.T) {
	t.Run("should success authenticate", func(t *testing.T) {
		ldapMock := &MockConnection{}

		ldapMock.On("Bind").Return(nil)
		ldapMock.On("Search").Return(&ldap.SearchResult{Entries: []*ldap.Entry{{DN: "test",
			Attributes: []*ldap.EntryAttribute{{Name: "test", Values: []string{"test"}}}}}}, nil)

		service := &LdapClient{
			BindDN:       "test",
			BindPassword: "test",
			Conn:         ldapMock,
		}

		isValid, data, err := service.Authenticate("test", "test")
		assert.NoError(t, err)
		assert.True(t, isValid)
		assert.NotNil(t, data)
	})

	t.Run("should return error too many entries", func(t *testing.T) {
		ldapMock := &MockConnection{}

		ldapMock.On("Bind").Return(nil)
		ldapMock.On("Search").Return(&ldap.SearchResult{Entries: []*ldap.Entry{{}, {}}}, nil)

		service := &LdapClient{
			BindDN:       "test",
			BindPassword: "test",
			Conn:         ldapMock,
		}

		isValid, data, err := service.Authenticate("test", "test")
		assert.Error(t, err)
		assert.Equal(t, ldapEnums.ErrorTooManyEntries, err)
		assert.False(t, isValid)
		assert.Nil(t, data)
	})

	t.Run("should return error user does not exist", func(t *testing.T) {
		ldapMock := &MockConnection{}

		ldapMock.On("Bind").Return(nil)
		ldapMock.On("Search").Return(&ldap.SearchResult{}, nil)

		service := &LdapClient{
			BindDN:       "test",
			BindPassword: "test",
			Conn:         ldapMock,
		}

		isValid, data, err := service.Authenticate("test", "test")
		assert.Error(t, err)
		assert.Equal(t, ldapEnums.ErrorUserDoesNotExist, err)
		assert.False(t, isValid)
		assert.Nil(t, data)
	})

	t.Run("should return error when empty bind user or password", func(t *testing.T) {
		ldapMock := &MockConnection{}

		ldapMock.On("Bind").Return(nil)
		ldapMock.On("Search").Return(&ldap.SearchResult{}, nil)

		service := &LdapClient{
			Conn: ldapMock,
		}

		isValid, data, err := service.Authenticate("test", "test")
		assert.Error(t, err)
		assert.Equal(t, ldapEnums.ErrorEmptyBindDNOrBindPassword, err)
		assert.False(t, isValid)
		assert.Nil(t, data)
	})

	t.Run("should return error when while searching", func(t *testing.T) {
		ldapMock := &MockConnection{}

		ldapMock.On("Bind").Return(nil)
		ldapMock.On("Search").Return(&ldap.SearchResult{}, errors.New("test"))

		service := &LdapClient{
			BindDN:       "test",
			BindPassword: "test",
			Conn:         ldapMock,
		}

		isValid, data, err := service.Authenticate("test", "test")
		assert.Error(t, err)
		assert.False(t, isValid)
		assert.Nil(t, data)
	})

	t.Run("should return when binding with user data", func(t *testing.T) {
		ldapMock := &MockConnection{}

		ldapMock.On("Bind").Once().Return(nil)
		ldapMock.On("Bind").Return(errors.New("test"))
		ldapMock.On("Search").Return(&ldap.SearchResult{Entries: []*ldap.Entry{{DN: "test",
			Attributes: []*ldap.EntryAttribute{{Name: "test", Values: []string{"test"}}}}}}, nil)

		service := &LdapClient{
			BindDN:       "test",
			BindPassword: "test",
			Conn:         ldapMock,
		}

		isValid, data, err := service.Authenticate("test", "test")
		assert.Error(t, err)
		assert.False(t, isValid)
		assert.Nil(t, data)
	})

	t.Run("should return error when failed to connect", func(t *testing.T) {
		service := &LdapClient{}

		isValid, data, err := service.Authenticate("test", "test")
		assert.Error(t, err)
		assert.False(t, isValid)
		assert.Nil(t, data)
	})
}

func TestClose(t *testing.T) {
	t.Run("should success close connection", func(t *testing.T) {
		ldapMock := &MockConnection{}

		ldapMock.On("Close").Return(nil)

		service := &LdapClient{
			Conn: ldapMock,
		}

		assert.NotPanics(t, func() {
			service.Close()
		})
	})
}

func TestGetGroupsOfUser(t *testing.T) {
	t.Run("should success get groups", func(t *testing.T) {
		ldapMock := &MockConnection{}

		ldapMock.On("Bind").Return(nil)
		ldapMock.On("Search").Return(&ldap.SearchResult{Entries: []*ldap.Entry{{DN: "test",
			Attributes: []*ldap.EntryAttribute{{Name: "cn", Values: []string{"test"}}}}}}, nil)

		service := &LdapClient{
			BindDN:       "test",
			BindPassword: "test",
			Conn:         ldapMock,
		}

		groups, err := service.GetUserGroups("test")
		assert.NotEmpty(t, groups)
		assert.NoError(t, err)
	})

	t.Run("should return error while searching groups", func(t *testing.T) {
		ldapMock := &MockConnection{}

		ldapMock.On("Bind").Return(nil)
		ldapMock.On("Search").Return(&ldap.SearchResult{}, errors.New("test"))

		service := &LdapClient{
			BindDN:       "test",
			BindPassword: "test",
			Conn:         ldapMock,
		}

		groups, err := service.GetUserGroups("test")
		assert.Nil(t, groups)
		assert.Error(t, err)
	})

	t.Run("should return error while biding with env vars", func(t *testing.T) {
		ldapMock := &MockConnection{}

		ldapMock.On("Bind").Return(errors.New("test"))

		service := &LdapClient{
			BindDN:       "test",
			BindPassword: "test",
			Conn:         ldapMock,
		}

		groups, err := service.GetUserGroups("test")
		assert.Nil(t, groups)
		assert.Error(t, err)
	})

	t.Run("should return error while connecting", func(t *testing.T) {
		service := &LdapClient{}

		groups, err := service.GetUserGroups("test")
		assert.Nil(t, groups)
		assert.Error(t, err)
	})
}

func TestCheck(t *testing.T) {
	t.Run("should return no true when ldap is healthy", func(t *testing.T) {
		ldapMock := &MockConnection{}

		ldapMock.On("Search").Return(&ldap.SearchResult{}, nil)

		service := &LdapClient{
			Conn: ldapMock,
		}

		assert.True(t, service.IsAvailable())
	})

	t.Run("should return false when ldap is not healthy", func(t *testing.T) {
		ldapMock := &MockConnection{}

		ldapMock.On("Search").Return(&ldap.SearchResult{}, errors.New("test"))

		service := &LdapClient{
			Conn: ldapMock,
		}

		assert.False(t, service.IsAvailable())
	})

	t.Run("should return false when connecting to ldap return error", func(t *testing.T) {
		service := &LdapClient{}

		assert.False(t, service.IsAvailable())
	})
}

func TestSetLDAPServiceConnection(t *testing.T) {
	t.Run("should success set connection", func(t *testing.T) {
		service := &LdapClient{}

		assert.NoError(t, service.setLDAPServiceConnection(&ldap.Conn{}, nil))
		assert.NotNil(t, service.Conn)
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("should success create user with case sensitive", func(t *testing.T) {
		service := &LdapClient{}

		searchResult := &ldap.SearchResult{
			Entries: []*ldap.Entry{
				{
					DN: "test",
					Attributes: []*ldap.EntryAttribute{
						{
							Name:   "sAMAccountName",
							Values: []string{"test"},
						},
						{
							Name:   "mail",
							Values: []string{"test"},
						},
					},
				},
			},
		}

		result := service.createUser(searchResult)
		assert.Equal(t, "test", result["sAMAccountName"])
		assert.Equal(t, "test", result["mail"])
	})

	t.Run("should success create user with lower case", func(t *testing.T) {
		service := &LdapClient{}

		searchResult := &ldap.SearchResult{
			Entries: []*ldap.Entry{
				{
					DN: "test",
					Attributes: []*ldap.EntryAttribute{
						{
							Name:   "samaccountname",
							Values: []string{"test"},
						},
						{
							Name:   "mail",
							Values: []string{"test"},
						},
					},
				},
			},
		}

		result := service.createUser(searchResult)
		assert.Equal(t, "test", result["sAMAccountName"])
		assert.Equal(t, "test", result["mail"])
	})
}
