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
	"context"
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/services/tracer"

	"github.com/opentracing/opentracing-go"

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
	GetAnalysis(ctx context.Context, analysisID uuid.UUID) (*analysis.Analysis, error)
	SaveAnalysis(ctx context.Context, analysisEntity *analysis.Analysis) (uuid.UUID, error)
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

func (c *Controller) GetAnalysis(ctx context.Context, analysisID uuid.UUID) (*analysis.Analysis, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetAnalysis")
	defer span.Finish()
	response := c.repoAnalysis.FindAnalysisByID(ctx, analysisID)
	if err := response.GetError(); err != nil {
		tracer.SetSpanError(span, err)
		return nil, err
	}
	if response.GetData() == nil {
		tracer.SetSpanError(span, enums.ErrorNotFoundRecords)
		return nil, enums.ErrorNotFoundRecords
	}

	return response.GetData().(*analysis.Analysis), nil
}

func (c *Controller) SaveAnalysis(ctx context.Context, analysisEntity *analysis.Analysis) (uuid.UUID, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "SaveAnalysis")
	defer span.Finish()
	analysisEntity, err := c.createRepositoryIfNotExists(ctx, analysisEntity)
	if err != nil {
		return uuid.Nil, err
	}
	analysisDecorated, err := c.decorateAnalysisEntityAndSaveOnDatabase(ctx, analysisEntity)
	if err != nil {
		return uuid.Nil, err
	}
	if err := c.publishInBroker(ctx, analysisDecorated.ID); err != nil {
		return uuid.Nil, err
	}
	return analysisDecorated.ID, nil
}

func (c *Controller) createRepositoryIfNotExists(ctx context.Context, analysisEntity *analysis.Analysis) (*analysis.Analysis, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "createRepositoryIfNotExists")
	defer span.Finish()
	if analysisEntity.RepositoryID == uuid.Nil {
		analysisEntity.SetRepositoryID(uuid.New())
		repositoryID, err := c.repoRepository.FindRepository(ctx, analysisEntity.WorkspaceID, analysisEntity.RepositoryName)
		if err != nil {
			if err == enums.ErrorNotFoundRecords {
				return analysisEntity, c.repoRepository.CreateRepository(ctx, analysisEntity.RepositoryID,
					analysisEntity.WorkspaceID, analysisEntity.RepositoryName)
			}
			return nil, err
		}
		analysisEntity.SetRepositoryID(repositoryID)
	}
	return analysisEntity, nil
}

func (c *Controller) decorateAnalysisEntityAndSaveOnDatabase(ctx context.Context,
	analysisEntity *analysis.Analysis) (*analysis.Analysis, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "decorateAnalysisEntityAndSaveOnDatabase")
	defer span.Finish()
	analysisDecorated := c.decoratorAnalysisToSave(analysisEntity)
	return analysisDecorated, c.createNewAnalysis(ctx, analysisDecorated)
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

func (c *Controller) createNewAnalysis(ctx context.Context, newAnalysis *analysis.Analysis) error {
	return c.repoAnalysis.CreateFullAnalysis(ctx, newAnalysis)
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

func (c *Controller) publishInBroker(ctx context.Context, analysisID uuid.UUID) error {
	response, err := c.GetAnalysis(ctx, analysisID)
	if err != nil {
		return err
	}

	return c.broker.Publish("", exchange.NewAnalysis,
		exchange.Fanout, response.ToBytes())
}
