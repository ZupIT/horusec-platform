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
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/entities/cli"
	brokerService "github.com/ZupIT/horusec-devkit/pkg/services/broker"
	brokerConfig "github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
)

type IController interface {
	SaveAnalysis(analysisData *cli.AnalysisData) (uuid.UUID, error)
	GetAnalysis(analysisID uuid.UUID) (*analysis.Analysis, error)
}

type Controller struct {
	broker        brokerService.IBroker
	databaseRead  database.IDatabaseRead
	databaseWrite database.IDatabaseWrite
	brokerConfig  brokerConfig.IConfig
}

func NewAnalysisController(broker brokerService.IBroker, brokerConfiguration brokerConfig.IConfig,
	databaseConnection *database.Connection) IController {
	return &Controller{
		broker:        broker,
		brokerConfig:  brokerConfiguration,
		databaseRead:  databaseConnection.Read,
		databaseWrite: databaseConnection.Write,
	}
}

func (c *Controller) SaveAnalysis(analysisData *cli.AnalysisData) (uuid.UUID, error) {
	panic("implement me")
}

func (c *Controller) GetAnalysis(analysisID uuid.UUID) (*analysis.Analysis, error) {
	panic("implement me")
}

// func (c *Controller) SaveAnalysis(analysisData *apiEntities.AnalysisData) (uuid.UUID, error) {
// 	company, err := c.repoCompany.GetByID(analysisData.Analysis.CompanyID)
// 	if err != nil {
// 		return uuid.Nil, err
// 	}
// 	repo, err := c.getRepositoryOrCreateIfNotExist(analysisData, company)
// 	if err != nil {
// 		return uuid.Nil, err
// 	}
// 	c.setDefaultContentToCreate(analysisData.Analysis, company, repo)
// 	analysis := c.removeAnalysisVulnerabilityWithHashDuplicate(analysisData.Analysis)
// 	return c.createAnalyzeAndVulnerabilities(analysis)
// }
//
// func (c *Controller) getRepositoryOrCreateIfNotExist(
// 	analysisData *apiEntities.AnalysisData, company *accountEntities.Company) (
// 	repo *accountEntities.Repository, err error) {
// 	if analysisData.RepositoryName != "" && analysisData.Analysis.RepositoryID == uuid.Nil {
// 		repo, err = c.repoRepository.GetByName(analysisData.Analysis.CompanyID, analysisData.RepositoryName)
// 		if err == errorsEnums.ErrNotFoundRecords {
// 			return c.createRepository(analysisData, company)
// 		}
// 		return repo, err
// 	}
//
// 	return c.repoRepository.Get(analysisData.Analysis.RepositoryID)
// }
//
// func (c *Controller) setDefaultContentToCreate(analysis *horusecEntities.Analysis,
// 	company *accountEntities.Company, repo *accountEntities.Repository) {
// 	analysis.
// 		SetCompanyName(company.Name).
// 		SetRepositoryName(repo.Name).
// 		SetRepositoryID(repo.RepositoryID).
// 		SetupIDInAnalysisContents()
// }
//
// func (c *Controller) createRepository(
// 	analysisData *apiEntities.AnalysisData, company *accountEntities.Company) (*accountEntities.Repository, error) {
// 	repo := &accountEntities.Repository{
// 		RepositoryID:    uuid.New(),
// 		CompanyID:       analysisData.Analysis.CompanyID,
// 		Name:            analysisData.RepositoryName,
// 		CreatedAt:       time.Now(),
// 		AuthzMember:     company.AuthzMember,
// 		AuthzAdmin:      company.AuthzAdmin,
// 		AuthzSupervisor: company.AuthzAdmin,
// 	}
//
// 	return repo, c.repoRepository.Create(repo, nil)
// }
//
// func (c *Controller) createAnalyzeAndVulnerabilities(analysis *horusecEntities.Analysis) (uuid.UUID, error) {
// 	conn := c.postgresWrite.StartTransaction()
// 	if err := c.repoAnalysis.Create(analysis, conn); err != nil {
// 		logger.LogError(
// 			"{HORUSEC_API} Error in rollback transaction analysis",
// 			conn.RollbackTransaction().GetError(),
// 		)
// 		return uuid.Nil, err
// 	}
// 	if err := conn.CommitTransaction().GetError(); err != nil {
// 		return uuid.Nil, err
// 	}
// 	return analysis.GetID(), c.publishToWebhook(analysis)
// }
//
// func (c *Controller) GetAnalysis(analysisID uuid.UUID) (*horusecEntities.Analysis, error) {
// 	return c.repoAnalysis.GetByID(analysisID)
// }
//
// func (c *Controller) removeAnalysisVulnerabilityWithHashDuplicate(
// 	analysis *horusecEntities.Analysis) *horusecEntities.Analysis {
// 	newAnalysis := analysis.GetAnalysisWithoutAnalysisVulnerabilities()
// 	for keyObservable := range analysis.AnalysisVulnerabilities {
// 		observable := analysis.AnalysisVulnerabilities[keyObservable]
// 		if !c.hasDuplicatedHash(newAnalysis, &observable) {
// 			newAnalysis.AnalysisVulnerabilities = append(newAnalysis.AnalysisVulnerabilities, observable)
// 		}
// 	}
// 	return newAnalysis
// }
//
// func (c *Controller) hasDuplicatedHash(
// 	newAnalysis *horusecEntities.Analysis, observable *horusecEntities.AnalysisVulnerabilities) bool {
// 	for keyCurrent := range newAnalysis.AnalysisVulnerabilities {
// 		current := newAnalysis.AnalysisVulnerabilities[keyCurrent]
// 		if observable.Vulnerability.VulnHash == current.Vulnerability.VulnHash {
// 			return true
// 		}
// 	}
// 	return false
// }
//
// func (c *Controller) publishToWebhook(analysis *horusecEntities.Analysis) error {
// 	if !c.config.IsDisabledBroker() {
// 		return c.broker.Publish(queues.HorusecWebhookDispatch.ToString(), "", "", analysis.ToBytes())
// 	}
// 	return nil
// }
