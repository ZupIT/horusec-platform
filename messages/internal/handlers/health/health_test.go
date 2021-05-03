package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/broker"

	"github.com/ZupIT/horusec-platform/messages/internal/services/mailer"
)

func TestNewHealthHandler(t *testing.T) {
	t.Run("should success create a new health handler", func(t *testing.T) {
		assert.NotNil(t, NewHealthHandler(nil, nil))
	})
}

func TestGet(t *testing.T) {
	t.Run("should return 200 when healthy", func(t *testing.T) {
		mailerMock := &mailer.Mock{}
		mailerMock.On("IsAvailable").Return(true)

		brokerMock := &broker.Mock{}
		brokerMock.On("IsAvailable").Return(true)

		handler := NewHealthHandler(brokerMock, mailerMock)

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		handler.Get(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when unhealthy broker", func(t *testing.T) {
		mailerMock := &mailer.Mock{}
		mailerMock.On("IsAvailable").Return(true)

		brokerMock := &broker.Mock{}
		brokerMock.On("IsAvailable").Return(false)

		handler := NewHealthHandler(brokerMock, mailerMock)

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		handler.Get(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 500 when unhealthy mailer", func(t *testing.T) {
		brokerMock := &broker.Mock{}

		mailerMock := &mailer.Mock{}
		mailerMock.On("IsAvailable").Return(false)

		handler := NewHealthHandler(brokerMock, mailerMock)

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		handler.Get(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestOptions(t *testing.T) {
	t.Run("should return 204 when options", func(t *testing.T) {
		mailerMock := &mailer.Mock{}
		brokerMock := &broker.Mock{}

		handler := NewHealthHandler(brokerMock, mailerMock)

		r, _ := http.NewRequest(http.MethodOptions, "test", nil)
		w := httptest.NewRecorder()

		handler.Options(w, r)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}
