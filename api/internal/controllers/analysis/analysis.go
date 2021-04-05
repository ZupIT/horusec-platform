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

import "C"
import (
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/entities/cli"
	"github.com/ZupIT/horusec-devkit/pkg/enums/queues"
	appConfiguration "github.com/ZupIT/horusec-devkit/pkg/services/app"
	brokerService "github.com/ZupIT/horusec-devkit/pkg/services/broker"
	brokerConfiguration "github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
)

type IController interface {
	GetAnalysis(analysisID uuid.UUID) (*analysis.Analysis, error)
	SaveAnalysis(analysisData *cli.AnalysisData) (uuid.UUID, error)
}

type Controller struct {
	broker        brokerService.IBroker
	databaseRead  database.IDatabaseRead
	databaseWrite database.IDatabaseWrite
	brokerConfig  brokerConfiguration.IConfig
	appConfig     appConfiguration.IConfig
}

func NewAnalysisController(broker brokerService.IBroker, brokerConfig brokerConfiguration.IConfig,
	databaseConnection *database.Connection, appConfig appConfiguration.IConfig) IController {
	return &Controller{
		appConfig:     appConfig,
		broker:        broker,
		brokerConfig:  brokerConfig,
		databaseRead:  databaseConnection.Read,
		databaseWrite: databaseConnection.Write,
	}
}

func (c *Controller) GetAnalysis(analysisID uuid.UUID) (*analysis.Analysis, error) {
	response := c.databaseRead.Find(
		&analysis.Analysis{},
		map[string]interface{}{"analysis_id": analysisID},
		(&analysis.Analysis{}).GetTable())

	if response.GetError() != nil {
		return nil, response.GetError()
	}
	if response.GetData() == nil {
		return nil, enums.ErrorNotFoundRecords
	}

	return response.GetData().(*analysis.Analysis), nil
}

func (c *Controller) SaveAnalysis(analysisData *cli.AnalysisData) (uuid.UUID, error) {
	analysisDecorated, err := c.decoratorAnalysisToSave(analysisData)
	if err != nil {
		return uuid.Nil, err
	}
	if err := c.createNewAnalysis(analysisDecorated); err != nil {
		return uuid.Nil, err
	}
	return analysisDecorated.ID, c.publishToWebhookAnalysis(analysisDecorated)
}

func (c *Controller) decoratorAnalysisToSave(data *cli.AnalysisData) (*analysis.Analysis, error) {
	newAnalysis, err := c.extractBaseOfTheAnalysis(data)
	if err != nil {
		return nil, err
	}
	for keyObservable := range data.Analysis.AnalysisVulnerabilities {
		observable := data.Analysis.AnalysisVulnerabilities[keyObservable]
		if !c.hasDuplicatedHash(newAnalysis, &observable) {
			newAnalysis.AnalysisVulnerabilities = append(newAnalysis.AnalysisVulnerabilities, observable)
		}
	}
	return newAnalysis, nil
}

func (c *Controller) publishToWebhookAnalysis(analysisData *analysis.Analysis) error {
	if !c.appConfig.IsBrokerDisabled() {
		return c.broker.Publish(queues.HorusecWebhookDispatch.ToString(), "", "", analysisData.ToBytes())
	}
	return nil
}

func (c *Controller) extractBaseOfTheAnalysis(data *cli.AnalysisData) (*analysis.Analysis, error) {
	workspace, repo, err := c.getWorkspaceAndRepository(data.Analysis.CompanyID, data.Analysis.RepositoryID, data.RepositoryName)
	if err != nil {
		return nil, err
	}
	return &analysis.Analysis{
		ID:             uuid.New(),
		RepositoryID:   repo["repositoryID"].(uuid.UUID),
		RepositoryName: repo["name"].(string),
		CompanyID:      workspace["workspaceID"].(uuid.UUID),
		CompanyName:    workspace["name"].(string),
		Status:         data.Analysis.Status,
		Errors:         data.Analysis.Errors,
		CreatedAt:      data.Analysis.CreatedAt,
		FinishedAt:     data.Analysis.FinishedAt,
	}, nil
}

// TODO: remove map and get data from database
func (c *Controller) getWorkspaceAndRepository(
	workspaceID uuid.UUID, repositoryID uuid.UUID, repositoryName string) (
	map[string]interface{}, map[string]interface{}, error) {
	repo, err := c.getRepositoryByRepositoryIDOrName(repositoryID, repositoryName)
	if err != nil {
		return map[string]interface{}{}, map[string]interface{}{}, err
	}
	workspace, err := c.getWorkspaceByID(workspaceID)
	if err != nil {
		return map[string]interface{}{}, map[string]interface{}{}, err
	}
	return repo, workspace, nil
}

// TODO: remove map and get data from database
func (c *Controller) getRepositoryByRepositoryIDOrName(repositoryID uuid.UUID, repositoryName string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"repositoryID": repositoryID,
		"name":         repositoryName,
	}, nil
}

// TODO: remove map and get data from database
func (c *Controller) getWorkspaceByID(workspace uuid.UUID) (map[string]interface{}, error) {
	return map[string]interface{}{
		"workspaceID": workspace,
		"name":        uuid.New().String(),
	}, nil
}

func (c *Controller) hasDuplicatedHash(newAnalysis *analysis.Analysis, observable *analysis.RelationshipAnalysisVuln) bool {
	for keyCurrent := range newAnalysis.AnalysisVulnerabilities {
		current := newAnalysis.AnalysisVulnerabilities[keyCurrent]
		if observable.Vulnerability.VulnHash == current.Vulnerability.VulnHash {
			return true
		}
	}
	return false
}

func (c *Controller) createNewAnalysis(newAnalysis *analysis.Analysis) error {
	return c.databaseWrite.Create(newAnalysis, newAnalysis.GetTable()).GetError()
}
