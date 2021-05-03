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

package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ZupIT/horusec-devkit/pkg/services/broker"

	"google.golang.org/grpc"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/health"

	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	t.Run("should return 204 when options", func(t *testing.T) {
		db := &database.Connection{
			Read:  &database.Mock{},
			Write: &database.Mock{},
		}
		handler := NewHealthHandler(db, &grpc.ClientConn{}, &broker.Mock{})
		r, _ := http.NewRequest(http.MethodOptions, "vulnerability/health", nil)
		w := httptest.NewRecorder()

		handler.Options(w, r)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}

func TestGet(t *testing.T) {
	t.Run("should return 200 everything its ok", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		brokerMock.On("IsAvailable").Return(true)
		mockGrpcService := &health.MockHealthCheckClient{}
		mockGrpcService.On("IsAvailable").Return(true, "")
		dbMockRead := &database.Mock{}
		dbMockRead.On("IsAvailable").Return(true)
		dbMockWrite := &database.Mock{}
		dbMockWrite.On("IsAvailable").Return(true)

		handler := Handler{
			databaseRead:           dbMockRead,
			databaseWrite:          dbMockWrite,
			broker:                 brokerMock,
			grpcHealthCheckService: mockGrpcService,
		}

		r, _ := http.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		handler.Get(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("should return 500 when broker is not alive", func(t *testing.T) {
		mockGrpcService := &health.MockHealthCheckClient{}
		mockGrpcService.On("IsAvailable").Return(true, "")
		brokerMock := &broker.Mock{}
		brokerMock.On("IsAvailable").Return(false)
		dbMockRead := &database.Mock{}
		dbMockRead.On("IsAvailable").Return(true)
		dbMockWrite := &database.Mock{}
		dbMockWrite.On("IsAvailable").Return(true)

		handler := Handler{
			databaseRead:           dbMockRead,
			databaseWrite:          dbMockWrite,
			broker:                 brokerMock,
			grpcHealthCheckService: mockGrpcService,
		}

		r, _ := http.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		handler.Get(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
	t.Run("should return 500 when database write is not healthy", func(t *testing.T) {
		mockGrpcService := &health.MockHealthCheckClient{}
		mockGrpcService.On("IsAvailable").Return(true, "")
		dbMockRead := &database.Mock{}
		dbMockRead.On("IsAvailable").Return(true)
		dbMockWrite := &database.Mock{}
		dbMockWrite.On("IsAvailable").Return(false)
		brokerMock := &broker.Mock{}
		brokerMock.On("IsAvailable").Return(true)

		handler := Handler{
			databaseRead:           dbMockRead,
			databaseWrite:          dbMockWrite,
			broker:                 brokerMock,
			grpcHealthCheckService: mockGrpcService,
		}

		r, _ := http.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		handler.Get(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 500 when database read is not healthy", func(t *testing.T) {
		mockGrpcService := &health.MockHealthCheckClient{}
		mockGrpcService.On("IsAvailable").Return(true, "")
		dbMockRead := &database.Mock{}
		dbMockRead.On("IsAvailable").Return(false)
		dbMockWrite := &database.Mock{}
		dbMockWrite.On("IsAvailable").Return(true)
		brokerMock := &broker.Mock{}
		brokerMock.On("IsAvailable").Return(true)

		handler := Handler{
			databaseRead:           dbMockRead,
			databaseWrite:          dbMockWrite,
			broker:                 brokerMock,
			grpcHealthCheckService: mockGrpcService,
		}

		r, _ := http.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		handler.Get(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 500 when failed to connect to grpc", func(t *testing.T) {
		mockGrpcService := &health.MockHealthCheckClient{}
		mockGrpcService.On("IsAvailable").Return(false, "error on connect grpc")
		dbMockRead := &database.Mock{}
		dbMockRead.On("IsAvailable").Return(true)
		dbMockWrite := &database.Mock{}
		dbMockWrite.On("IsAvailable").Return(true)
		brokerMock := &broker.Mock{}
		brokerMock.On("IsAvailable").Return(true)

		handler := Handler{
			databaseRead:           dbMockRead,
			databaseWrite:          dbMockWrite,
			broker:                 brokerMock,
			grpcHealthCheckService: mockGrpcService,
		}

		r, _ := http.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		handler.Get(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
