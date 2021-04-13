package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"

	repositoryEntities "github.com/ZupIT/horusec-platform/core/internal/entities/repository"
	roleEntities "github.com/ZupIT/horusec-platform/core/internal/entities/role"
	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
	repositoryEnums "github.com/ZupIT/horusec-platform/core/internal/enums/repository"
	workspaceRepository "github.com/ZupIT/horusec-platform/core/internal/repositories/workspace"
	repositoriesUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/repository"
)

type IRepository interface {
	GetRepositoryByName(workspaceID uuid.UUID, name string) (*repositoryEntities.Repository, error)
	GetRepository(repositoryID uuid.UUID) (*repositoryEntities.Repository, error)
	GetAccountRepository(accountID, repositoryID uuid.UUID) (*repositoryEntities.AccountRepository, error)
	ListRepositoriesAuthTypeHorusec(accountID, workspaceID uuid.UUID) (*[]repositoryEntities.Response, error)
	ListRepositoriesAuthTypeLdap(workspaceID uuid.UUID, permissions []string) (*[]repositoryEntities.Response, error)
	IsNotMemberOfWorkspace(accountID, workspaceID uuid.UUID) bool
	ListAllRepositoryUsers(repositoryID uuid.UUID) (*[]roleEntities.Response, error)
	GetWorkspace(workspaceID uuid.UUID) (*workspaceEntities.Workspace, error)
}

type Repository struct {
	databaseRead        database.IDatabaseRead
	databaseWrite       database.IDatabaseWrite
	useCases            repositoriesUseCases.IUseCases
	workspaceRepository workspaceRepository.IRepository
}

func NewRepositoryRepository(databaseConnection *database.Connection,
	useCases repositoriesUseCases.IUseCases, repository workspaceRepository.IRepository) IRepository {
	return &Repository{
		databaseRead:        databaseConnection.Read,
		databaseWrite:       databaseConnection.Write,
		useCases:            useCases,
		workspaceRepository: repository,
	}
}

func (r *Repository) GetRepositoryByName(workspaceID uuid.UUID, name string) (*repositoryEntities.Repository, error) {
	repository := &repositoryEntities.Repository{}

	return repository, r.databaseRead.Find(repository, r.useCases.FilterRepositoryByNameAndWorkspace(workspaceID, name),
		repositoryEnums.DatabaseRepositoryTable).GetError()
}

func (r *Repository) GetRepository(repositoryID uuid.UUID) (*repositoryEntities.Repository, error) {
	repository := &repositoryEntities.Repository{}

	return repository, r.databaseRead.Find(repository, r.useCases.FilterRepositoryByID(repositoryID),
		repositoryEnums.DatabaseRepositoryTable).GetError()
}

func (r *Repository) GetAccountRepository(accountID,
	repositoryID uuid.UUID) (*repositoryEntities.AccountRepository, error) {
	accountRepository := &repositoryEntities.AccountRepository{}

	return accountRepository, r.databaseRead.Find(accountRepository, r.useCases.FilterAccountRepositoryByID(
		accountID, repositoryID), repositoryEnums.DatabaseAccountRepositoryTable).GetError()
}

func (r *Repository) ListRepositoriesAuthTypeHorusec(accountID,
	workspaceID uuid.UUID) (*[]repositoryEntities.Response, error) {
	accountWorkspace, err := r.workspaceRepository.GetAccountWorkspace(accountID, workspaceID)
	if err != nil {
		return nil, err
	}

	if accountWorkspace.Role == account.Admin {
		return r.listRepositoriesWhenWorkspaceAdmin(accountID, workspaceID)
	}

	return r.listRepositoriesByRoles(accountID, workspaceID)
}

func (r *Repository) listRepositoriesWhenWorkspaceAdmin(accountID,
	workspaceID uuid.UUID) (*[]repositoryEntities.Response, error) {
	repositories := &[]repositoryEntities.Response{}

	return repositories, r.databaseRead.Raw(
		r.queryListRepositoriesWhenWorkspaceAdmin(), repositories, accountID, workspaceID).GetErrorExceptNotFound()
}

func (r *Repository) queryListRepositoriesWhenWorkspaceAdmin() string {
	return `
			SELECT repo.repository_id, repo.workspace_id, repo.description, repo.name, 'admin' AS role, 
				   repo.created_at, repo.updated_at
			FROM repositories AS repo
		    INNER JOIN account_workspace AS aw ON aw.workspace_id = repo.workspace_id AND aw.account_id = ?
			WHERE repo.workspace_id = ?
	`
}

func (r *Repository) listRepositoriesByRoles(accountID,
	workspaceID uuid.UUID) (*[]repositoryEntities.Response, error) {
	repositories := &[]repositoryEntities.Response{}

	return repositories, r.databaseRead.Raw(r.queryListRepositoriesByRoles(), repositories,
		sql.Named("accountID", accountID), sql.Named("workspaceID", workspaceID)).GetErrorExceptNotFound()
}

func (r *Repository) queryListRepositoriesByRoles() string {
	return `
			SELECT repo.repository_id, repo.workspace_id, repo.description, repo.name, ar.role,
			  	   repo.created_at, repo.updated_at
		    FROM repositories AS repo
			INNER JOIN account_repository AS ar ON ar.repository_id = repo.repository_id AND ar.account_id = @accountID
			WHERE ar.workspace_id = @workspaceID AND accountRepo.account_id = @accountID
	`
}

func (r *Repository) ListRepositoriesAuthTypeLdap(workspaceID uuid.UUID,
	permissions []string) (*[]repositoryEntities.Response, error) {
	repositories := &[]repositoryEntities.Response{}

	return repositories, r.databaseRead.Raw(r.queryListRepositoriesAuthTypeLdap(), repositories,
		sql.Named("workspaceID", workspaceID),
		sql.Named("permissions", pq.StringArray(permissions))).GetErrorExceptNotFound()
}

//nolint:funlen // query needs more than 15 lines
func (r *Repository) queryListRepositoriesAuthTypeLdap() string {
	return `
			SELECT * 
			FROM (
				SELECT repo.repository_id, repo.workspace_id, repo.description, repo.name, 'admin' AS role,
					   repo.authz_admin, repo.authz_member, repo.authz_supervisor, repo.created_at, repo.updated_at
				FROM repositories AS repo
				WHERE repo.workspace_id = @workspaceID AND @permissions && repo.authz_admin
			) AS admin

			UNION ALL (
				SELECT * FROM (
					SELECT repo.repository_id, repo.workspace_id, repo.description, repo.name, 'supervisor' AS role,
					       repo.authz_admin, repo.authz_member, repo.authz_supervisor, repo.created_at, repo.updated_at
					FROM repositories AS repo
					WHERE repo.workspace_id = @workspaceID AND @permissions && repo.authz_supervisor
				) AS supervisor
				WHERE supervisor.repository_id NOT IN (SELECT repo.workspace_id FROM repositories AS repo 
					  WHERE repo.workspace_id = @workspaceID AND @permissions && repo.authz_admin) 

				UNION ALL

				SELECT * FROM (
					SELECT repo.repository_id, repo.workspace_id, repo.description, repo.name, 'member' AS role,
						   repo.authz_admin, repo.authz_member, repo.authz_supervisor, repo.created_at, repo.updated_at
					FROM repositories AS repo
					WHERE repo.workspace_id = @workspaceID AND @permissions && repo.authz_member	
				) AS member
					WHERE member.repository_id 
					NOT IN (
						SELECT repo.workspace_id FROM repositories AS repo 
						WHERE repo.workspace_id = @workspaceID AND @permissions && repo.authz_admin
					) 
					AND member.repository_id 
					NOT IN (
						SELECT repo.workspace_id FROM repositories AS repo 
						WHERE repo.workspace_id = @workspaceID AND @permissions && repo.authz_supervisor
					)
			)
	`
}

func (r *Repository) IsNotMemberOfWorkspace(accountID, workspaceID uuid.UUID) bool {
	accountWorkspace, err := r.workspaceRepository.GetAccountWorkspace(accountID, workspaceID)
	if err != nil || accountWorkspace == nil {
		return true
	}

	return false
}

func (r *Repository) ListAllRepositoryUsers(repositoryID uuid.UUID) (*[]roleEntities.Response, error) {
	users := &[]roleEntities.Response{}

	return users, r.databaseRead.Raw(r.queryListAllRepositoryUsers(), users, repositoryID).GetErrorExceptNotFound()
}

func (r *Repository) queryListAllRepositoryUsers() string {
	return `
			SELECT ac.email, ac.username, ar.role, ac.account_id
			FROM accounts AS ac
			INNER JOIN account_repository AS ar ON ar.account_id = ac.account_id
			WHERE ar.repository_id = ?
	`
}

func (r *Repository) GetWorkspace(workspaceID uuid.UUID) (*workspaceEntities.Workspace, error) {
	return r.workspaceRepository.GetWorkspace(workspaceID)
}
