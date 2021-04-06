package workspace

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type IRepository interface {
	ListWorkspacesAuthTypeHorusec(accountID uuid.UUID) (*[]workspaceEntities.Response, error)
	ListWorkspacesAuthTypeLdap(permissions []string) (*[]workspaceEntities.Response, error)
}

type Repository struct {
	databaseRead  database.IDatabaseRead
	databaseWrite database.IDatabaseWrite
}

func NewWorkspaceRepository(databaseConnection *database.Connection) IRepository {
	return &Repository{
		databaseRead:  databaseConnection.Read,
		databaseWrite: databaseConnection.Write,
	}
}

func (r *Repository) ListWorkspacesAuthTypeHorusec(accountID uuid.UUID) (*[]workspaceEntities.Response, error) {
	workspaces := &[]workspaceEntities.Response{}

	return workspaces, r.databaseRead.Raw(r.queryListWorkspacesAuthTypeHorusec(), workspaces, accountID).GetError()
}

func (r *Repository) queryListWorkspacesAuthTypeHorusec() string {
	return `
			SELECT ws.workspace_id, ws.name, ws.description, aw.role, ws.created_at, ws.updated_at
			FROM workspaces AS ws
			INNER JOIN account_workspace AS aw ON aw.workspace_id = ws.workspace_id
			WHERE aw.account_id = ?
	`
}

func (r *Repository) ListWorkspacesAuthTypeLdap(permissions []string) (*[]workspaceEntities.Response, error) {
	workspaces := &[]workspaceEntities.Response{}

	response := r.databaseRead.Raw(r.queryListWorkspacesAuthTypeLdap(), workspaces, pq.StringArray(permissions), pq.StringArray(permissions))

	return workspaces, response.GetError()
}

func (r *Repository) queryListWorkspacesAuthTypeLdap() string {
	return `
			SELECT * 
			FROM (
				SELECT ws.workspace_id, ws.name, ws.description, 'admin' AS role, ws.authz_admin, 
			   	 	   ws.authz_member, ws.created_at, ws.updated_at
				FROM workspaces AS ws 
				WHERE ? && ws.authz_admin
			) AS admin

			UNION ALL

			SELECT * 
			FROM (
				SELECT ws.workspace_id, ws.name, ws.description, 'member' AS role, ws.authz_admin, 
			   		   ws.authz_member, ws.created_at, ws.updated_at
				FROM workspaces AS ws 
				WHERE ? && ws.authz_member
			) AS member 
			WHERE member.company_id NOT IN (SELECT company_id FROM admin)
	`
}

func (r *Repository) subQueryListWorkspacesAuthTypeLdapWhenAdmin() string {
	return `
		SELECT ws.workspace_id, ws.name, ws.description, 'admin' AS role, ws.authz_admin,
			   ws.authz_member, ws.created_at, ws.updated_at
		FROM workspaces AS ws
		WHERE ? && ws.authz_admin
	`
}

func (r *Repository) subQueryListWorkspacesAuthTypeLdapWhenMember() string {
	return `
		SELECT ws.workspace_id, ws.name, ws.description, 'member' AS role, ws.authz_admin,
			   ws.authz_member, ws.created_at, ws.updated_at
		FROM workspaces AS ws
		WHERE ? && ws.authz_member
	`
}
