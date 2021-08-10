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

package repository

import (
	"context"
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/services/tracer"

	"github.com/opentracing/opentracing-go"

	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"

	"github.com/ZupIT/horusec-platform/api/internal/entities/core"
	"github.com/ZupIT/horusec-platform/api/internal/enums"
)

type IRepository interface {
	CreateRepository(ctx context.Context, ID, workspaceID uuid.UUID, name string) error
	FindRepository(ctx context.Context, workspaceID uuid.UUID, name string) (uuid.UUID, error)
}

type Repository struct {
	databaseWrite database.IDatabaseWrite
	databaseRead  database.IDatabaseRead
}

func NewRepositoriesRepository(connection *database.Connection) IRepository {
	return &Repository{
		databaseWrite: connection.Write,
		databaseRead:  connection.Read,
	}
}

func (r *Repository) FindRepository(ctx context.Context, workspaceID uuid.UUID, name string) (uuid.UUID, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "FindRepository")
	defer span.Finish()
	repository := &core.Repository{}
	condition := map[string]interface{}{
		"workspace_id": workspaceID,
		"name":         name,
	}

	return repository.RepositoryID, r.databaseRead.Find(repository, condition,
		enums.DatabaseRepositoryTable).GetError()
}

// nolint
func (r *Repository) CreateRepository(ctx context.Context, repositoryID, workspaceID uuid.UUID, name string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateRepository")
	defer span.Finish()
	workspace, err := r.GetWorkspace(ctx, workspaceID)
	if err != nil {
		tracer.SetSpanError(span, err)
		return err
	}
	entity := map[string]interface{}{
		"repository_id":    repositoryID,
		"workspace_id":     workspaceID,
		"name":             name,
		"created_at":       time.Now(),
		"updated_at":       time.Now(),
		"authz_member":     workspace.AuthzMember,
		"authz_admin":      workspace.AuthzAdmin,
		"authz_supervisor": workspace.AuthzAdmin,
	}
	return r.databaseWrite.Create(entity, enums.DatabaseRepositoryTable).GetError()
}

func (r *Repository) GetWorkspace(ctx context.Context, workspaceID uuid.UUID) (*core.Workspace, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "GetWorkspace")
	defer span.Finish()
	workspace := &core.Workspace{}
	condition := map[string]interface{}{
		"workspace_id": workspaceID,
	}

	return workspace, r.databaseRead.Find(workspace, condition,
		enums.DatabaseWorkspaceTable).GetError()
}
