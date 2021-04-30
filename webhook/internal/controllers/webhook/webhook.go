package webhook

import (
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-platform/webhook/internal/entities/webhook"
	repositoryWebhook "github.com/ZupIT/horusec-platform/webhook/internal/repositories/webhook"
)

type IWebhookController interface {
	Save(entity *webhook.Webhook) (uuid.UUID, error)
	Update(entity *webhook.Webhook, webhookID uuid.UUID) error
	ListAll(workspaceID uuid.UUID) (*[]webhook.Webhook, error)
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
	entity = entity.GenerateID().GenerateCreateAt()
	err := c.repository.Save(entity)
	if err != nil {
		return uuid.Nil, err
	}
	return entity.WebhookID, nil
}

func (c *Controller) Update(entity *webhook.Webhook, webhookID uuid.UUID) error {
	return c.repository.Update(entity, webhookID)
}

func (c *Controller) ListAll(workspaceID uuid.UUID) (*[]webhook.Webhook, error) {
	return c.repository.ListAll(workspaceID)
}

func (c *Controller) Remove(webhookID uuid.UUID) error {
	return c.repository.Remove(webhookID)
}
