package workspace

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"

	roleEntities "github.com/ZupIT/horusec-platform/core/internal/entities/role"
	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
	workspaceEnums "github.com/ZupIT/horusec-platform/core/internal/enums/workspace"
	workspaceUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/workspace"
)

type IRepository interface {
	ListWorkspacesAuthTypeHorusec(accountID uuid.UUID) (*[]workspaceEntities.Response, error)
	ListWorkspacesAuthTypeLdap(permissions []string) (*[]workspaceEntities.Response, error)
	GetWorkspace(workspaceID uuid.UUID) (*workspaceEntities.Workspace, error)
	GetAccountWorkspace(accountID, workspaceID uuid.UUID) (*workspaceEntities.AccountWorkspace, error)
	ListAllWorkspaceUsers(workspaceID uuid.UUID) (*[]roleEntities.Response, error)
}

type Repository struct {
	databaseRead  database.IDatabaseRead
	databaseWrite database.IDatabaseWrite
	useCases      workspaceUseCases.IUseCases
}

func NewWorkspaceRepository(databaseConnection *database.Connection, useCases workspaceUseCases.IUseCases) IRepository {
	return &Repository{
		databaseRead:  databaseConnection.Read,
		databaseWrite: databaseConnection.Write,
		useCases:      useCases,
	}
}

func (r *Repository) GetWorkspace(workspaceID uuid.UUID) (*workspaceEntities.Workspace, error) {
	workspace := &workspaceEntities.Workspace{}

	return workspace, r.databaseRead.Find(workspace, r.useCases.FilterWorkspaceByID(workspaceID),
		workspaceEnums.DatabaseWorkspaceTable).GetError()
}

func (r *Repository) GetAccountWorkspace(accountID,
	workspaceID uuid.UUID) (*workspaceEntities.AccountWorkspace, error) {
	accountWorkspace := &workspaceEntities.AccountWorkspace{}

	return accountWorkspace, r.databaseRead.Find(accountWorkspace, r.useCases.FilterAccountWorkspaceByID(
		accountID, workspaceID), workspaceEnums.DatabaseAccountWorkspaceTable).GetError()
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

	return workspaces, r.databaseRead.Raw(r.queryListWorkspacesAuthTypeLdap(), workspaces,
		sql.Named("permissions", pq.StringArray(permissions))).GetError()
}

//nolint:funlen // query needs more than 15 lines
func (r *Repository) queryListWorkspacesAuthTypeLdap() string {
	return `
			SELECT * 
			FROM (
				SELECT ws.workspace_id, ws.name, ws.description, 'admin' AS role, ws.authz_admin, 
			   	 	   ws.authz_member, ws.created_at, ws.updated_at
				FROM workspaces AS ws 
				WHERE @permissions && ws.authz_admin
			) AS admin

			UNION ALL

			SELECT * 
			FROM (
				SELECT ws.workspace_id, ws.name, ws.description, 'member' AS role, ws.authz_admin, 
			   		   ws.authz_member, ws.created_at, ws.updated_at
				FROM workspaces AS ws 
				WHERE @permissions && ws.authz_member
			) AS member 
			WHERE member.workspace_id NOT IN (
			      SELECT ws.workspace_id FROM workspaces AS ws WHERE @permissions && ws.authz_admin)
	`
}

func (r *Repository) ListAllWorkspaceUsers(workspaceID uuid.UUID) (*[]roleEntities.Response, error) {
	users := &[]roleEntities.Response{}

	return users, r.databaseRead.Raw(r.queryListAllWorkspaceUsers(), users, workspaceID).GetError()
}

func (r *Repository) queryListAllWorkspaceUsers() string {
	return `
			SELECT ac.email, ac.username, aw.role, ac.account_id
			FROM accounts AS ac
			INNER JOIN account_workspace AS aw ON ws.account_id = ac.account_id
			WHERE aw.workspace_id = ?
	`
}
