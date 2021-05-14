package dashboard

import (
	"errors"
	"testing"
	"time"

	response2 "github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/response"

	database2 "github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/database"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var zeroValue = 0

func TestRepoDashboard_Save(t *testing.T) {
	t.Run("Should save new content without error", func(t *testing.T) {
		dbWrite := &database.Mock{}
		dbWrite.On("Create").Return(response.NewResponse(0, nil, nil))
		conn := &database.Connection{
			Read:  &database.Mock{},
			Write: dbWrite,
		}
		repo := NewRepoDashboard(conn)
		entity := &database2.VulnerabilitiesByAuthor{
			Author: "horusec@zup.com.br",
			Vulnerability: response2.Vulnerability{
				VulnerabilityID: uuid.New(),
				CreatedAt:       time.Now(),
				Active:          true,
				WorkspaceID:     uuid.New(),
				RepositoryID:    uuid.New(),
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
		entity := &database2.VulnerabilitiesByAuthor{
			Author: "horusec@zup.com.br",
			Vulnerability: response2.Vulnerability{
				VulnerabilityID: uuid.New(),
				CreatedAt:       time.Now(),
				Active:          true,
				WorkspaceID:     uuid.New(),
				RepositoryID:    uuid.New(),
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
			"active":        true,
			"repository_id": uuid.New(),
		}
		err := repo.Inactive(condition, (&database2.VulnerabilitiesByAuthor{}).GetTable())
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
			"active":        true,
			"repository_id": uuid.New(),
		}
		err := repo.Inactive(condition, (&database2.VulnerabilitiesByAuthor{}).GetTable())
		assert.Error(t, err)
	})
}

func TestRepoDashboard_GetDashboardTotalDevelopers(t *testing.T) {
	t.Run("Should return TotalDevelopers without error", func(t *testing.T) {
		dbRead := &database.Mock{}
		count := 1
		dbRead.On("Raw").Return(response.NewResponse(0, nil, &count))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		total, err := repo.GetDashboardTotalDevelopers(filter)
		assert.NoError(t, err)
		assert.Equal(t, 1, total)
	})
	t.Run("Should return TotalDevelopers without error if error is ErrorNotFoundRecords", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, enums.ErrorNotFoundRecords, &zeroValue))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		total, err := repo.GetDashboardTotalDevelopers(filter)
		assert.NoError(t, err)
		assert.Equal(t, 0, total)
	})
	t.Run("Should return TotalDevelopers with error", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, errors.New("unknown error"), &zeroValue))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		total, err := repo.GetDashboardTotalDevelopers(filter)
		assert.Error(t, err)
		assert.Equal(t, 0, total)
	})
}

func TestRepoDashboard_GetDashboardTotalRepositories(t *testing.T) {
	t.Run("Should return TotalRepositories without error", func(t *testing.T) {
		dbRead := &database.Mock{}
		count := 1
		dbRead.On("Raw").Return(response.NewResponse(0, nil, &count))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		total, err := repo.GetDashboardTotalRepositories(filter)
		assert.NoError(t, err)
		assert.Equal(t, 1, total)
	})
	t.Run("Should return TotalRepositories without error if error is ErrorNotFoundRecords", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, enums.ErrorNotFoundRecords, &zeroValue))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		total, err := repo.GetDashboardTotalRepositories(filter)
		assert.NoError(t, err)
		assert.Equal(t, 0, total)
	})
	t.Run("Should return TotalRepositories with error", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, errors.New("unknown error"), &zeroValue))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		total, err := repo.GetDashboardTotalRepositories(filter)
		assert.Error(t, err)
		assert.Equal(t, 0, total)
	})
}

func TestRepoDashboard_GetDashboardVulnBySeverity(t *testing.T) {
	t.Run("Should return GetDashboardVulnBySeverity with error", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, errors.New("unexpected error"), nil))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		res, err := repo.GetDashboardVulnBySeverity(filter)
		assert.Error(t, err)
		assert.NotNil(t, res)
	})
	t.Run("Should return GetDashboardVulnBySeverity without error when data is empty", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, nil, nil))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		res, err := repo.GetDashboardVulnBySeverity(filter)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
	t.Run("Should return GetDashboardVulnBySeverity without error", func(t *testing.T) {
		vulnID := uuid.New()
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, nil, &response2.Vulnerability{VulnerabilityID: vulnID}))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		res, err := repo.GetDashboardVulnBySeverity(filter)
		assert.NoError(t, err)
		assert.Equal(t, vulnID, res.VulnerabilityID)
	})
}
func TestRepoDashboard_GetDashboardVulnByAuthor(t *testing.T) {
	t.Run("Should return GetDashboardVulnByAuthor with error", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, errors.New("unexpected error"), nil))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		res, err := repo.GetDashboardVulnByAuthor(filter)
		assert.Error(t, err)
		assert.NotNil(t, res)
	})
	t.Run("Should return GetDashboardVulnByAuthor without error when data is empty", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, nil, &[]*database2.VulnerabilitiesByAuthor{}))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		res, err := repo.GetDashboardVulnByAuthor(filter)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
	t.Run("Should return GetDashboardVulnByAuthor without error", func(t *testing.T) {
		dbRead := &database.Mock{}
		vulnID := uuid.New()
		dbRead.On("Raw").Return(response.NewResponse(0, nil, &[]*database2.VulnerabilitiesByAuthor{{
			Vulnerability: response2.Vulnerability{VulnerabilityID: vulnID},
		}}))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		res, err := repo.GetDashboardVulnByAuthor(filter)
		assert.NoError(t, err)
		assert.Len(t, res, 1)
		assert.Equal(t, vulnID, res[0].VulnerabilityID)
	})
}
func TestRepoDashboard_GetDashboardVulnByRepository(t *testing.T) {
	t.Run("Should return GetDashboardVulnByRepository with error", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, errors.New("unexpected error"), nil))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		res, err := repo.GetDashboardVulnByRepository(filter)
		assert.Error(t, err)
		assert.NotNil(t, res)
	})
	t.Run("Should return GetDashboardVulnByRepository without error when data is empty", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, nil, nil))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		res, err := repo.GetDashboardVulnByRepository(filter)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
	t.Run("Should return GetDashboardVulnByRepository without error", func(t *testing.T) {
		dbRead := &database.Mock{}
		vulnID := uuid.New()
		dbRead.On("Raw").Return(response.NewResponse(0, nil, &[]*database2.VulnerabilitiesByRepository{{
			Vulnerability: response2.Vulnerability{VulnerabilityID: vulnID},
		}}))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		res, err := repo.GetDashboardVulnByRepository(filter)
		assert.NoError(t, err)
		assert.Len(t, res, 1)
		assert.Equal(t, vulnID, res[0].VulnerabilityID)
	})
}
func TestRepoDashboard_GetDashboardVulnByLanguage(t *testing.T) {
	t.Run("Should return GetDashboardVulnByLanguage with error", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, errors.New("unexpected error"), nil))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		res, err := repo.GetDashboardVulnByLanguage(filter)
		assert.Error(t, err)
		assert.NotNil(t, res)
	})
	t.Run("Should return GetDashboardVulnByLanguage without error when data is empty", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, nil, nil))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		res, err := repo.GetDashboardVulnByLanguage(filter)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
	t.Run("Should return GetDashboardVulnByLanguage without error", func(t *testing.T) {
		dbRead := &database.Mock{}
		vulnID := uuid.New()
		dbRead.On("Raw").Return(response.NewResponse(0, nil, &[]*database2.VulnerabilitiesByLanguage{{
			Vulnerability: response2.Vulnerability{VulnerabilityID: vulnID},
		}}))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		res, err := repo.GetDashboardVulnByLanguage(filter)
		assert.NoError(t, err)
		assert.Len(t, res, 1)
		assert.Equal(t, vulnID, res[0].VulnerabilityID)
	})
}
func TestRepoDashboard_GetDashboardVulnByTime(t *testing.T) {
	t.Run("Should return GetDashboardVulnByTime with error", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, errors.New("unexpected error"), nil))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		res, err := repo.GetDashboardVulnByTime(filter)
		assert.Error(t, err)
		assert.NotNil(t, res)
	})
	t.Run("Should return GetDashboardVulnByTime without error when data is empty", func(t *testing.T) {
		dbRead := &database.Mock{}
		dbRead.On("Raw").Return(response.NewResponse(0, nil, nil))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		res, err := repo.GetDashboardVulnByTime(filter)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
	t.Run("Should return GetDashboardVulnByTime without error", func(t *testing.T) {
		dbRead := &database.Mock{}
		vulnID := uuid.New()
		dbRead.On("Raw").Return(response.NewResponse(0, nil, &[]*database2.VulnerabilitiesByTime{{
			Vulnerability: response2.Vulnerability{VulnerabilityID: vulnID},
		}}))
		conn := &database.Connection{
			Read:  dbRead,
			Write: &database.Mock{},
		}
		repo := NewRepoDashboard(conn)
		filter := &database2.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
		}
		res, err := repo.GetDashboardVulnByTime(filter)
		assert.NoError(t, err)
		assert.Len(t, res, 1)
		assert.Equal(t, vulnID, res[0].VulnerabilityID)
	})
}
