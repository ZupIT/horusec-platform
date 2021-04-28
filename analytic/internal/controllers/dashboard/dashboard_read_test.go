package dashboard

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
	repoDashboard "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
)

func TestControllerRead_GetAllCharts(t *testing.T) {
	t.Run("Should return all charts without errors", func(t *testing.T) {
		repoMock := &repoDashboard.Mock{}
		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
		repoMock.On("GetDashboardTotalRepositories").Return(0, nil)
		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
		repoMock.On("GetDashboardVulnByRepository").Return([]*dashboard.VulnerabilitiesByRepository{}, nil)
		repoMock.On("GetDashboardVulnByLanguage").Return([]*dashboard.VulnerabilitiesByLanguage{}, nil)
		repoMock.On("GetDashboardVulnByTime").Return([]*dashboard.VulnerabilitiesByTime{}, nil)
		response, err := NewControllerDashboardRead(repoMock).GetAllCharts(&dashboard.FilterDashboard{})
		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})
	t.Run("Should return error when GetDashboardTotalDevelopers", func(t *testing.T) {
		repoMock := &repoDashboard.Mock{}
		repoMock.On("GetDashboardTotalDevelopers").Return(0, errors.New("unexpected error"))
		repoMock.On("GetDashboardTotalRepositories").Return(0, nil)
		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
		repoMock.On("GetDashboardVulnByRepository").Return([]*dashboard.VulnerabilitiesByRepository{}, nil)
		repoMock.On("GetDashboardVulnByLanguage").Return([]*dashboard.VulnerabilitiesByLanguage{}, nil)
		repoMock.On("GetDashboardVulnByTime").Return([]*dashboard.VulnerabilitiesByTime{}, nil)
		response, err := NewControllerDashboardRead(repoMock).GetAllCharts(&dashboard.FilterDashboard{})
		assert.Error(t, err)
		assert.Empty(t, response)
	})
	t.Run("Should return error when GetDashboardTotalRepositories", func(t *testing.T) {
		repoMock := &repoDashboard.Mock{}
		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
		repoMock.On("GetDashboardTotalRepositories").Return(0, errors.New("unexpected error"))
		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
		repoMock.On("GetDashboardVulnByRepository").Return([]*dashboard.VulnerabilitiesByRepository{}, nil)
		repoMock.On("GetDashboardVulnByLanguage").Return([]*dashboard.VulnerabilitiesByLanguage{}, nil)
		repoMock.On("GetDashboardVulnByTime").Return([]*dashboard.VulnerabilitiesByTime{}, nil)
		response, err := NewControllerDashboardRead(repoMock).GetAllCharts(&dashboard.FilterDashboard{})
		assert.Error(t, err)
		assert.Empty(t, response)
	})
	t.Run("Should return error when GetDashboardVulnBySeverity", func(t *testing.T) {
		repoMock := &repoDashboard.Mock{}
		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
		repoMock.On("GetDashboardTotalRepositories").Return(0, nil)
		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, errors.New("unexpected error"))
		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
		repoMock.On("GetDashboardVulnByRepository").Return([]*dashboard.VulnerabilitiesByRepository{}, nil)
		repoMock.On("GetDashboardVulnByLanguage").Return([]*dashboard.VulnerabilitiesByLanguage{}, nil)
		repoMock.On("GetDashboardVulnByTime").Return([]*dashboard.VulnerabilitiesByTime{}, nil)
		response, err := NewControllerDashboardRead(repoMock).GetAllCharts(&dashboard.FilterDashboard{})
		assert.Error(t, err)
		assert.Empty(t, response)
	})
	t.Run("Should return error when GetDashboardVulnByAuthor", func(t *testing.T) {
		repoMock := &repoDashboard.Mock{}
		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
		repoMock.On("GetDashboardTotalRepositories").Return(0, nil)
		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, errors.New("unexpected error"))
		repoMock.On("GetDashboardVulnByRepository").Return([]*dashboard.VulnerabilitiesByRepository{}, nil)
		repoMock.On("GetDashboardVulnByLanguage").Return([]*dashboard.VulnerabilitiesByLanguage{}, nil)
		repoMock.On("GetDashboardVulnByTime").Return([]*dashboard.VulnerabilitiesByTime{}, nil)
		response, err := NewControllerDashboardRead(repoMock).GetAllCharts(&dashboard.FilterDashboard{})
		assert.Error(t, err)
		assert.Empty(t, response)
	})
	t.Run("Should return error when VulnerabilitiesByRepository", func(t *testing.T) {
		repoMock := &repoDashboard.Mock{}
		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
		repoMock.On("GetDashboardTotalRepositories").Return(0, nil)
		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
		repoMock.On("GetDashboardVulnByRepository").Return([]*dashboard.VulnerabilitiesByRepository{}, errors.New("unexpected error"))
		repoMock.On("GetDashboardVulnByLanguage").Return([]*dashboard.VulnerabilitiesByLanguage{}, nil)
		repoMock.On("GetDashboardVulnByTime").Return([]*dashboard.VulnerabilitiesByTime{}, nil)
		response, err := NewControllerDashboardRead(repoMock).GetAllCharts(&dashboard.FilterDashboard{})
		assert.Error(t, err)
		assert.Empty(t, response)
	})
	t.Run("Should return error when GetDashboardVulnByLanguage", func(t *testing.T) {
		repoMock := &repoDashboard.Mock{}
		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
		repoMock.On("GetDashboardTotalRepositories").Return(0, nil)
		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
		repoMock.On("GetDashboardVulnByRepository").Return([]*dashboard.VulnerabilitiesByRepository{}, nil)
		repoMock.On("GetDashboardVulnByLanguage").Return([]*dashboard.VulnerabilitiesByLanguage{}, errors.New("unexpected error"))
		repoMock.On("GetDashboardVulnByTime").Return([]*dashboard.VulnerabilitiesByTime{}, nil)
		response, err := NewControllerDashboardRead(repoMock).GetAllCharts(&dashboard.FilterDashboard{})
		assert.Error(t, err)
		assert.Empty(t, response)
	})
	t.Run("Should return error when GetDashboardVulnByTime", func(t *testing.T) {
		repoMock := &repoDashboard.Mock{}
		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
		repoMock.On("GetDashboardTotalRepositories").Return(0, nil)
		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
		repoMock.On("GetDashboardVulnByRepository").Return([]*dashboard.VulnerabilitiesByRepository{}, nil)
		repoMock.On("GetDashboardVulnByLanguage").Return([]*dashboard.VulnerabilitiesByLanguage{}, nil)
		repoMock.On("GetDashboardVulnByTime").Return([]*dashboard.VulnerabilitiesByTime{}, errors.New("unexpected error"))
		response, err := NewControllerDashboardRead(repoMock).GetAllCharts(&dashboard.FilterDashboard{})
		assert.Error(t, err)
		assert.Empty(t, response)
	})
}
