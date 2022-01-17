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
	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
	enumsdashboard "github.com/ZupIT/horusec-platform/analytic/internal/enums/dashboard"
	repositoriesdashboard "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
	usecasesdashboard "github.com/ZupIT/horusec-platform/analytic/internal/usecases/dashboard"
)

type IController interface {
	GetAllDashboardChartsWorkspace(filter *dashboard.Filter) (*dashboard.Response, error)
	GetAllDashboardChartsRepository(filter *dashboard.Filter) (*dashboard.Response, error)
	AddVulnerabilitiesByAuthor(entity *analysis.Analysis) error
	AddVulnerabilitiesByRepository(entity *analysis.Analysis) error
	AddVulnerabilitiesByLanguage(entity *analysis.Analysis) error
	AddVulnerabilitiesByTime(entity *analysis.Analysis) error
}

type Controller struct {
	repoRepository      repositoriesdashboard.IRepoRepository
	workspaceRepository repositoriesdashboard.IWorkspaceRepository
	useCases            usecasesdashboard.IUseCases
	databaseWrite       database.IDatabaseWrite
}

func NewDashboardController(repoRepository repositoriesdashboard.IRepoRepository,
	workspaceRepository repositoriesdashboard.IWorkspaceRepository, connection *database.Connection,
	useCases usecasesdashboard.IUseCases) IController {
	return &Controller{
		repoRepository:      repoRepository,
		workspaceRepository: workspaceRepository,
		databaseWrite:       connection.Write,
		useCases:            useCases,
	}
}

func (c *Controller) AddVulnerabilitiesByAuthor(entity *analysis.Analysis) error {
	return c.databaseWrite.Create(c.useCases.ParseAnalysisToVulnerabilitiesByAuthor(entity),
		enumsdashboard.TableVulnerabilitiesByAuthor).GetError()
}

func (c *Controller) AddVulnerabilitiesByRepository(entity *analysis.Analysis) error {
	return c.databaseWrite.Create(c.useCases.ParseAnalysisToVulnerabilitiesByRepository(entity),
		enumsdashboard.TableVulnerabilitiesByRepository).GetError()
}

func (c *Controller) AddVulnerabilitiesByLanguage(entity *analysis.Analysis) error {
	return c.databaseWrite.Create(c.useCases.ParseAnalysisToVulnerabilitiesByLanguage(entity),
		enumsdashboard.TableVulnerabilitiesByLanguage).GetError()
}

func (c *Controller) AddVulnerabilitiesByTime(entity *analysis.Analysis) error {
	return c.databaseWrite.Create(c.useCases.ParseAnalysisToVulnerabilitiesByTime(entity),
		enumsdashboard.TableVulnerabilitiesByTime).GetError()
}

func (c *Controller) GetAllDashboardChartsWorkspace(filter *dashboard.Filter) (*dashboard.Response, error) {
	response := &dashboard.Response{}

	if err := response.SetTotalAuthors(c.workspaceRepository.GetDashboardTotalDevelopers(filter)); err != nil {
		return nil, err
	}

	if err := response.SetTotalRepositories(c.workspaceRepository.GetDashboardTotalRepositories(filter)); err != nil {
		return nil, err
	}

	return c.getChartsBySeverityAndAuthorWorkspace(filter, response)
}

func (c *Controller) getChartsBySeverityAndAuthorWorkspace(filter *dashboard.Filter,
	response *dashboard.Response) (*dashboard.Response, error) {
	if err := response.SetChartBySeverity(c.workspaceRepository.GetDashboardVulnBySeverity(filter)); err != nil {
		return nil, err
	}

	if err := response.SetChartByAuthor(c.workspaceRepository.GetDashboardVulnByAuthor(filter)); err != nil {
		return nil, err
	}

	return c.getChartsByRepositoryAndLanguageWorkspace(filter, response)
}

func (c *Controller) getChartsByRepositoryAndLanguageWorkspace(filter *dashboard.Filter,
	response *dashboard.Response) (*dashboard.Response, error) {
	if err := response.SetChartByRepository(c.workspaceRepository.GetDashboardVulnByRepository(filter)); err != nil {
		return nil, err
	}

	if err := response.SetChartByLanguage(c.workspaceRepository.GetDashboardVulnByLanguage(filter)); err != nil {
		return nil, err
	}

	return c.getChartByTimeWorkspace(filter, response)
}

func (c *Controller) getChartByTimeWorkspace(filter *dashboard.Filter,
	response *dashboard.Response) (*dashboard.Response, error) {
	if err := response.SetChartByTime(c.workspaceRepository.GetDashboardVulnByTime(filter)); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Controller) GetAllDashboardChartsRepository(filter *dashboard.Filter) (*dashboard.Response, error) {
	response := &dashboard.Response{}

	if err := response.SetTotalAuthors(c.repoRepository.GetDashboardTotalDevelopers(filter)); err != nil {
		return nil, err
	}

	if err := response.SetChartBySeverity(c.repoRepository.GetDashboardVulnBySeverity(filter)); err != nil {
		return nil, err
	}

	return c.getChartsByLanguageAndAuthorRepository(filter, response)
}

func (c *Controller) getChartsByLanguageAndAuthorRepository(filter *dashboard.Filter,
	response *dashboard.Response) (*dashboard.Response, error) {
	if err := response.SetChartByAuthor(c.repoRepository.GetDashboardVulnByAuthor(filter)); err != nil {
		return nil, err
	}

	if err := response.SetChartByLanguage(c.repoRepository.GetDashboardVulnByLanguage(filter)); err != nil {
		return nil, err
	}

	return c.getChartByTimeRepository(filter, response)
}

func (c *Controller) getChartByTimeRepository(filter *dashboard.Filter,
	response *dashboard.Response) (*dashboard.Response, error) {
	if err := response.SetChartByTime(c.repoRepository.GetDashboardVulnByTime(filter)); err != nil {
		return nil, err
	}

	return response, nil
}
