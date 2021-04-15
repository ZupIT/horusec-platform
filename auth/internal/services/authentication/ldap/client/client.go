package client

import (
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/go-ldap/ldap/v3"

	"github.com/ZupIT/horusec-devkit/pkg/utils/env"

	ldapEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication/ldap"
)

type ILdapClient interface {
	Connect() error
	Close()
	Authenticate(username, password string) (bool, map[string]string, error)
	GetUserGroups(userDN string) ([]string, error)
	IsAvailable() bool
}

type LdapClient struct {
	Host               string
	Port               int
	Base               string
	BindDN             string
	BindPassword       string
	ServerName         string
	InsecureSkipVerify bool
	UseSSL             bool
	SkipTLS            bool
	ClientCertificates []tls.Certificate
	Conn               IConnection
	UserFilter         string
}

func NewLdapClient() ILdapClient {
	return &LdapClient{
		Host:               env.GetEnvOrDefault(ldapEnums.EnvLdapHost, ""),
		Port:               env.GetEnvOrDefaultInt(ldapEnums.EnvLdapPort, 389),
		Base:               env.GetEnvOrDefault(ldapEnums.EnvLdapBase, ""),
		BindDN:             env.GetEnvOrDefault(ldapEnums.EnvLdapBindDn, ""),
		BindPassword:       env.GetEnvOrDefault(ldapEnums.EnvLdapBindPassword, ""),
		UseSSL:             env.GetEnvOrDefaultBool(ldapEnums.EnvLdapUseSSL, false),
		SkipTLS:            env.GetEnvOrDefaultBool(ldapEnums.EnvLdapSkipTLS, true),
		InsecureSkipVerify: env.GetEnvOrDefaultBool(ldapEnums.EnvLdapInsecureSkipVerify, true),
		UserFilter:         env.GetEnvOrDefault(ldapEnums.EnvLdapUserFilter, ldapEnums.DefaultLdapUserFilter),
	}
}

func (l *LdapClient) Connect() error {
	if l.Conn != nil {
		return nil
	}

	if l.UseSSL {
		return l.dialWithSSL()
	}

	return l.dialWithoutSSL()
}

//nolint:gosec // false optional dial type
func (l *LdapClient) dialWithoutSSL() error {
	conn, err := ldap.Dial("tcp", l.getLdapURL())
	if err != nil {
		return err
	}

	if !l.SkipTLS {
		err = conn.StartTLS(&tls.Config{InsecureSkipVerify: true})
	}

	return l.setLDAPServiceConnection(conn, err)
}

//nolint:gosec // false optional dial type
func (l *LdapClient) dialWithSSL() error {
	config := &tls.Config{
		InsecureSkipVerify: l.InsecureSkipVerify,
		ServerName:         l.ServerName,
	}

	if l.ClientCertificates != nil && len(l.ClientCertificates) > 0 {
		config.Certificates = l.ClientCertificates
	}

	return l.setLDAPServiceConnection(ldap.DialTLS("tcp", l.getLdapURL(), config))
}

func (l *LdapClient) getLdapURL() string {
	return fmt.Sprintf("%s:%d", l.Host, l.Port)
}

func (l *LdapClient) setLDAPServiceConnection(conn *ldap.Conn, err error) error {
	if err != nil {
		return err
	}

	l.Conn = conn
	return nil
}

func (l *LdapClient) Close() {
	if l.Conn != nil {
		l.Conn.Close()
		l.Conn = nil
	}
}

func (l *LdapClient) connectAndBind() error {
	if err := l.Connect(); err != nil {
		return err
	}

	return l.bindByEnvVars()
}

func (l *LdapClient) Authenticate(username, password string) (bool, map[string]string, error) {
	if err := l.connectAndBind(); err != nil {
		return false, nil, err
	}

	return l.searchAndCreateUser(username, password)
}

func (l *LdapClient) searchAndCreateUser(username, password string) (bool, map[string]string, error) {
	searchResult, err := l.searchUserByUsername(username)
	if err != nil {
		return false, nil, err
	}

	if err := l.Conn.Bind(l.getDNBySearchResult(searchResult), password); err != nil {
		return false, nil, err
	}

	return true, l.createUser(searchResult), nil
}

func (l *LdapClient) getDNBySearchResult(searchResult *ldap.SearchResult) string {
	return searchResult.Entries[0].DN
}

func (l *LdapClient) bindByEnvVars() error {
	if l.BindDN != "" && l.BindPassword != "" {
		return l.Conn.Bind(l.BindDN, l.BindPassword)
	}

	return ldapEnums.ErrorEmptyBindDNOrBindPassword
}

func (l *LdapClient) searchUserByUsername(username string) (*ldap.SearchResult, error) {
	searchResult, err := l.Conn.Search(l.newSearchRequestByUserFilter(username))
	if err != nil {
		return nil, err
	}

	return searchResult, l.validateSearchResult(searchResult)
}

func (l *LdapClient) newSearchRequestByUserFilter(username string) *ldap.SearchRequest {
	return ldap.NewSearchRequest(
		l.Base,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(l.UserFilter, username),
		[]string{"sAMAccountName", "mail"},
		nil,
	)
}

func (l *LdapClient) validateSearchResult(searchResult *ldap.SearchResult) error {
	if searchResult == nil || searchResult.Entries == nil || len(searchResult.Entries) < 1 {
		return ldapEnums.ErrorUserDoesNotExist
	}

	if len(searchResult.Entries) > 1 {
		return ldapEnums.ErrorTooManyEntries
	}

	return nil
}

func (l *LdapClient) createUser(searchResult *ldap.SearchResult) map[string]string {
	user := map[string]string{"dn": l.getDNBySearchResult(searchResult)}

	for _, attr := range []string{"sAMAccountName", "mail"} {
		if value := searchResult.Entries[0].GetAttributeValue(attr); value != "" {
			user[attr] = value
		} else {
			user[attr] = searchResult.Entries[0].GetAttributeValue(strings.ToLower(attr))
		}
	}

	return user
}

func (l *LdapClient) GetUserGroups(userDN string) ([]string, error) {
	if err := l.connectAndBind(); err != nil {
		return nil, err
	}

	return l.getGroupsByDN(userDN)
}

func (l *LdapClient) getGroupsByDN(userDN string) ([]string, error) {
	searchResult, err := l.Conn.Search(l.newSearchRequestByGroupMember(userDN))
	if err != nil {
		return nil, err
	}

	return l.getGroupsNames(searchResult), nil
}

func (l *LdapClient) newSearchRequestByGroupMember(userDN string) *ldap.SearchRequest {
	return ldap.NewSearchRequest(
		l.Base,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(member=%s)", userDN),
		[]string{"cn"},
		nil,
	)
}

func (l *LdapClient) getGroupsNames(searchResult *ldap.SearchResult) []string {
	var groups []string

	for _, entry := range searchResult.Entries {
		groups = append(groups, entry.GetAttributeValue("cn"))
	}

	return groups
}

func (l *LdapClient) IsAvailable() bool {
	if err := l.Connect(); err != nil {
		return false
	}

	_, err := l.Conn.Search(l.newSearchRequestHealthCheck())
	return err == nil
}

func (l *LdapClient) newSearchRequestHealthCheck() *ldap.SearchRequest {
	return ldap.NewSearchRequest(
		l.Base,
		ldap.ScopeBaseObject, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(%s)", l.getLdapURL()),
		[]string{},
		nil,
	)
}
