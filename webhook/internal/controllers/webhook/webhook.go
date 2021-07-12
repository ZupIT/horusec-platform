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
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-platform/webhook/internal/enums"

	"github.com/ZupIT/horusec-platform/webhook/internal/entities/webhook"
	repositoryWebhook "github.com/ZupIT/horusec-platform/webhook/internal/repositories/webhook"
)

type IWebhookController interface {
	Save(entity *webhook.Webhook) (uuid.UUID, error)
	Update(entity *webhook.Webhook, webhookID uuid.UUID) error
	ListAll(workspaceID uuid.UUID) (*[]webhook.WithRepository, error)
	Remove(webhookID uuid.UUID) error
}

type Controller struct {
	repository repositoryWebhook.IWebhookRepository
}

func NewWebhookController(repository repositoryWebhook.IWebhookRepository) IWebhookController {
	return &Controller{
		repository: repository,
	}
}

func (c *Controller) Save(entity *webhook.Webhook) (uuid.UUID, error) {
	existing, err := c.repository.ListOne(map[string]interface{}{"repository_id": entity.RepositoryID})
	if err != nil {
		return uuid.Nil, err
	}
	if existing.WebhookID != uuid.Nil {
		return uuid.Nil, enums.ErrorWebhookDuplicate
	}
	entity = entity.GenerateID().GenerateCreateAt()
	if err := c.repository.Save(entity); err != nil {
		return uuid.Nil, err
	}
	return entity.WebhookID, nil
}

func (c *Controller) Update(entity *webhook.Webhook, webhookID uuid.UUID) error {
	entity = entity.GenerateUpdatedAt()
	return c.repository.Update(entity, webhookID)
}

func (c *Controller) ListAll(workspaceID uuid.UUID) (*[]webhook.WithRepository, error) {
	return c.repository.ListAll(workspaceID)
}

func (c *Controller) Remove(webhookID uuid.UUID) error {
	return c.repository.Remove(webhookID)
}
