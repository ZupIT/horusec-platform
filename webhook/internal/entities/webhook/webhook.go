package webhook

import (
	"time"

	"github.com/google/uuid"
)

type Webhook struct {
	WebhookID    uuid.UUID  `json:"webhookID" gorm:"primary_key"`
	Description  string     `json:"description"`
	URL          string     `json:"url" example:"http://my-domain.io/api"`
	Method       string     `json:"method" example:"POST" enums:"POST"`
	Headers      HeaderType `json:"headers"`
	Repository   Repository `json:"repository" gorm:"-"`
	RepositoryID uuid.UUID  `json:"repositoryID" example:"00000000-0000-0000-0000-000000000000"`
	WorkspaceID  uuid.UUID  `json:"workspaceID" example:"00000000-0000-0000-0000-000000000000"`
	CreatedAt    time.Time  `json:"createdAt" example:"2021-12-30T23:59:59Z"`
	UpdatedAt    time.Time  `json:"updatedAt" example:"2021-12-30T23:59:59Z"`
}

func (w *Webhook) GetTable() string {
	return "webhooks"
}

func (w *Webhook) GenerateID() *Webhook {
	w.WebhookID = uuid.New()
	return w
}

func (w *Webhook) GenerateCreateAt() *Webhook {
	w.CreatedAt = time.Now()
	return w
}

func (w *Webhook) GenerateUpdatedAt() *Webhook {
	w.UpdatedAt = time.Now()
	return w
}
