package dashboard

import (
	"errors"
	"testing"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"
	"github.com/stretchr/testify/assert"

	repoDashboard "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
)

func TestControllerWrite_AddVulnerabilitiesByAuthor(t *testing.T) {
	t.Run("Should AddVulnerabilitiesByAuthor with success", func(t *testing.T) {
		repoMock := &repoDashboard.Mock{}
		repoMock.On("Inactive").Return(nil)
		repoMock.On("Save").Return(nil)
		entity := &analysis.Analysis{
			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
				{Vulnerability: vulnerability.Vulnerability{}},
			},
		}
		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByAuthor(entity)
		assert.NoError(t, err)
	})
	t.Run("Should AddVulnerabilitiesByAuthor with error when inactive vulns", func(t *testing.T) {
		repoMock := &repoDashboard.Mock{}
		repoMock.On("Inactive").Return(errors.New("unexpected error"))
		repoMock.On("Save").Return(nil)
		entity := &analysis.Analysis{
			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
				{Vulnerability: vulnerability.Vulnerability{}},
			},
		}
		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByAuthor(entity)
		assert.Error(t, err)
	})
	t.Run("Should AddVulnerabilitiesByAuthor with error when save vulns", func(t *testing.T) {
		repoMock := &repoDashboard.Mock{}
		repoMock.On("Inactive").Return(nil)
		repoMock.On("Save").Return(errors.New("unexpected error"))
		entity := &analysis.Analysis{
			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
				{Vulnerability: vulnerability.Vulnerability{}},
			},
		}
		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByAuthor(entity)
		assert.Error(t, err)
	})
}

func TestControllerWrite_AddVulnerabilitiesByLanguage(t *testing.T) {
	t.Run("Should AddVulnerabilitiesByLanguage with success", func(t *testing.T) {
		repoMock := &repoDashboard.Mock{}
		repoMock.On("Inactive").Return(nil)
		repoMock.On("Save").Return(nil)
		entity := &analysis.Analysis{
			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
				{Vulnerability: vulnerability.Vulnerability{}},
			},
		}
		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByLanguage(entity)
		assert.NoError(t, err)
	})
	t.Run("Should AddVulnerabilitiesByLanguage with error when inactive vulns", func(t *testing.T) {
		repoMock := &repoDashboard.Mock{}
		repoMock.On("Inactive").Return(errors.New("unexpected error"))
		repoMock.On("Save").Return(nil)
		entity := &analysis.Analysis{
			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
				{Vulnerability: vulnerability.Vulnerability{}},
			},
		}
		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByLanguage(entity)
		assert.Error(t, err)
	})
	t.Run("Should AddVulnerabilitiesByLanguage with error when save vulns", func(t *testing.T) {
		repoMock := &repoDashboard.Mock{}
		repoMock.On("Inactive").Return(nil)
		repoMock.On("Save").Return(errors.New("unexpected error"))
		entity := &analysis.Analysis{
			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
				{Vulnerability: vulnerability.Vulnerability{}},
			},
		}
		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByLanguage(entity)
		assert.Error(t, err)
	})
}

func TestControllerWrite_AddVulnerabilitiesByRepository(t *testing.T) {
	t.Run("Should AddVulnerabilitiesByRepository with success", func(t *testing.T) {
		repoMock := &repoDashboard.Mock{}
		repoMock.On("Inactive").Return(nil)
		repoMock.On("Save").Return(nil)
		entity := &analysis.Analysis{
			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
				{Vulnerability: vulnerability.Vulnerability{}},
			},
		}
		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByRepository(entity)
		assert.NoError(t, err)
	})
	t.Run("Should AddVulnerabilitiesByRepository with error when inactive vulns", func(t *testing.T) {
		repoMock := &repoDashboard.Mock{}
		repoMock.On("Inactive").Return(errors.New("unexpected error"))
		repoMock.On("Save").Return(nil)
		entity := &analysis.Analysis{
			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
				{Vulnerability: vulnerability.Vulnerability{}},
			},
		}
		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByRepository(entity)
		assert.Error(t, err)
	})
	t.Run("Should AddVulnerabilitiesByRepository with error when save vulns", func(t *testing.T) {
		repoMock := &repoDashboard.Mock{}
		repoMock.On("Inactive").Return(nil)
		repoMock.On("Save").Return(errors.New("unexpected error"))
		entity := &analysis.Analysis{
			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
				{Vulnerability: vulnerability.Vulnerability{}},
			},
		}
		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByRepository(entity)
		assert.Error(t, err)
	})
}

func TestControllerWrite_AddVulnerabilitiesByTime(t *testing.T) {
	t.Run("Should AddVulnerabilitiesByTime with success", func(t *testing.T) {
		repoMock := &repoDashboard.Mock{}
		repoMock.On("Inactive").Return(nil)
		repoMock.On("Save").Return(nil)
		entity := &analysis.Analysis{
			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
				{Vulnerability: vulnerability.Vulnerability{}},
			},
		}
		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByTime(entity)
		assert.NoError(t, err)
	})
	t.Run("Should AddVulnerabilitiesByTime with error when save vulns", func(t *testing.T) {
		repoMock := &repoDashboard.Mock{}
		repoMock.On("Inactive").Return(nil)
		repoMock.On("Save").Return(errors.New("unexpected error"))
		entity := &analysis.Analysis{
			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
				{Vulnerability: vulnerability.Vulnerability{}},
			},
		}
		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByTime(entity)
		assert.Error(t, err)
	})
}
