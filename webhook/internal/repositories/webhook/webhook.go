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
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-platform/webhook/internal/entities/webhook"
)

type IWebhookRepository interface {
	Save(entity *webhook.Webhook) error
	Update(entity *webhook.Webhook, webhookID uuid.UUID) error
	ListAll(workspaceID uuid.UUID) (entities *[]webhook.WithRepository, err error)
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

func (r *Repository) ListAll(workspaceID uuid.UUID) (entities *[]webhook.WithRepository, err error) {
	condition := map[string]interface{}{"workspace_id": workspaceID}
	preloads := map[string][]interface{}{
		"Repository": {},
	}
	res := r.dbRead.FindPreload(&[]webhook.WithRepository{}, condition, preloads, (&webhook.WithRepository{}).GetTable())
	if res.GetErrorExceptNotFound() != nil {
		return &[]webhook.WithRepository{}, res.GetErrorExceptNotFound()
	}
	if res.GetData() == nil {
		return &[]webhook.WithRepository{}, nil
	}
	return res.GetData().(*[]webhook.WithRepository), nil
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
