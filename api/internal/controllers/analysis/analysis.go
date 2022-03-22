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

package analysis

import (
	"errors"
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"

	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/enums/exchange"
	appConfiguration "github.com/ZupIT/horusec-devkit/pkg/services/app"
	brokerService "github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"

	repoAnalysis "github.com/ZupIT/horusec-platform/api/internal/repositories/analysis"
	"github.com/ZupIT/horusec-platform/api/internal/repositories/repository"
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
	res := c.repoAnalysis.FindAnalysisByID(analysisID)
	if res.GetError() != nil {
		return nil, res.GetError()
	}
	if res.GetData() == nil {
		return nil, enums.ErrorNotFoundRecords
	}

	return res.GetData().(*analysis.Analysis), nil
}

// nolint
func (c *Controller) SaveAnalysis(analysisEntity *analysis.Analysis) (uuid.UUID, error) {
	analysisEntity, err := c.createRepositoryIfNotExists(analysisEntity)
	if err != nil {
		return uuid.Nil, err
	}

	//TODO: REMOVE treatCompatibility IN v2.10.0 VERSION
	if err := c.treatCompatibility(analysisEntity); err != nil {
		return uuid.Nil, err
	}

	analysisDecorated, err := c.decorateAnalysisEntityAndSaveOnDatabase(analysisEntity)
	if err != nil {
		return uuid.Nil, err
	}

	if err := c.publishInBroker(analysisDecorated.ID); err != nil {
		return uuid.Nil, err
	}
	return analysisDecorated.ID, nil
}

func (c *Controller) createRepositoryIfNotExists(analysisEntity *analysis.Analysis) (*analysis.Analysis, error) {
	if analysisEntity.RepositoryID == uuid.Nil {
		analysisEntity.SetRepositoryID(uuid.New())
		repositoryID, err := c.repoRepository.FindRepository(analysisEntity.WorkspaceID, analysisEntity.RepositoryName)
		if err != nil {
			if err == enums.ErrorNotFoundRecords {
				return analysisEntity, c.repoRepository.CreateRepository(analysisEntity.RepositoryID,
					analysisEntity.WorkspaceID, analysisEntity.RepositoryName)
			}
			return nil, err
		}
		analysisEntity.SetRepositoryID(repositoryID)
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
		ID:             analysisEntity.ID,
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

func (c *Controller) publishInBroker(analysisID uuid.UUID) error {
	res, err := c.GetAnalysis(analysisID)
	if err != nil {
		return err
	}

	return c.broker.Publish("", exchange.NewAnalysis,
		exchange.Fanout, res.ToBytes())
}

// TODO:REMOVE ALL BELOW AFTER v2.10.0
// treatCompatibility checks if the field Analysis.AnalysisVulnerabilities[i].DeprecatedHashes exists
// and if so, find them on database and updates it with the correct field Analysis.AnalysisVulnerabilities[i].VulnHash.
// this is only a temporary fix to maintain compatibility between versions,
// it will be deleted when v2.10.0 is released
// nolint
func (c *Controller) treatCompatibility(analysisEntity *analysis.Analysis) error {
	if !c.existsDeprecatedHashesSlice(analysisEntity.AnalysisVulnerabilities) {
		return nil
	}
	deprecatedHashes := make([]string, 0)

	for i := range analysisEntity.AnalysisVulnerabilities {
		deprecatedHashes = append(
			deprecatedHashes,
			analysisEntity.AnalysisVulnerabilities[i].Vulnerability.DeprecatedHashes...)
	}

	if err := c.saveUpdates(deprecatedHashes, analysisEntity); err != nil {
		return err
	}

	return nil
}

func (c *Controller) saveUpdates(hashSlice []string, analysisEntity *analysis.Analysis) error {
	res := c.repoAnalysis.FindVulnerabilitiesByHashSliceInRepository(hashSlice, analysisEntity.RepositoryID)
	if res.GetError() != nil {
		return res.GetError()
	}
	mapHashToID, err := c.parseResIds(res)
	if err != nil {
		return err
	}
	query, values := c.mountUpdateQuery(analysisEntity, mapHashToID)

	if err := c.repoAnalysis.RawQuery(query, values); err != nil {
		return err
	}
	return nil
}

// mountUpdateQuery iterates over rawAnalysis.AnalysisVulnerabilities and
// checks if some vuln.Vulnerability.DeprecatedHashes is present on
// mapHashToId then creates and update statement to update the
// deprecated Hash value to the new one (that is present in rawAnalysis.Vulnerability.VulnHash field)
func (c *Controller) mountUpdateQuery(
	rawAnalysis *analysis.Analysis, mapHashToID map[string]uuid.UUID,
) (string, []string) {
	query := ""
	values := make([]string, 0)
	for i := range rawAnalysis.AnalysisVulnerabilities {
		vuln := rawAnalysis.AnalysisVulnerabilities[i]
		for _, hash := range vuln.Vulnerability.DeprecatedHashes {
			if mapHashToID[hash] != uuid.Nil {
				query += "UPDATE vulnerabilities SET vuln_hash =?  where vulnerability_id = ? ;\n"
				values = append(values, vuln.Vulnerability.VulnHash, mapHashToID[hash].String())
			}
		}
	}
	return query, values
}

// existsDeprecatedHashesSlice checks if []analysis.AnalysisVulnerabilities.Vulnerability
// has a field called DeprecatedHashes
func (c *Controller) existsDeprecatedHashesSlice(vulns []analysis.AnalysisVulnerabilities) bool {
	if len(vulns) > 0 {
		if vulns[0].Vulnerability.DeprecatedHashes != nil {
			return true
		}
	}
	return false
}

// parseResIds makes a map[hash] id that already exists on database for further manipulation
func (c *Controller) parseResIds(res response.IResponse) (map[string]uuid.UUID, error) {
	if res.GetData() == nil {
		return nil, errors.New("nil response.GetData")
	}
	mapIds := res.GetData().(*[]map[string]interface{})
	mapIDHash := make(map[string]uuid.UUID, len(*mapIds))
	for _, id := range *mapIds {
		mapIDHash[id["vuln_hash"].(string)] = uuid.MustParse(id["vulnerability_id"].(string))
	}
	return mapIDHash, nil
}
