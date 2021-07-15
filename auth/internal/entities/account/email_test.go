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

package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmail(t *testing.T) {
	t.Run("should return no error when valid data", func(t *testing.T) {
		data := &Email{
			Email: "test@test.com",
		}

		assert.NoError(t, data.Validate())
	})

	t.Run("should return error when invalid data", func(t *testing.T) {
		data := &Email{
			Email: "test",
		}

		assert.Error(t, data.Validate())
	})
}

func TestToBytesEmail(t *testing.T) {
	t.Run("should success parse to bytes", func(t *testing.T) {
		data := &Email{
			Email: "test@test.com",
		}

		assert.NotEmpty(t, data.ToBytes())
	})
}
