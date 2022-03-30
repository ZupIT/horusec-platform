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
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"

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

func (c *Controller) SaveAnalysis(analysisEntity *analysis.Analysis) (uuid.UUID, error) {
	analysisEntity, err := c.createRepositoryIfNotExists(analysisEntity)
	if err != nil {
		return uuid.Nil, err
	}
	analysisDecorated, err := c.saveNewAnalysisInDatabase(analysisEntity)
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

func (c *Controller) saveNewAnalysisInDatabase(newAnalysis *analysis.Analysis) (*analysis.Analysis, error) {
	//TODO: REMOVE treatHashCompatibility IN v2.10.0 VERSION
	if err := c.treatHashCompatibility(newAnalysis); err != nil {
		return nil, err
	}

	analysisDecorated := c.decoratorAnalysisToSave(newAnalysis)

	if err := c.repoAnalysis.CreateFullAnalysis(analysisDecorated); err != nil {
		return nil, err
	}

	return analysisDecorated, nil
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
// treatHashCompatibility checks if the field Analysis.AnalysisVulnerabilities[i].DeprecatedHashes exists
// and if so, find them on database and updates it with the correct field Analysis.AnalysisVulnerabilities[i].VulnHash.
// this is only a temporary fix to maintain compatibility between versions,
// it will be deleted when v2.10.0 is released
// nolint:funlen,gocyclo // funlen and gocyclo is not necessary in this method deprecated
func (c *Controller) treatHashCompatibility(analysisEntity *analysis.Analysis) error {
	if !c.existsDeprecatedHashesSlice(analysisEntity.AnalysisVulnerabilities) {
		return nil
	}
	deprecatedHashes := []string{}
	for index := range analysisEntity.AnalysisVulnerabilities {
		manyToMany := analysisEntity.AnalysisVulnerabilities[index]
		deprecatedHashes = append(deprecatedHashes, manyToMany.Vulnerability.DeprecatedHashes...)
	}
	mapHashToVulnerabilityID, err := c.getMapHashToVulnerabilityID(deprecatedHashes, analysisEntity)
	if err != nil || len(mapHashToVulnerabilityID) == 0 {
		// When the database cannot find data, it means that the user is using the new CLI version,
		// but Horusec-Platform is new and has not yet received any analysis for this repository.
		if err == enums.ErrorNotFoundRecords {
			return nil
		}
		return err
	}
	return c.repoAnalysis.SaveTreatCompatibility(mapHashToVulnerabilityID, analysisEntity)
}

// existsDeprecatedHashesSlice will get all vulnerabilities from database with deprecated hashes
// nolint:funlen // funlen is not necessary in this method deprecated
func (c *Controller) getMapHashToVulnerabilityID(deprecatedHashes []string,
	analysisEntity *analysis.Analysis) (map[string]uuid.UUID, error) {
	mapHashToVulnerabilityID := make(map[string]uuid.UUID, 0)
	res := c.repoAnalysis.FindAllVulnerabilitiesByHashesAndRepository(deprecatedHashes, analysisEntity.RepositoryID)
	if err := res.GetError(); err != nil {
		return map[string]uuid.UUID{}, err
	}
	if res.GetData() == nil {
		return map[string]uuid.UUID{}, nil
	}
	vulnerabilities := *res.GetData().(*[]vulnerability.Vulnerability)
	for index := range vulnerabilities {
		vuln := vulnerabilities[index]
		mapHashToVulnerabilityID[vuln.VulnHash] = vuln.VulnerabilityID
	}
	return mapHashToVulnerabilityID, nil
}

// existsDeprecatedHashesSlice checks if []analysis.AnalysisVulnerabilities.Vulnerability
// has a field called DeprecatedHashes
func (c *Controller) existsDeprecatedHashesSlice(sliceManyToMany []analysis.AnalysisVulnerabilities) bool {
	for index := range sliceManyToMany {
		manyToMany := sliceManyToMany[index]
		if len(manyToMany.Vulnerability.DeprecatedHashes) > 0 {
			return true
		}
	}
	return false
}
