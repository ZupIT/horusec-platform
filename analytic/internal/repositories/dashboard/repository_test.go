package dashboard

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
)

func TestGetDashboardTotalDevelopers(t *testing.T) {
	t.Run("should return total developers without error", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(response.NewResponse(0, nil, 1))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewRepoDashboard(connection)

		total, err := repository.GetDashboardTotalDevelopers(&dashboard.Filter{})
		assert.NoError(t, err)
		assert.NotNil(t, total)
	})
}

func TestGetDashboardVulnBySeverity(t *testing.T) {
	t.Run("should return get vulns by severity without errors", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(
			response.NewResponse(0, nil, &dashboard.Vulnerability{}))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewRepoDashboard(connection)

		result, err := repository.GetDashboardVulnBySeverity(&dashboard.Filter{})
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestGetDashboardVulnByAuthor(t *testing.T) {
	t.Run("should return get vulns by author without errors", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(
			response.NewResponse(0, nil, &[]*dashboard.VulnerabilitiesByAuthor{}))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewRepoDashboard(connection)

		_, err := repository.GetDashboardVulnByAuthor(&dashboard.Filter{})
		assert.NoError(t, err)
	})
}

func TestGetDashboardVulnByLanguage(t *testing.T) {
	t.Run("should return get vulns by language without errors", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(
			response.NewResponse(0, nil, []*dashboard.VulnerabilitiesByLanguage{}))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewRepoDashboard(connection)

		_, err := repository.GetDashboardVulnByLanguage(&dashboard.Filter{})
		assert.NoError(t, err)
	})
}

func TestGetDashboardVulnByTime(t *testing.T) {
	t.Run("should return get vulns by time without errors", func(t *testing.T) {
		databaseReadMock := &database.Mock{}
		databaseReadMock.On("Raw").Return(
			response.NewResponse(0, nil, []*dashboard.VulnerabilitiesByTime{}))

		connection := &database.Connection{
			Read:  databaseReadMock,
			Write: &database.Mock{},
		}

		repository := NewRepoDashboard(connection)

		_, err := repository.GetDashboardVulnByTime(&dashboard.Filter{})
		assert.NoError(t, err)
	})
}
