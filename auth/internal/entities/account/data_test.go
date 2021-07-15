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

func TestValidate(t *testing.T) {
	t.Run("should return no error when valid data", func(t *testing.T) {
		data := &Data{
			Email:    "test@test.com",
			Password: "Test@123",
			Username: "test",
		}

		assert.NoError(t, data.Validate())
	})

	t.Run("should return error when invalid data email", func(t *testing.T) {
		data := &Data{
			Email:    "test",
			Password: "Test@123",
			Username: "test",
		}

		assert.Error(t, data.Validate())
	})

	t.Run("should return error when invalid data password", func(t *testing.T) {
		data := &Data{
			Email:    "test@test.com",
			Password: "test",
			Username: "test",
		}

		assert.Error(t, data.Validate())
	})
}

func TestToAccount(t *testing.T) {
	t.Run("should success parse data to account", func(t *testing.T) {
		data := &Data{
			Email:    "test@test.com",
			Password: "Test@123",
			Username: "test",
		}

		account := data.ToAccount()
		assert.Equal(t, data.Email, account.Email)
		assert.NotEmpty(t, account.Password)
		assert.Equal(t, data.Username, account.Username)
	})
}

func TestToBytesData(t *testing.T) {
	t.Run("should success parse to bytes", func(t *testing.T) {
		data := &Data{
			Email:    "test",
			Password: "test",
			Username: "test",
		}

		assert.NotEmpty(t, data.ToBytes())
	})
}
