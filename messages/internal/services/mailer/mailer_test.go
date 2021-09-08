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

package mailer

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/gomail.v2"

	"github.com/ZupIT/horusec-platform/messages/internal/services/mailer/config/enums"
)

func TestNewMailerService(t *testing.T) {
	t.Run("should success create a new mailer service", func(t *testing.T) {
		_ = os.Setenv(enums.EnvSMTPUsername, "test")
		_ = os.Setenv(enums.EnvSMTPPassword, "test")
		_ = os.Setenv(enums.EnvSMTPHost, "test")

		assert.NotNil(t, NewMailerService())
	})

	t.Run("should panic when invalid config", func(t *testing.T) {
		_ = os.Setenv(enums.EnvSMTPUsername, "")
		_ = os.Setenv(enums.EnvSMTPPassword, "")
		_ = os.Setenv(enums.EnvSMTPHost, "")

		assert.Panics(t, func() {
			_ = NewMailerService()
		})
	})
}

func TestSendEmail(t *testing.T) {
	t.Run("should return error when failed to send email", func(t *testing.T) {
		_ = os.Setenv(enums.EnvSMTPUsername, "test")
		_ = os.Setenv(enums.EnvSMTPPassword, "test")
		_ = os.Setenv(enums.EnvSMTPHost, "test")

		service := NewMailerService()

		assert.Error(t, service.SendEmail(&gomail.Message{}))
	})
}

func TestNoop(t *testing.T) {
	t.Run("should return error when failed to noop", func(t *testing.T) {
		_ = os.Setenv(enums.EnvSMTPUsername, "test")
		_ = os.Setenv(enums.EnvSMTPPassword, "test")
		_ = os.Setenv(enums.EnvSMTPHost, "test")

		service := NewMailerService()

		assert.Error(t, service.Noop())
	})
}

func TestIsAvailable(t *testing.T) {
	t.Run("should return false when it is not available", func(t *testing.T) {
		_ = os.Setenv(enums.EnvSMTPUsername, "test")
		_ = os.Setenv(enums.EnvSMTPPassword, "test")
		_ = os.Setenv(enums.EnvSMTPHost, "test")

		service := NewMailerService()

		assert.False(t, service.IsAvailable())
	})
}

func TestGetFromHeader(t *testing.T) {
	t.Run("should success get from header", func(t *testing.T) {
		_ = os.Setenv(enums.EnvSMTPUsername, "test")
		_ = os.Setenv(enums.EnvSMTPPassword, "test")
		_ = os.Setenv(enums.EnvSMTPHost, "test")

		service := NewMailerService()

		assert.Equal(t, "horusec@zup.com.br", service.GetFromHeader())
	})
}
