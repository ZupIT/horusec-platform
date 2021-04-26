package dashboard

import (
	"errors"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRepoDashboard_Save(t *testing.T) {
	t.Run("Should save new content without error", func(t *testing.T) {
		dbWrite := &database.Mock{}
		dbWrite.On("Create").Return(response.NewResponse(0, nil, nil))
		conn := &database.Connection{
			Read:  &database.Mock{},
			Write: dbWrite,
		}
		repo := NewRepoDashboard(conn)
		entity := &dashboard.VulnerabilitiesByAuthor{
			Author: "horusec@zup.com.br",
			Vulnerability: dashboard.Vulnerability{
				VulnerabilityID:       uuid.New(),
				CreatedAt:             time.Now(),
				Active:                true,
				WorkspaceID:           uuid.New(),
				RepositoryID:          uuid.New(),
			},
		}
		err := repo.Save(entity, entity.GetTable())
		assert.NoError(t, err)
	})
	t.Run("Should not save new content because exists error", func(t *testing.T) {
		dbWrite := &database.Mock{}
		dbWrite.On("Create").Return(response.NewResponse(0, errors.New("unknown error"), nil))
		conn := &database.Connection{
			Read:  &database.Mock{},
			Write: dbWrite,
		}
		repo := NewRepoDashboard(conn)
		entity := &dashboard.VulnerabilitiesByAuthor{
			Author: "horusec@zup.com.br",
			Vulnerability: dashboard.Vulnerability{
				VulnerabilityID:       uuid.New(),
				CreatedAt:             time.Now(),
				Active:                true,
				WorkspaceID:           uuid.New(),
				RepositoryID:          uuid.New(),
			},
		}
		err := repo.Save(entity, entity.GetTable())
		assert.Error(t, err)
	})
}

func TestRepoDashboard_Inactive(t *testing.T) {
	t.Run("Should save new content without error", func(t *testing.T) {
		dbWrite := &database.Mock{}
		dbWrite.On("Update").Return(response.NewResponse(0, nil, nil))
		conn := &database.Connection{
			Read:  &database.Mock{},
			Write: dbWrite,
		}
		repo := NewRepoDashboard(conn)
		condition := map[string]interface{}{
			"active": true,
			"repository_id": uuid.New(),
		}
		err := repo.Inactive(condition, (&dashboard.VulnerabilitiesByAuthor{}).GetTable())
		assert.NoError(t, err)
	})
	t.Run("Should not save new content because exists error", func(t *testing.T) {
		dbWrite := &database.Mock{}
		dbWrite.On("Update").Return(response.NewResponse(0, errors.New("unknown error"), nil))
		conn := &database.Connection{
			Read:  &database.Mock{},
			Write: dbWrite,
		}
		repo := NewRepoDashboard(conn)
		condition := map[string]interface{}{
			"active": true,
			"repository_id": uuid.New(),
		}
		err := repo.Inactive(condition, (&dashboard.VulnerabilitiesByAuthor{}).GetTable())
		assert.Error(t, err)
	})
}

func TestRepoDashboard_GetDashboardTotalDevelopers(t *testing.T) {
	t.Run("Should return TotalDevelopers without error", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, nil, 1))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &dashboard.FilterDashboard{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		_, err := repo.GetDashboardTotalDevelopers(filter)
		assert.NoError(t, err)
	})
	t.Run("Should return TotalDevelopers without error if error is ErrorNotFoundRecords", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, enums.ErrorNotFoundRecords, 0))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &dashboard.FilterDashboard{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		_, err := repo.GetDashboardTotalDevelopers(filter)
		assert.NoError(t, err)
	})
	t.Run("Should return TotalDevelopers with error", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, errors.New("unknown error"), 0))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &dashboard.FilterDashboard{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		_, err := repo.GetDashboardTotalDevelopers(filter)
		assert.Error(t, err)
	})
}

func TestRepoDashboard_GetDashboardTotalRepositories(t *testing.T) {
	t.Run("Should return TotalRepositories without error", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, nil, 1))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &dashboard.FilterDashboard{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		_, err := repo.GetDashboardTotalRepositories(filter)
		assert.NoError(t, err)
	})
	t.Run("Should return TotalRepositories without error if error is ErrorNotFoundRecords", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, enums.ErrorNotFoundRecords, 0))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &dashboard.FilterDashboard{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		_, err := repo.GetDashboardTotalRepositories(filter)
		assert.NoError(t, err)
	})
	t.Run("Should return TotalRepositories with error", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, errors.New("unknown error"), 0))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &dashboard.FilterDashboard{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		_, err := repo.GetDashboardTotalRepositories(filter)
		assert.Error(t, err)
	})
}
