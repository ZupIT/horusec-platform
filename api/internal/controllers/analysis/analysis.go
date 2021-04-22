// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

package analysis

import (
	"github.com/ZupIT/horusec-devkit/pkg/enums/exchange"
	"time"

	"github.com/google/uuid"

	repoAnalysis "github.com/ZupIT/horusec-platform/api/internal/repositories/analysis"
	"github.com/ZupIT/horusec-platform/api/internal/repositories/repository"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/enums/queues"
	appConfiguration "github.com/ZupIT/horusec-devkit/pkg/services/app"
	brokerService "github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
)

type IController interface {
	GetAnalysis(analysisID uuid.UUID) (*analysis.Analysis, error)
	SaveAnalysis(analysisEntity *analysis.Analysis) (uuid.UUID, error)
}

type Controller struct {
	broker         brokerService.IBroker
	repoRepository repository.IRepository
	repoAnalysis   repoAnalysis.IAnalysis
	appConfig      appConfiguration.IConfig
}

func NewAnalysisController(broker brokerService.IBroker, appConfig appConfiguration.IConfig,
	repositoriesRepository repository.IRepository, repositoriesAnalysis repoAnalysis.IAnalysis) IController {
	return &Controller{
		repoRepository: repositoriesRepository,
		repoAnalysis:   repositoriesAnalysis,
		appConfig:      appConfig,
		broker:         broker,
	}
}

func (c *Controller) GetAnalysis(analysisID uuid.UUID) (*analysis.Analysis, error) {
	response := c.repoAnalysis.FindAnalysisByID(analysisID)
	if response.GetError() != nil {
		return nil, response.GetError()
	}
	if response.GetData() == nil {
		return nil, enums.ErrorNotFoundRecords
	}

	return response.GetData().(*analysis.Analysis), nil
}

func (c *Controller) SaveAnalysis(analysisEntity *analysis.Analysis) (uuid.UUID, error) {
	analysisEntity, err := c.createRepositoryIfNotExists(analysisEntity)
	if err != nil {
		return uuid.Nil, err
	}
	analysisDecorated, err := c.decorateAnalysisEntityAndSaveOnDatabase(analysisEntity)
	if err != nil {
		return uuid.Nil, err
	}
	if err := c.publishInBroker(analysisDecorated); err != nil {
		return uuid.Nil, err
	}
	return analysisDecorated.ID, nil
}

func (c *Controller) createRepositoryIfNotExists(analysisEntity *analysis.Analysis) (*analysis.Analysis, error) {
	if analysisEntity.RepositoryID == uuid.Nil {
		analysisEntity.RepositoryID = uuid.New()
		repositoryID, err := c.repoRepository.FindRepository(analysisEntity.WorkspaceID, analysisEntity.RepositoryName)
		if err != nil {
			if err == enums.ErrorNotFoundRecords {
				return analysisEntity, c.repoRepository.CreateRepository(analysisEntity.RepositoryID,
					analysisEntity.WorkspaceID, analysisEntity.RepositoryName)
			}
			return nil, err
		}
		analysisEntity.RepositoryID = repositoryID
	}
	return analysisEntity, nil
}

func (c *Controller) decorateAnalysisEntityAndSaveOnDatabase(
	analysisEntity *analysis.Analysis) (*analysis.Analysis, error) {
	analysisDecorated := c.decoratorAnalysisToSave(analysisEntity)
	return analysisDecorated, c.createNewAnalysis(analysisDecorated)
}

func (c *Controller) decoratorAnalysisToSave(analysisEntity *analysis.Analysis) *analysis.Analysis {
	newAnalysis := c.extractBaseOfTheAnalysis(analysisEntity)
	for keyObservable := range analysisEntity.AnalysisVulnerabilities {
		observable := analysisEntity.AnalysisVulnerabilities[keyObservable]
		if !c.hasDuplicatedHash(newAnalysis, &observable) {
			newAnalysis.AnalysisVulnerabilities = append(newAnalysis.AnalysisVulnerabilities,
				analysis.AnalysisVulnerabilities{
					VulnerabilityID: observable.Vulnerability.VulnerabilityID,
					AnalysisID:      newAnalysis.ID,
					CreatedAt:       time.Now(),
					Vulnerability:   observable.Vulnerability,
				})
		}
	}
	return newAnalysis
}

func (c *Controller) createNewAnalysis(newAnalysis *analysis.Analysis) error {
	return c.repoAnalysis.CreateFullAnalysis(newAnalysis)
}

func (c *Controller) extractBaseOfTheAnalysis(analysisEntity *analysis.Analysis) *analysis.Analysis {
	return &analysis.Analysis{
		ID:             uuid.New(),
		RepositoryID:   analysisEntity.RepositoryID,
		RepositoryName: analysisEntity.RepositoryName,
		WorkspaceID:    analysisEntity.WorkspaceID,
		WorkspaceName:  analysisEntity.WorkspaceName,
		Status:         analysisEntity.Status,
		Errors:         analysisEntity.Errors,
		CreatedAt:      analysisEntity.CreatedAt,
		FinishedAt:     analysisEntity.FinishedAt,
	}
}

func (c *Controller) hasDuplicatedHash(
	newAnalysis *analysis.Analysis, observable *analysis.AnalysisVulnerabilities) bool {
	for keyCurrent := range newAnalysis.AnalysisVulnerabilities {
		current := newAnalysis.AnalysisVulnerabilities[keyCurrent]
		if observable.Vulnerability.VulnHash == current.Vulnerability.VulnHash {
			return true
		}
	}
	return false
}

func (c *Controller) publishInBroker(analysisData *analysis.Analysis) error {
	if !c.appConfig.IsBrokerDisabled() {
		return c.broker.Publish(queues.HorusecNewAnalysis.ToString(), exchange.NewAnalysis.ToString(),
			exchange.Fanout.ToString(), analysisData.ToBytes())
	}
	return nil
}
