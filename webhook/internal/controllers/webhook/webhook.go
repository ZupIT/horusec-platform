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

func (c *Controller) ListAll(workspaceID uuid.UUID) (*[]webhook.Webhook, error) {
	webhooks, err := c.repository.ListAll(workspaceID)
	if err != nil {
		return &[]webhook.Webhook{}, err
	}
	allWebhooks := *webhooks
	for key := range allWebhooks {
		repoName, err := c.repository.GetRepositoryByID(allWebhooks[key].RepositoryID)
		if err != nil {
			return &[]webhook.Webhook{}, err
		}
		allWebhooks[key].Repository = webhook.Repository{Name: repoName}
	}
	return &allWebhooks, nil
}

func (c *Controller) extractRepositoriesID(webhooks *[]webhook.Webhook) (repositoriesID []uuid.UUID) {
	return repositoriesID
}

func (c *Controller) Remove(webhookID uuid.UUID) error {
	return c.repository.Remove(webhookID)
}
