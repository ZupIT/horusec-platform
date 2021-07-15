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
	"time"

	"github.com/google/uuid"
)

type Webhook struct {
	WebhookID    uuid.UUID  `json:"webhookID" gorm:"primary_key"`
	Description  string     `json:"description"`
	URL          string     `json:"url" example:"http://my-domain.io/api"`
	Method       string     `json:"method" example:"POST" enums:"POST"`
	Headers      HeaderType `json:"headers"`
	RepositoryID uuid.UUID  `json:"repositoryID" example:"00000000-0000-0000-0000-000000000000"`
	WorkspaceID  uuid.UUID  `json:"workspaceID" example:"00000000-0000-0000-0000-000000000000"`
	CreatedAt    time.Time  `json:"createdAt" example:"2021-12-30T23:59:59Z"`
	UpdatedAt    time.Time  `json:"updatedAt" example:"2021-12-30T23:59:59Z"`
}

type WithRepository struct {
	Webhook
	Repository Repository `json:"repository" gorm:"foreignKey:RepositoryID;references:RepositoryID"`
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
