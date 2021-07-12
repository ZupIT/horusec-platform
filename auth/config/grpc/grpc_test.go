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

package grpc

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-platform/auth/config/grpc/enums"
	authHandler "github.com/ZupIT/horusec-platform/auth/internal/handlers/authentication"
)

func TestNewAuthGRPCServer(t *testing.T) {
	t.Run("should success create server without certs", func(t *testing.T) {
		server := NewAuthGRPCServer(&authHandler.Handler{})
		assert.NotNil(t, server)
	})

	t.Run("should panic when failed to get certs", func(t *testing.T) {
		_ = os.Setenv(enums.EnvGrpcUseCerts, "true")

		assert.Panics(t, func() {
			_ = NewAuthGRPCServer(&authHandler.Handler{})
		})
	})
}
