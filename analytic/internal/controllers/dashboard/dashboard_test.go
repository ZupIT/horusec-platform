package dashboard

//func TestControllerRead_GetAllCharts(t *testing.T) {
//	t.Run("Should return all charts without errors", func(t *testing.T) {
//		repoMock := &repository.Mock{}
//		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
//		repoMock.On("GetDashboardTotalRepositories").Return(0, nil)
//		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
//		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
//		repoMock.On("GetDashboardVulnByRepository").Return([]*dashboard.VulnerabilitiesByRepository{}, nil)
//		repoMock.On("GetDashboardVulnByLanguage").Return([]*dashboard.VulnerabilitiesByLanguage{}, nil)
//		repoMock.On("GetDashboardVulnByTime").Return([]*dashboard.VulnerabilitiesByTime{}, nil)
//		response, err := NewControllerDashboardRead(repoMock).GetAllDashboardCharts(&dashboard.FilterDashboard{})
//		assert.NoError(t, err)
//		assert.NotEmpty(t, response)
//	})
//	t.Run("Should return error when GetDashboardTotalDevelopers", func(t *testing.T) {
//		repoMock := &repository.Mock{}
//		repoMock.On("GetDashboardTotalDevelopers").Return(0, errors.New("unexpected error"))
//		repoMock.On("GetDashboardTotalRepositories").Return(0, nil)
//		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
//		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
//		repoMock.On("GetDashboardVulnByRepository").Return([]*dashboard.VulnerabilitiesByRepository{}, nil)
//		repoMock.On("GetDashboardVulnByLanguage").Return([]*dashboard.VulnerabilitiesByLanguage{}, nil)
//		repoMock.On("GetDashboardVulnByTime").Return([]*dashboard.VulnerabilitiesByTime{}, nil)
//		response, err := NewControllerDashboardRead(repoMock).GetAllDashboardCharts(&dashboard.FilterDashboard{})
//		assert.Error(t, err)
//		assert.Empty(t, response)
//	})
//	t.Run("Should return error when GetDashboardTotalRepositories", func(t *testing.T) {
//		repoMock := &repository.Mock{}
//		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
//		repoMock.On("GetDashboardTotalRepositories").Return(0, errors.New("unexpected error"))
//		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
//		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
//		repoMock.On("GetDashboardVulnByRepository").Return([]*dashboard.VulnerabilitiesByRepository{}, nil)
//		repoMock.On("GetDashboardVulnByLanguage").Return([]*dashboard.VulnerabilitiesByLanguage{}, nil)
//		repoMock.On("GetDashboardVulnByTime").Return([]*dashboard.VulnerabilitiesByTime{}, nil)
//		response, err := NewControllerDashboardRead(repoMock).GetAllDashboardCharts(&dashboard.FilterDashboard{})
//		assert.Error(t, err)
//		assert.Empty(t, response)
//	})
//	t.Run("Should return error when GetDashboardVulnBySeverity", func(t *testing.T) {
//		repoMock := &repository.Mock{}
//		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
//		repoMock.On("GetDashboardTotalRepositories").Return(0, nil)
//		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, errors.New("unexpected error"))
//		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
//		repoMock.On("GetDashboardVulnByRepository").Return([]*dashboard.VulnerabilitiesByRepository{}, nil)
//		repoMock.On("GetDashboardVulnByLanguage").Return([]*dashboard.VulnerabilitiesByLanguage{}, nil)
//		repoMock.On("GetDashboardVulnByTime").Return([]*dashboard.VulnerabilitiesByTime{}, nil)
//		response, err := NewControllerDashboardRead(repoMock).GetAllDashboardCharts(&dashboard.FilterDashboard{})
//		assert.Error(t, err)
//		assert.Empty(t, response)
//	})
//	t.Run("Should return error when GetDashboardVulnByAuthor", func(t *testing.T) {
//		repoMock := &repository.Mock{}
//		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
//		repoMock.On("GetDashboardTotalRepositories").Return(0, nil)
//		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
//		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, errors.New("unexpected error"))
//		repoMock.On("GetDashboardVulnByRepository").Return([]*dashboard.VulnerabilitiesByRepository{}, nil)
//		repoMock.On("GetDashboardVulnByLanguage").Return([]*dashboard.VulnerabilitiesByLanguage{}, nil)
//		repoMock.On("GetDashboardVulnByTime").Return([]*dashboard.VulnerabilitiesByTime{}, nil)
//		response, err := NewControllerDashboardRead(repoMock).GetAllDashboardCharts(&dashboard.FilterDashboard{})
//		assert.Error(t, err)
//		assert.Empty(t, response)
//	})
//	t.Run("Should return error when VulnerabilitiesByRepository", func(t *testing.T) {
//		repoMock := &repository.Mock{}
//		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
//		repoMock.On("GetDashboardTotalRepositories").Return(0, nil)
//		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
//		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
//		repoMock.On("GetDashboardVulnByRepository").Return([]*dashboard.VulnerabilitiesByRepository{}, errors.New("unexpected error"))
//		repoMock.On("GetDashboardVulnByLanguage").Return([]*dashboard.VulnerabilitiesByLanguage{}, nil)
//		repoMock.On("GetDashboardVulnByTime").Return([]*dashboard.VulnerabilitiesByTime{}, nil)
//		response, err := NewControllerDashboardRead(repoMock).GetAllDashboardCharts(&dashboard.FilterDashboard{})
//		assert.Error(t, err)
//		assert.Empty(t, response)
//	})
//	t.Run("Should return error when GetDashboardVulnByLanguage", func(t *testing.T) {
//		repoMock := &repository.Mock{}
//		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
//		repoMock.On("GetDashboardTotalRepositories").Return(0, nil)
//		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
//		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
//		repoMock.On("GetDashboardVulnByRepository").Return([]*dashboard.VulnerabilitiesByRepository{}, nil)
//		repoMock.On("GetDashboardVulnByLanguage").Return([]*dashboard.VulnerabilitiesByLanguage{}, errors.New("unexpected error"))
//		repoMock.On("GetDashboardVulnByTime").Return([]*dashboard.VulnerabilitiesByTime{}, nil)
//		response, err := NewControllerDashboardRead(repoMock).GetAllDashboardCharts(&dashboard.FilterDashboard{})
//		assert.Error(t, err)
//		assert.Empty(t, response)
//	})
//	t.Run("Should return error when GetDashboardVulnByTime", func(t *testing.T) {
//		repoMock := &repository.Mock{}
//		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
//		repoMock.On("GetDashboardTotalRepositories").Return(0, nil)
//		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
//		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
//		repoMock.On("GetDashboardVulnByRepository").Return([]*dashboard.VulnerabilitiesByRepository{}, nil)
//		repoMock.On("GetDashboardVulnByLanguage").Return([]*dashboard.VulnerabilitiesByLanguage{}, nil)
//		repoMock.On("GetDashboardVulnByTime").Return([]*dashboard.VulnerabilitiesByTime{}, errors.New("unexpected error"))
//		response, err := NewControllerDashboardRead(repoMock).GetAllDashboardCharts(&dashboard.FilterDashboard{})
//		assert.Error(t, err)
//		assert.Empty(t, response)
//	})
//}
//
//func TestControllerWrite_AddVulnerabilitiesByAuthor(t *testing.T) {
//	t.Run("Should AddVulnerabilitiesByAuthor with success", func(t *testing.T) {
//		repoMock := &repository.Mock{}
//		repoMock.On("Inactive").Return(nil)
//		repoMock.On("Save").Return(nil)
//		entity := &analysis.Analysis{
//			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
//				{Vulnerability: vulnerability.Vulnerability{}},
//			},
//		}
//		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByAuthor(entity)
//		assert.NoError(t, err)
//	})
//	t.Run("Should AddVulnerabilitiesByAuthor with error when inactive vulns", func(t *testing.T) {
//		repoMock := &repository.Mock{}
//		repoMock.On("Inactive").Return(errors.New("unexpected error"))
//		repoMock.On("Save").Return(nil)
//		entity := &analysis.Analysis{
//			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
//				{Vulnerability: vulnerability.Vulnerability{}},
//			},
//		}
//		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByAuthor(entity)
//		assert.Error(t, err)
//	})
//	t.Run("Should AddVulnerabilitiesByAuthor with error when save vulns", func(t *testing.T) {
//		repoMock := &repository.Mock{}
//		repoMock.On("Inactive").Return(nil)
//		repoMock.On("Save").Return(errors.New("unexpected error"))
//		entity := &analysis.Analysis{
//			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
//				{Vulnerability: vulnerability.Vulnerability{}},
//			},
//		}
//		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByAuthor(entity)
//		assert.Error(t, err)
//	})
//}
//
//func TestControllerWrite_AddVulnerabilitiesByLanguage(t *testing.T) {
//	t.Run("Should AddVulnerabilitiesByLanguage with success", func(t *testing.T) {
//		repoMock := &repository.Mock{}
//		repoMock.On("Inactive").Return(nil)
//		repoMock.On("Save").Return(nil)
//		entity := &analysis.Analysis{
//			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
//				{Vulnerability: vulnerability.Vulnerability{}},
//			},
//		}
//		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByLanguage(entity)
//		assert.NoError(t, err)
//	})
//	t.Run("Should AddVulnerabilitiesByLanguage with error when inactive vulns", func(t *testing.T) {
//		repoMock := &repository.Mock{}
//		repoMock.On("Inactive").Return(errors.New("unexpected error"))
//		repoMock.On("Save").Return(nil)
//		entity := &analysis.Analysis{
//			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
//				{Vulnerability: vulnerability.Vulnerability{}},
//			},
//		}
//		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByLanguage(entity)
//		assert.Error(t, err)
//	})
//	t.Run("Should AddVulnerabilitiesByLanguage with error when save vulns", func(t *testing.T) {
//		repoMock := &repository.Mock{}
//		repoMock.On("Inactive").Return(nil)
//		repoMock.On("Save").Return(errors.New("unexpected error"))
//		entity := &analysis.Analysis{
//			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
//				{Vulnerability: vulnerability.Vulnerability{}},
//			},
//		}
//		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByLanguage(entity)
//		assert.Error(t, err)
//	})
//}
//
//func TestControllerWrite_AddVulnerabilitiesByRepository(t *testing.T) {
//	t.Run("Should AddVulnerabilitiesByRepository with success", func(t *testing.T) {
//		repoMock := &repository.Mock{}
//		repoMock.On("Inactive").Return(nil)
//		repoMock.On("Save").Return(nil)
//		entity := &analysis.Analysis{
//			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
//				{Vulnerability: vulnerability.Vulnerability{}},
//			},
//		}
//		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByRepository(entity)
//		assert.NoError(t, err)
//	})
//	t.Run("Should AddVulnerabilitiesByRepository with error when inactive vulns", func(t *testing.T) {
//		repoMock := &repository.Mock{}
//		repoMock.On("Inactive").Return(errors.New("unexpected error"))
//		repoMock.On("Save").Return(nil)
//		entity := &analysis.Analysis{
//			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
//				{Vulnerability: vulnerability.Vulnerability{}},
//			},
//		}
//		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByRepository(entity)
//		assert.Error(t, err)
//	})
//	t.Run("Should AddVulnerabilitiesByRepository with error when save vulns", func(t *testing.T) {
//		repoMock := &repository.Mock{}
//		repoMock.On("Inactive").Return(nil)
//		repoMock.On("Save").Return(errors.New("unexpected error"))
//		entity := &analysis.Analysis{
//			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
//				{Vulnerability: vulnerability.Vulnerability{}},
//			},
//		}
//		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByRepository(entity)
//		assert.Error(t, err)
//	})
//}
//
//func TestControllerWrite_AddVulnerabilitiesByTime(t *testing.T) {
//	t.Run("Should AddVulnerabilitiesByTime with success", func(t *testing.T) {
//		repoMock := &repository.Mock{}
//		repoMock.On("Inactive").Return(nil)
//		repoMock.On("Save").Return(nil)
//		entity := &analysis.Analysis{
//			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
//				{Vulnerability: vulnerability.Vulnerability{}},
//			},
//		}
//		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByTime(entity)
//		assert.NoError(t, err)
//	})
//	t.Run("Should AddVulnerabilitiesByTime with error when save vulns", func(t *testing.T) {
//		repoMock := &repository.Mock{}
//		repoMock.On("Inactive").Return(nil)
//		repoMock.On("Save").Return(errors.New("unexpected error"))
//		entity := &analysis.Analysis{
//			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
//				{Vulnerability: vulnerability.Vulnerability{}},
//			},
//		}
//		err := NewControllerDashboardWrite(repoMock).AddVulnerabilitiesByTime(entity)
//		assert.Error(t, err)
//	})
//}
