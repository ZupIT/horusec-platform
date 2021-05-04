package webhook

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-platform/webhook/internal/entities/webhook"
)

type IWebhookRepository interface {
	Save(entity *webhook.Webhook) error
	Update(entity *webhook.Webhook, webhookID uuid.UUID) error
	ListAll(workspaceID uuid.UUID) (entities *[]webhook.Webhook, err error)
	ListOne(condition map[string]interface{}) (entity *webhook.Webhook, err error)
	Remove(webhookID uuid.UUID) error
}

type Repository struct {
	dbRead  database.IDatabaseRead
	dbWrite database.IDatabaseWrite
}

func NewWebhookRepository(connection *database.Connection) IWebhookRepository {
	return &Repository{
		dbRead:  connection.Read,
		dbWrite: connection.Write,
	}
}

func (r *Repository) Save(entity *webhook.Webhook) error {
	return r.dbWrite.Create(entity, entity.GetTable()).GetErrorExceptNotFound()
}

func (r *Repository) Update(entity *webhook.Webhook, webhookID uuid.UUID) error {
	condition := map[string]interface{}{"webhook_id": webhookID}
	return r.dbWrite.Update(entity, condition, entity.GetTable()).GetError()
}

func (r *Repository) ListAll(workspaceID uuid.UUID) (entities *[]webhook.Webhook, err error) {
	condition := map[string]interface{}{"workspace_id": workspaceID}
	res := r.dbRead.Find(&[]webhook.Webhook{}, condition, (&webhook.Webhook{}).GetTable())
	if res.GetErrorExceptNotFound() != nil {
		return &[]webhook.Webhook{}, res.GetErrorExceptNotFound()
	}
	if res.GetData() == nil {
		return &[]webhook.Webhook{}, nil
	}
	return res.GetData().(*[]webhook.Webhook), nil
}

func (r *Repository) ListOne(condition map[string]interface{}) (entity *webhook.Webhook, err error) {
	res := r.dbRead.Find(&webhook.Webhook{}, condition, (&webhook.Webhook{}).GetTable())
	if res.GetErrorExceptNotFound() != nil {
		return &webhook.Webhook{}, res.GetErrorExceptNotFound()
	}
	if res.GetData() == nil {
		return &webhook.Webhook{}, nil
	}
	return res.GetData().(*webhook.Webhook), nil
}

func (r *Repository) Remove(webhookID uuid.UUID) error {
	condition := map[string]interface{}{"webhook_id": webhookID}
	return r.dbWrite.Delete(condition, (&webhook.Webhook{}).GetTable()).GetError()
}
