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

package webhook

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestWebhook(t *testing.T) {
	t.Run("Should get table of webhook", func(t *testing.T) {
		wh := &Webhook{}
		assert.Equal(t, "webhooks", wh.GetTable())
	})
	t.Run("Should generate webhook ID", func(t *testing.T) {
		wh := &Webhook{}
		assert.NotEqual(t, uuid.Nil, wh.GenerateID())
	})
	t.Run("Should generate createdAt", func(t *testing.T) {
		wh := &Webhook{}
		assert.NotEqual(t, time.Time{}, wh.GenerateCreateAt())
	})
	t.Run("Should generate updatedAt", func(t *testing.T) {
		wh := &Webhook{}
		assert.NotEqual(t, time.Time{}, wh.GenerateUpdatedAt())
	})
}
