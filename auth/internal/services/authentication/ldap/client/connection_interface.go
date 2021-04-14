package client

import (
	"crypto/tls"
	"time"

	"github.com/go-ldap/ldap/v3"
)

type IConnection interface {
	Start()
	Close()
	SetTimeout(timeout time.Duration)
	StartTLS(config *tls.Config) error
	Search(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error)
	Bind(username, password string) error
}
