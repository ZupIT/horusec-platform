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
