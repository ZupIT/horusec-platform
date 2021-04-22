package analysis

import (
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
)

type IUseCase interface {
	ParseAnalysisToVulnerabilitiesByAuthor(analysis *analysis.Analysis) []dashboard.VulnerabilitiesByAuthor
	ParseAnalysisToVulnerabilitiesByRepository(analysis *analysis.Analysis) []dashboard.VulnerabilitiesByRepository
	ParseAnalysisToVulnerabilitiesByLanguage(analysis *analysis.Analysis) []dashboard.VulnerabilitiesByLanguage
	ParseAnalysisToVulnerabilitiesByTime(analysis *analysis.Analysis) []dashboard.VulnerabilitiesByTime
}

type UseCase struct{}

func NewUseCaseAnalysis() IUseCase {
	return &UseCase{}
}

func (u *UseCase) ParseAnalysisToVulnerabilitiesByAuthor(analysisEntity *analysis.Analysis) (entity []dashboard.VulnerabilitiesByAuthor) {
	for _, manyToMany := range analysisEntity.AnalysisVulnerabilities {
		index, exists := u.existsAuthorInList(entity, manyToMany.Vulnerability.CommitEmail)
		if !exists {
			entityToAppend := u.newVulnerabilitiesByAuthor(manyToMany.Vulnerability.CommitEmail, analysisEntity.WorkspaceID, analysisEntity.RepositoryID)
			entityToAppend.AddCountVulnerabilityBySeverity(1, manyToMany.Vulnerability.Severity, manyToMany.Vulnerability.Type)
			entity = append(entity, *entityToAppend)
		} else if len(entity) > 0 {
			entityToAppend := &entity[index]
			entityToAppend.AddCountVulnerabilityBySeverity(1, manyToMany.Vulnerability.Severity, manyToMany.Vulnerability.Type)
			entity[index] = *entityToAppend
		}
	}
	return entity
}

func (u *UseCase) existsAuthorInList(listVulns []dashboard.VulnerabilitiesByAuthor, author string) (int, bool) {
	for index, entity := range listVulns {
		if entity.Author == author {
			return index, true
		}
	}
	return 0, false
}

func (u *UseCase) newVulnerabilitiesByAuthor(author string, workspaceID, repositoryID uuid.UUID) *dashboard.VulnerabilitiesByAuthor {
	return &dashboard.VulnerabilitiesByAuthor{
		Author: author,
		Vulnerability: dashboard.Vulnerability{
			CreatedAt:    time.Now(),
			Active:       true,
			WorkspaceID:  workspaceID,
			RepositoryID: repositoryID,
		},
	}
}

func (u *UseCase) ParseAnalysisToVulnerabilitiesByRepository(analysisEntity *analysis.Analysis) (entity []dashboard.VulnerabilitiesByRepository) {
	for _, manyToMany := range analysisEntity.AnalysisVulnerabilities {
		index, exists := u.existsRepositoryInList(entity, analysisEntity.RepositoryID)
		if !exists {
			entityToAppend := u.newVulnerabilitiesByRepository(manyToMany.Vulnerability.CommitEmail, analysisEntity.WorkspaceID, analysisEntity.RepositoryID)
			entityToAppend.AddCountVulnerabilityBySeverity(1, manyToMany.Vulnerability.Severity, manyToMany.Vulnerability.Type)
			entity = append(entity, *entityToAppend)
		} else if len(entity) > 0 {
			entityToAppend := &entity[index]
			entityToAppend.AddCountVulnerabilityBySeverity(1, manyToMany.Vulnerability.Severity, manyToMany.Vulnerability.Type)
			entity[index] = *entityToAppend
		}
	}
	return entity
}

func (u *UseCase) existsRepositoryInList(listVulns []dashboard.VulnerabilitiesByRepository, repositoryID uuid.UUID) (int, bool) {
	for index, entity := range listVulns {
		if entity.RepositoryID == repositoryID {
			return index, true
		}
	}
	return 0, false
}

func (u *UseCase) newVulnerabilitiesByRepository(repositoryName string, workspaceID, repositoryID uuid.UUID) *dashboard.VulnerabilitiesByRepository {
	return &dashboard.VulnerabilitiesByRepository{
		RepositoryName: repositoryName,
		Vulnerability: dashboard.Vulnerability{
			CreatedAt:    time.Now(),
			Active:       true,
			WorkspaceID:  workspaceID,
			RepositoryID: repositoryID,
		},
	}
}

func (u *UseCase) ParseAnalysisToVulnerabilitiesByLanguage(analysisEntity *analysis.Analysis) (entity []dashboard.VulnerabilitiesByLanguage) {
	for _, manyToMany := range analysisEntity.AnalysisVulnerabilities {
		index, exists := u.existsLanguageInList(entity, manyToMany.Vulnerability.Language)
		if !exists {
			entityToAppend := u.newVulnerabilitiesByLanguage(manyToMany.Vulnerability.Language, analysisEntity.WorkspaceID, analysisEntity.RepositoryID)
			entityToAppend.AddCountVulnerabilityBySeverity(1, manyToMany.Vulnerability.Severity, manyToMany.Vulnerability.Type)
			entity = append(entity, *entityToAppend)
		} else if len(entity) > 0 {
			entityToAppend := &entity[index]
			entityToAppend.AddCountVulnerabilityBySeverity(1, manyToMany.Vulnerability.Severity, manyToMany.Vulnerability.Type)
			entity[index] = *entityToAppend
		}
	}
	return entity
}

func (u *UseCase) existsLanguageInList(listVulns []dashboard.VulnerabilitiesByLanguage, language languages.Language) (int, bool) {
	for index, entity := range listVulns {
		if entity.Language == language {
			return index, true
		}
	}
	return 0, false
}

func (u *UseCase) newVulnerabilitiesByLanguage(language languages.Language, workspaceID, repositoryID uuid.UUID) *dashboard.VulnerabilitiesByLanguage {
	return &dashboard.VulnerabilitiesByLanguage{
		Language: language,
		Vulnerability: dashboard.Vulnerability{
			CreatedAt:    time.Now(),
			Active:       true,
			WorkspaceID:  workspaceID,
			RepositoryID: repositoryID,
		},
	}
}

func (u *UseCase) ParseAnalysisToVulnerabilitiesByTime(analysisEntity *analysis.Analysis) (entity []dashboard.VulnerabilitiesByTime) {
	entityToAppend := &dashboard.VulnerabilitiesByTime{
		Vulnerability: dashboard.Vulnerability{
			CreatedAt:    time.Now(),
			Active:       true,
			WorkspaceID:  analysisEntity.WorkspaceID,
			RepositoryID: analysisEntity.RepositoryID,
		},
	}
	for _, manyToMany := range analysisEntity.AnalysisVulnerabilities {
		entityToAppend.AddCountVulnerabilityBySeverity(1, manyToMany.Vulnerability.Severity, manyToMany.Vulnerability.Type)
	}
	return []dashboard.VulnerabilitiesByTime{*entityToAppend}
}
