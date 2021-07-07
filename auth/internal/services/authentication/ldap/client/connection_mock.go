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
