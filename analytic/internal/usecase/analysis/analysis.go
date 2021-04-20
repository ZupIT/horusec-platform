package analysis

import (
	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
	"github.com/google/uuid"
	"time"
)

type IUseCase interface {
	ParseAnalysisToVulnerabilitiesByAuthor(analysis *analysis.Analysis) []dashboard.VulnerabilitiesByAuthor
	ParseAnalysisToVulnerabilitiesByRepository(analysis *analysis.Analysis) []dashboard.VulnerabilitiesByRepository
	ParseAnalysisToVulnerabilitiesByLanguage(analysis *analysis.Analysis) []dashboard.VulnerabilitiesByLanguage
	ParseAnalysisToVulnerabilitiesByTime(analysis *analysis.Analysis) []dashboard.VulnerabilitiesByTime
}

type UseCase struct {}

func NewUseCaseAnalysis() IUseCase {
	return &UseCase{}
}

func (u *UseCase) ParseAnalysisToVulnerabilitiesByAuthor(analysisEntity *analysis.Analysis) (entity []dashboard.VulnerabilitiesByAuthor) {
	for _, manyToMany := range analysisEntity.AnalysisVulnerabilities {
		if !u.existsAuthorInList(entity, manyToMany.Vulnerability.CommitEmail) {
			entityToAppend := u.newVulnerabilitiesByAuthor(manyToMany.Vulnerability.CommitEmail, analysisEntity.WorkspaceID, analysisEntity.RepositoryID)
			entityToAppend.AddCountVulnerabilityBySeverity(1, manyToMany.Vulnerability.Severity, manyToMany.Vulnerability.Type)
			entity = append(entity, *entityToAppend)
		}
	}
	return entity
}

func (u *UseCase) existsAuthorInList(listVulns []dashboard.VulnerabilitiesByAuthor, author string) bool {
	for _, entity := range listVulns {
		if entity.Author == author {
			return true
		}
	}
	return false
}

func (u *UseCase) newVulnerabilitiesByAuthor(author string, workspaceID, repositoryID uuid.UUID) *dashboard.VulnerabilitiesByAuthor {
	return &dashboard.VulnerabilitiesByAuthor{
		Author:        author,
		Vulnerability: dashboard.Vulnerability{
			CreatedAt:             time.Now(),
			Active:                true,
			WorkspaceID:           workspaceID,
			RepositoryID:          repositoryID,
		},
	}
}

func (u *UseCase) ParseAnalysisToVulnerabilitiesByRepository(analysisEntity *analysis.Analysis) (entity []dashboard.VulnerabilitiesByRepository) {
	for _, manyToMany := range analysisEntity.AnalysisVulnerabilities {
		if !u.existsRepositoryInList(entity, manyToMany.Vulnerability.CommitEmail) {
			entityToAppend := u.newVulnerabilitiesByRepository(manyToMany.Vulnerability.CommitEmail, analysisEntity.WorkspaceID, analysisEntity.RepositoryID)
			entityToAppend.AddCountVulnerabilityBySeverity(1, manyToMany.Vulnerability.Severity, manyToMany.Vulnerability.Type)
			entity = append(entity, *entityToAppend)
		}
	}
	return entity
}

func (u *UseCase) existsRepositoryInList(listVulns []dashboard.VulnerabilitiesByRepository, repositoryName string) bool {
	for _, entity := range listVulns {
		if entity.RepositoryName == repositoryName {
			return true
		}
	}
	return false
}

func (u *UseCase) newVulnerabilitiesByRepository(repositoryName string, workspaceID, repositoryID uuid.UUID) *dashboard.VulnerabilitiesByRepository {
	return &dashboard.VulnerabilitiesByRepository{
		RepositoryName:        repositoryName,
		Vulnerability: dashboard.Vulnerability{
			CreatedAt:             time.Now(),
			Active:                true,
			WorkspaceID:           workspaceID,
			RepositoryID:          repositoryID,
		},
	}
}

func (u *UseCase) ParseAnalysisToVulnerabilitiesByLanguage(analysisEntity *analysis.Analysis) (entity []dashboard.VulnerabilitiesByLanguage) {
	for _, manyToMany := range analysisEntity.AnalysisVulnerabilities {
		if !u.existsLanguageInList(entity, manyToMany.Vulnerability.Language) {
			entityToAppend := u.newVulnerabilitiesByLanguage(manyToMany.Vulnerability.Language, analysisEntity.WorkspaceID, analysisEntity.RepositoryID)
			entityToAppend.AddCountVulnerabilityBySeverity(1, manyToMany.Vulnerability.Severity, manyToMany.Vulnerability.Type)
			entity = append(entity, *entityToAppend)
		}
	}
	return entity
}

func (u *UseCase) existsLanguageInList(listVulns []dashboard.VulnerabilitiesByLanguage, language languages.Language) bool {
	for _, entity := range listVulns {
		if entity.Language == language {
			return true
		}
	}
	return false
}

func (u *UseCase) newVulnerabilitiesByLanguage(language languages.Language, workspaceID, repositoryID uuid.UUID) *dashboard.VulnerabilitiesByLanguage {
	return &dashboard.VulnerabilitiesByLanguage{
		Language:        language,
		Vulnerability: dashboard.Vulnerability{
			CreatedAt:             time.Now(),
			Active:                true,
			WorkspaceID:           workspaceID,
			RepositoryID:          repositoryID,
		},
	}
}

func (u *UseCase) ParseAnalysisToVulnerabilitiesByTime(analysisEntity *analysis.Analysis) (entity []dashboard.VulnerabilitiesByTime) {
	for _, manyToMany := range analysisEntity.AnalysisVulnerabilities {
		entityToAppend := &dashboard.VulnerabilitiesByTime{
			Vulnerability: dashboard.Vulnerability{
				CreatedAt:             time.Now(),
				Active:                true,
				WorkspaceID:           analysisEntity.WorkspaceID,
				RepositoryID:          analysisEntity.RepositoryID,
			},
		}
		entityToAppend.AddCountVulnerabilityBySeverity(1, manyToMany.Vulnerability.Severity, manyToMany.Vulnerability.Type)
		entity = append(entity, *entityToAppend)
	}
	return entity
}
