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

package email

import (
	"testing"

	"github.com/stretchr/testify/assert"

	emailEntities "github.com/ZupIT/horusec-devkit/pkg/entities/email"
	emailEnums "github.com/ZupIT/horusec-devkit/pkg/enums/email"

	"github.com/ZupIT/horusec-platform/messages/internal/services/mailer"
)

func TestNewController(t *testing.T) {
	t.Run("should success create a new controller", func(t *testing.T) {
		assert.NotNil(t, NewEmailController(nil))
	})
}

func TestSendEmail(t *testing.T) {
	t.Run("should success send email", func(t *testing.T) {
		mailerMock := &mailer.Mock{}
		mailerMock.On("SendEmail").Return(nil)
		mailerMock.On("GetFromHeader").Return("test")

		controller := NewEmailController(mailerMock)

		message := &emailEntities.Message{TemplateName: emailEnums.AccountConfirmation}
		assert.NoError(t, controller.SendEmail(message))
	})

	t.Run("should return error when failed to execute template", func(t *testing.T) {
		mailerMock := &mailer.Mock{}

		controller := NewEmailController(mailerMock)

		message := &emailEntities.Message{}
		assert.Error(t, controller.SendEmail(message))
	})
}
