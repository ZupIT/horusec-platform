package dashboard

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
)

func TestGetDashboardTotalDevelopersWorkspace(t *testing.T) {
	t.Run("should return total developers without error", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(response.NewResponse(0, nil, 1))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewWorkspaceDashboard(connection)

		total, err := repository.GetDashboardTotalDevelopers(&dashboard.Filter{})
		assert.NoError(t, err)
		assert.NotNil(t, total)
	})
}

func TestGetDashboardVulnBySeverityWorkspace(t *testing.T) {
	t.Run("should return get vulns by severity without errors", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(
			response.NewResponse(0, nil, &dashboard.Vulnerability{}))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewWorkspaceDashboard(connection)

		result, err := repository.GetDashboardVulnBySeverity(&dashboard.Filter{})
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestGetDashboardVulnByAuthorWorkspace(t *testing.T) {
	t.Run("should return get vulns by author without errors", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(
			response.NewResponse(0, nil, &[]*dashboard.VulnerabilitiesByAuthor{}))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewWorkspaceDashboard(connection)

		_, err := repository.GetDashboardVulnByAuthor(&dashboard.Filter{})
		assert.NoError(t, err)
	})
}

func TestGetDashboardVulnByLanguageWorkspace(t *testing.T) {
	t.Run("should return get vulns by language without errors", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(
			response.NewResponse(0, nil, []*dashboard.VulnerabilitiesByLanguage{}))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewWorkspaceDashboard(connection)

		_, err := repository.GetDashboardVulnByLanguage(&dashboard.Filter{})
		assert.NoError(t, err)
	})
}

func TestGetDashboardVulnByTimeWorkspace(t *testing.T) {
	t.Run("should return get vulns by time without errors", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(
			response.NewResponse(0, nil, []*dashboard.VulnerabilitiesByTime{}))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewWorkspaceDashboard(connection)

		_, err := repository.GetDashboardVulnByTime(&dashboard.Filter{})
		assert.NoError(t, err)
	})
}

func TestGetDashboardTotalRepositoriesWorkspace(t *testing.T) {
	t.Run("should return total count of repositories", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(
			response.NewResponse(0, nil, 1))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewWorkspaceDashboard(connection)

		_, err := repository.GetDashboardTotalRepositories(&dashboard.Filter{})
		assert.NoError(t, err)
	})
}

func TestGetDashboardVulnByRepositoryWorkspace(t *testing.T) {
	t.Run("should return total count by repository", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(
			response.NewResponse(0, nil, []*dashboard.VulnerabilitiesByRepository{}))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewWorkspaceDashboard(connection)

		_, err := repository.GetDashboardVulnByRepository(&dashboard.Filter{})
		assert.NoError(t, err)
	})
}
