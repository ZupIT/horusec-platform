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
	"errors"
	"testing"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-platform/webhook/internal/entities/webhook"
)

func TestRepository_ListAll(t *testing.T) {
	t.Run("Should return all webhooks without errors", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("FindPreload").Return(response.NewResponse(0, nil, &[]webhook.WithRepository{
			{
				Webhook: webhook.Webhook{WebhookID: uuid.New()},
			},
		}))
		dbWrite := &database.Mock{}
		connection := &database.Connection{
			Read:  dbRead,
			Write: dbWrite,
		}
		res, err := NewWebhookRepository(connection).ListAll(uuid.New())
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})
	t.Run("Should return error unknown on list all webhooks", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("FindPreload").Return(response.NewResponse(0, errors.New("unexpected error"), nil))
		dbWrite := &database.Mock{}
		connection := &database.Connection{
			Read:  dbRead,
			Write: dbWrite,
		}
		res, err := NewWebhookRepository(connection).ListAll(uuid.New())
		assert.Error(t, err)
		assert.Empty(t, res)
	})
	t.Run("Should return not error but return empty list if data is nil", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("FindPreload").Return(response.NewResponse(0, nil, nil))
		dbWrite := &database.Mock{}
		connection := &database.Connection{
			Read:  dbRead,
			Write: dbWrite,
		}
		res, err := NewWebhookRepository(connection).ListAll(uuid.New())
		assert.NoError(t, err)
		assert.Empty(t, res)
	})
}

func TestRepository_ListOne(t *testing.T) {
	t.Run("Should return one webhooks without errors", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Find").Return(response.NewResponse(0, nil, &webhook.Webhook{
			WebhookID: uuid.New(),
		}))
		dbWrite := &database.Mock{}
		connection := &database.Connection{
			Read:  dbRead,
			Write: dbWrite,
		}
		res, err := NewWebhookRepository(connection).ListOne(map[string]interface{}{"webhook_id": uuid.New()})
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})
	t.Run("Should return error unknown on list one webhooks", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Find").Return(response.NewResponse(0, errors.New("unexpected error"), nil))
		dbWrite := &database.Mock{}
		connection := &database.Connection{
			Read:  dbRead,
			Write: dbWrite,
		}
		res, err := NewWebhookRepository(connection).ListOne(map[string]interface{}{"webhook_id": uuid.New()})
		assert.Error(t, err)
		assert.Empty(t, res)
	})
	t.Run("Should return not error but return empty item if data is nil", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Find").Return(response.NewResponse(0, nil, nil))
		dbWrite := &database.Mock{}
		connection := &database.Connection{
			Read:  dbRead,
			Write: dbWrite,
		}
		res, err := NewWebhookRepository(connection).ListOne(map[string]interface{}{"webhook_id": uuid.New()})
		assert.NoError(t, err)
		assert.Empty(t, res)
	})
}

func TestRepository_Save(t *testing.T) {
	t.Run("Should save new webhook without error", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbWrite := &database.Mock{}
		dbWrite.On("Create").Return(response.NewResponse(0, nil, nil))
		connection := &database.Connection{
			Read:  dbRead,
			Write: dbWrite,
		}
		err := NewWebhookRepository(connection).Save(&webhook.Webhook{})
		assert.NoError(t, err)
	})
	t.Run("Should save new webhook with error", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbWrite := &database.Mock{}
		dbWrite.On("Create").Return(response.NewResponse(0, errors.New("unexpected error"), nil))
		connection := &database.Connection{
			Read:  dbRead,
			Write: dbWrite,
		}
		err := NewWebhookRepository(connection).Save(&webhook.Webhook{})
		assert.Error(t, err)
	})
}

func TestRepository_Update(t *testing.T) {
	t.Run("Should update webhook without error", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbWrite := &database.Mock{}
		dbWrite.On("Update").Return(response.NewResponse(0, nil, nil))
		connection := &database.Connection{
			Read:  dbRead,
			Write: dbWrite,
		}
		err := NewWebhookRepository(connection).Update(&webhook.Webhook{}, uuid.New())
		assert.NoError(t, err)
	})
	t.Run("Should update webhook with error", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbWrite := &database.Mock{}
		dbWrite.On("Update").Return(response.NewResponse(0, errors.New("unexpected error"), nil))
		connection := &database.Connection{
			Read:  dbRead,
			Write: dbWrite,
		}
		err := NewWebhookRepository(connection).Update(&webhook.Webhook{}, uuid.New())
		assert.Error(t, err)
	})
}

func TestRepository_Remove(t *testing.T) {
	t.Run("Should remove repository without error", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbWrite := &database.Mock{}
		dbWrite.On("Delete").Return(response.NewResponse(0, nil, nil))
		connection := &database.Connection{
			Read:  dbRead,
			Write: dbWrite,
		}
		err := NewWebhookRepository(connection).Remove(uuid.New())
		assert.NoError(t, err)
	})
	t.Run("Should remove webhook with error", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbWrite := &database.Mock{}
		dbWrite.On("Delete").Return(response.NewResponse(0, errors.New("unexpected error"), nil))
		connection := &database.Connection{
			Read:  dbRead,
			Write: dbWrite,
		}
		err := NewWebhookRepository(connection).Remove(uuid.New())
		assert.Error(t, err)
	})
}
