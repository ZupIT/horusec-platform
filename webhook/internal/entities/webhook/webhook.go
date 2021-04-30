package webhook

import (
	"time"

	"github.com/google/uuid"
)

type Webhook struct {
	WebhookID    uuid.UUID  `json:"webhookID" gorm:"primary_key"`
	Description  string     `json:"description"`
	URL          string     `json:"url"`
	Method       string     `json:"method" enums:"post"`
	Headers      HeaderType `json:"headers"`
	RepositoryID uuid.UUID  `json:"repositoryID"`
	WorkspaceID  uuid.UUID  `json:"workspaceID"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
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
