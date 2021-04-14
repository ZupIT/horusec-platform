package client

import (
	"crypto/tls"
	"time"

	mockUtils "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
	"github.com/go-ldap/ldap/v3"
	"github.com/stretchr/testify/mock"
)

type MockConnection struct {
	mock.Mock
}

func (m *MockConnection) Start() {
	_ = m.MethodCalled("Start")
}

func (m *MockConnection) Close() {
	_ = m.MethodCalled("Close")
}

func (m *MockConnection) SetTimeout(_ time.Duration) {
	_ = m.MethodCalled("SetTimeout")
}

func (m *MockConnection) StartTLS(_ *tls.Config) error {
	args := m.MethodCalled("StartTLS")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (m *MockConnection) Search(_ *ldap.SearchRequest) (*ldap.SearchResult, error) {
	args := m.MethodCalled("Search")
	return args.Get(0).(*ldap.SearchResult), mockUtils.ReturnNilOrError(args, 1)
}

func (m *MockConnection) Bind(_, _ string) error {
	args := m.MethodCalled("Bind")
	return mockUtils.ReturnNilOrError(args, 0)
}
