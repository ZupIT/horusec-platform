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

package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
)

func TestNewHealthHandler(t *testing.T) {
	t.Run("should success create a new health handler", func(t *testing.T) {
		assert.NotNil(t, NewHealthHandler(&database.Connection{}, nil))
	})
}

func TestGet(t *testing.T) {
	t.Run("should return 200 when healthy", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("IsAvailable").Return(true)

		brokerMock := &broker.Mock{}
		brokerMock.On("IsAvailable").Return(true)

		handler := NewHealthHandler(&database.Connection{Read: databaseMock, Write: databaseMock}, brokerMock)

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		handler.Get(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when unhealthy broker", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("IsAvailable").Return(true)

		brokerMock := &broker.Mock{}
		brokerMock.On("IsAvailable").Return(false)

		handler := NewHealthHandler(&database.Connection{Read: databaseMock, Write: databaseMock}, brokerMock)

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		handler.Get(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 500 when unhealthy database", func(t *testing.T) {
		brokerMock := &broker.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("IsAvailable").Return(false)

		handler := NewHealthHandler(&database.Connection{Read: databaseMock, Write: databaseMock}, brokerMock)

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		handler.Get(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestOptions(t *testing.T) {
	t.Run("should return 204 when options", func(t *testing.T) {
		databaseMock := &database.Mock{}
		brokerMock := &broker.Mock{}

		handler := NewHealthHandler(&database.Connection{Read: databaseMock, Write: databaseMock}, brokerMock)

		r, _ := http.NewRequest(http.MethodOptions, "test", nil)
		w := httptest.NewRecorder()

		handler.Options(w, r)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}
