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
