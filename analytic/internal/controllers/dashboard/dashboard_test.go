// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dashboard

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	analysisEntities "github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
	dashboardRepository "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
	dashboardUseCases "github.com/ZupIT/horusec-platform/analytic/internal/usecases/dashboard"
)

func TestGetAllDashboardChartsWorkspace(t *testing.T) {
	t.Run("should return all charts without errors", func(t *testing.T) {
		repoMock := &dashboardRepository.Mock{}

		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
		repoMock.On("GetDashboardTotalRepositories").Return(0, nil)
		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
		repoMock.On("GetDashboardVulnByRepository").Return([]*dashboard.VulnerabilitiesByRepository{}, nil)
		repoMock.On("GetDashboardVulnByLanguage").Return([]*dashboard.VulnerabilitiesByLanguage{}, nil)
		repoMock.On("GetDashboardVulnByTime").Return([]*dashboard.VulnerabilitiesByTime{}, nil)

		controller := NewDashboardController(repoMock, repoMock, &database.Connection{},
			dashboardUseCases.NewUseCaseDashboard())

		result, err := controller.GetAllDashboardChartsWorkspace(&dashboard.Filter{})
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("should return error when getting vuln by time", func(t *testing.T) {
		repoMock := &dashboardRepository.Mock{}

		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
		repoMock.On("GetDashboardTotalRepositories").Return(0, nil)
		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
		repoMock.On("GetDashboardVulnByRepository").Return([]*dashboard.VulnerabilitiesByRepository{}, nil)
		repoMock.On("GetDashboardVulnByLanguage").Return([]*dashboard.VulnerabilitiesByLanguage{}, nil)
		repoMock.On("GetDashboardVulnByTime").Return(
			[]*dashboard.VulnerabilitiesByTime{}, errors.New("test"))

		controller := NewDashboardController(repoMock, repoMock, &database.Connection{},
			dashboardUseCases.NewUseCaseDashboard())

		result, err := controller.GetAllDashboardChartsWorkspace(&dashboard.Filter{})
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when getting vuln by language", func(t *testing.T) {
		repoMock := &dashboardRepository.Mock{}

		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
		repoMock.On("GetDashboardTotalRepositories").Return(0, nil)
		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
		repoMock.On("GetDashboardVulnByRepository").Return([]*dashboard.VulnerabilitiesByRepository{}, nil)
		repoMock.On("GetDashboardVulnByLanguage").Return(
			[]*dashboard.VulnerabilitiesByLanguage{}, errors.New("test"))

		controller := NewDashboardController(repoMock, repoMock, &database.Connection{},
			dashboardUseCases.NewUseCaseDashboard())

		result, err := controller.GetAllDashboardChartsWorkspace(&dashboard.Filter{})
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when getting vuln by repository", func(t *testing.T) {
		repoMock := &dashboardRepository.Mock{}

		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
		repoMock.On("GetDashboardTotalRepositories").Return(0, nil)
		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
		repoMock.On("GetDashboardVulnByRepository").Return(
			[]*dashboard.VulnerabilitiesByRepository{}, errors.New("test"))

		controller := NewDashboardController(repoMock, repoMock, &database.Connection{},
			dashboardUseCases.NewUseCaseDashboard())

		result, err := controller.GetAllDashboardChartsWorkspace(&dashboard.Filter{})
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when getting vuln by author", func(t *testing.T) {
		repoMock := &dashboardRepository.Mock{}

		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
		repoMock.On("GetDashboardTotalRepositories").Return(0, nil)
		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
		repoMock.On("GetDashboardVulnByAuthor").Return(
			[]*dashboard.VulnerabilitiesByAuthor{}, errors.New("test"))

		controller := NewDashboardController(repoMock, repoMock, &database.Connection{},
			dashboardUseCases.NewUseCaseDashboard())

		result, err := controller.GetAllDashboardChartsWorkspace(&dashboard.Filter{})
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when getting vuln by severity", func(t *testing.T) {
		repoMock := &dashboardRepository.Mock{}

		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
		repoMock.On("GetDashboardTotalRepositories").Return(0, nil)
		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, errors.New("test"))

		controller := NewDashboardController(repoMock, repoMock, &database.Connection{},
			dashboardUseCases.NewUseCaseDashboard())

		result, err := controller.GetAllDashboardChartsWorkspace(&dashboard.Filter{})
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when getting total repositories", func(t *testing.T) {
		repoMock := &dashboardRepository.Mock{}

		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
		repoMock.On("GetDashboardTotalRepositories").Return(0, errors.New("test"))

		controller := NewDashboardController(repoMock, repoMock, &database.Connection{},
			dashboardUseCases.NewUseCaseDashboard())

		result, err := controller.GetAllDashboardChartsWorkspace(&dashboard.Filter{})
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when getting total developers", func(t *testing.T) {
		repoMock := &dashboardRepository.Mock{}

		repoMock.On("GetDashboardTotalDevelopers").Return(0, errors.New("test"))

		controller := NewDashboardController(repoMock, repoMock, &database.Connection{},
			dashboardUseCases.NewUseCaseDashboard())

		result, err := controller.GetAllDashboardChartsWorkspace(&dashboard.Filter{})
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestAddVulnerabilitiesByAuthor(t *testing.T) {
	t.Run("should success add vulnerabilities", func(t *testing.T) {
		repoMock := &dashboardRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})

		controller := NewDashboardController(repoMock, repoMock,
			&database.Connection{Write: databaseMock, Read: databaseMock}, dashboardUseCases.NewUseCaseDashboard())

		assert.NoError(t, controller.AddVulnerabilitiesByAuthor(&analysisEntities.Analysis{}))
	})
}

func TestAddVulnerabilitiesByLanguage(t *testing.T) {
	t.Run("should success add vulnerabilities", func(t *testing.T) {
		repoMock := &dashboardRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})

		controller := NewDashboardController(repoMock, repoMock,
			&database.Connection{Write: databaseMock, Read: databaseMock}, dashboardUseCases.NewUseCaseDashboard())

		assert.NoError(t, controller.AddVulnerabilitiesByLanguage(&analysisEntities.Analysis{}))
	})
}

func TestAddVulnerabilitiesByRepository(t *testing.T) {
	t.Run("should success add vulnerabilities", func(t *testing.T) {
		repoMock := &dashboardRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})

		controller := NewDashboardController(repoMock, repoMock,
			&database.Connection{Write: databaseMock, Read: databaseMock}, dashboardUseCases.NewUseCaseDashboard())

		assert.NoError(t, controller.AddVulnerabilitiesByRepository(&analysisEntities.Analysis{}))
	})
}

func TestAddVulnerabilitiesByTime(t *testing.T) {
	t.Run("should success add vulnerabilities", func(t *testing.T) {
		repoMock := &dashboardRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})

		controller := NewDashboardController(repoMock, repoMock,
			&database.Connection{Write: databaseMock, Read: databaseMock}, dashboardUseCases.NewUseCaseDashboard())

		assert.NoError(t, controller.AddVulnerabilitiesByTime(&analysisEntities.Analysis{}))
	})
}

func TestGetAllDashboardChartsRepository(t *testing.T) {
	t.Run("should return all charts without errors", func(t *testing.T) {
		repoMock := &dashboardRepository.Mock{}

		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
		repoMock.On("GetDashboardVulnByLanguage").Return([]*dashboard.VulnerabilitiesByLanguage{}, nil)
		repoMock.On("GetDashboardVulnByTime").Return([]*dashboard.VulnerabilitiesByTime{}, nil)

		controller := NewDashboardController(repoMock, repoMock, &database.Connection{},
			dashboardUseCases.NewUseCaseDashboard())

		result, err := controller.GetAllDashboardChartsRepository(&dashboard.Filter{})
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("should return error when getting vuln by time", func(t *testing.T) {
		repoMock := &dashboardRepository.Mock{}

		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
		repoMock.On("GetDashboardVulnByLanguage").Return([]*dashboard.VulnerabilitiesByLanguage{}, nil)
		repoMock.On("GetDashboardVulnByTime").Return(
			[]*dashboard.VulnerabilitiesByTime{}, errors.New("test"))

		controller := NewDashboardController(repoMock, repoMock, &database.Connection{},
			dashboardUseCases.NewUseCaseDashboard())

		result, err := controller.GetAllDashboardChartsRepository(&dashboard.Filter{})
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when getting vuln by language", func(t *testing.T) {
		repoMock := &dashboardRepository.Mock{}

		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
		repoMock.On("GetDashboardVulnByAuthor").Return([]*dashboard.VulnerabilitiesByAuthor{}, nil)
		repoMock.On("GetDashboardVulnByLanguage").Return(
			[]*dashboard.VulnerabilitiesByLanguage{}, errors.New("test"))

		controller := NewDashboardController(repoMock, repoMock, &database.Connection{},
			dashboardUseCases.NewUseCaseDashboard())

		result, err := controller.GetAllDashboardChartsRepository(&dashboard.Filter{})
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when getting vuln by author", func(t *testing.T) {
		repoMock := &dashboardRepository.Mock{}

		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, nil)
		repoMock.On("GetDashboardVulnByAuthor").Return(
			[]*dashboard.VulnerabilitiesByAuthor{}, errors.New("test"))

		controller := NewDashboardController(repoMock, repoMock, &database.Connection{},
			dashboardUseCases.NewUseCaseDashboard())

		result, err := controller.GetAllDashboardChartsRepository(&dashboard.Filter{})
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when getting vuln by severity", func(t *testing.T) {
		repoMock := &dashboardRepository.Mock{}

		repoMock.On("GetDashboardTotalDevelopers").Return(0, nil)
		repoMock.On("GetDashboardVulnBySeverity").Return(&dashboard.Vulnerability{}, errors.New("test"))

		controller := NewDashboardController(repoMock, repoMock, &database.Connection{},
			dashboardUseCases.NewUseCaseDashboard())

		result, err := controller.GetAllDashboardChartsRepository(&dashboard.Filter{})
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when getting total developers", func(t *testing.T) {
		repoMock := &dashboardRepository.Mock{}

		repoMock.On("GetDashboardTotalDevelopers").Return(0, errors.New("test"))

		controller := NewDashboardController(repoMock, repoMock, &database.Connection{},
			dashboardUseCases.NewUseCaseDashboard())

		result, err := controller.GetAllDashboardChartsRepository(&dashboard.Filter{})
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
