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

package cors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCorsConfig(t *testing.T) {
	t.Run("should success create a new cors config", func(t *testing.T) {
		config := NewCorsConfig()
		assert.Equal(t, []string{"*"}, config.AllowedOrigins)
		assert.Equal(t, []string{"GET", "OPTIONS"}, config.AllowedMethods)
		assert.Equal(t, []string{"Accept", "headers", "X-Horusec-Authorization", "Content-Type"}, config.AllowedHeaders)
		assert.True(t, config.AllowCredentials)
		assert.Equal(t, 300, config.MaxAge)
	})
}
