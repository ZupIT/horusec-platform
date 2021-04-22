package analysis

import (
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"

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

func (u *UseCase) ParseAnalysisToVulnerabilitiesByAuthor(
	analysisEntity *analysis.Analysis) (entitiesByAuthor []dashboard.VulnerabilitiesByAuthor) {
	for indexNN := range analysisEntity.AnalysisVulnerabilities {
		manyToMany := analysisEntity.AnalysisVulnerabilities[indexNN]
		entitiesByAuthor = u.appendEntitiesByAuthor(&manyToMany,
			analysisEntity.WorkspaceID, analysisEntity.RepositoryID, entitiesByAuthor)
	}
	return entitiesByAuthor
}

// nolint:dupl // method is not necessary join with others
func (u *UseCase) appendEntitiesByAuthor(manyToMany *analysis.AnalysisVulnerabilities,
	workspaceID, repositoryID uuid.UUID,
	entitiesByAuthor []dashboard.VulnerabilitiesByAuthor) []dashboard.VulnerabilitiesByAuthor {
	index, exists := u.existsAuthorInList(entitiesByAuthor, manyToMany.Vulnerability.CommitEmail)
	if !exists {
		toAppend := u.newVulnerabilitiesByAuthor(manyToMany.Vulnerability.CommitEmail,
			workspaceID, repositoryID, &manyToMany.Vulnerability)
		entitiesByAuthor = append(entitiesByAuthor, toAppend)
	} else if len(entitiesByAuthor) > 0 {
		entityToAppend := &entitiesByAuthor[index]
		entityToAppend.AddCountVulnerabilityBySeverity(
			1, manyToMany.Vulnerability.Severity, manyToMany.Vulnerability.Type)
		entitiesByAuthor[index] = *entityToAppend
	}
	return entitiesByAuthor
}

func (u *UseCase) existsAuthorInList(listVulns []dashboard.VulnerabilitiesByAuthor, author string) (int, bool) {
	for index := range listVulns {
		if listVulns[index].Author == author {
			return index, true
		}
	}
	return 0, false
}

func (u *UseCase) newVulnerabilitiesByAuthor(author string,
	workspaceID, repositoryID uuid.UUID, vuln *vulnerability.Vulnerability) dashboard.VulnerabilitiesByAuthor {
	entity := &dashboard.VulnerabilitiesByAuthor{
		Author: author,
		Vulnerability: dashboard.Vulnerability{
			VulnerabilityID: uuid.New(),
			CreatedAt:       time.Now(),
			Active:          true,
			WorkspaceID:     workspaceID,
			RepositoryID:    repositoryID,
		},
	}
	entity.AddCountVulnerabilityBySeverity(1, vuln.Severity, vuln.Type)
	return *entity
}

func (u *UseCase) ParseAnalysisToVulnerabilitiesByRepository(
	analysisEntity *analysis.Analysis) (entitiesByRepository []dashboard.VulnerabilitiesByRepository) {
	for indexNN := range analysisEntity.AnalysisVulnerabilities {
		entitiesByRepository = u.appendEntitiesByRepository(&analysisEntity.AnalysisVulnerabilities[indexNN],
			analysisEntity.WorkspaceID, analysisEntity.RepositoryID, entitiesByRepository)
	}
	return entitiesByRepository
}

func (u *UseCase) appendEntitiesByRepository(manyToMany *analysis.AnalysisVulnerabilities,
	workspaceID, repositoryID uuid.UUID,
	entitiesByRepository []dashboard.VulnerabilitiesByRepository) []dashboard.VulnerabilitiesByRepository {
	index, exists := u.existsRepositoryInList(entitiesByRepository, repositoryID)
	if !exists {
		toAppend := u.newVulnerabilitiesByRepository(manyToMany.Vulnerability.CommitEmail,
			workspaceID, repositoryID, &manyToMany.Vulnerability)
		entitiesByRepository = append(entitiesByRepository, toAppend)
	} else if len(entitiesByRepository) > 0 {
		entityToAppend := &entitiesByRepository[index]
		entityToAppend.AddCountVulnerabilityBySeverity(
			1, manyToMany.Vulnerability.Severity, manyToMany.Vulnerability.Type)
		entitiesByRepository[index] = *entityToAppend
	}
	return entitiesByRepository
}

func (u *UseCase) existsRepositoryInList(listVulns []dashboard.VulnerabilitiesByRepository,
	repositoryID uuid.UUID) (int, bool) {
	for index := range listVulns {
		if listVulns[index].RepositoryID == repositoryID {
			return index, true
		}
	}
	return 0, false
}

func (u *UseCase) newVulnerabilitiesByRepository(repositoryName string, workspaceID, repositoryID uuid.UUID,
	vuln *vulnerability.Vulnerability) dashboard.VulnerabilitiesByRepository {
	entity := &dashboard.VulnerabilitiesByRepository{
		RepositoryName: repositoryName,
		Vulnerability: dashboard.Vulnerability{
			VulnerabilityID: uuid.New(),
			CreatedAt:       time.Now(),
			Active:          true,
			WorkspaceID:     workspaceID,
			RepositoryID:    repositoryID,
		},
	}
	entity.AddCountVulnerabilityBySeverity(1, vuln.Severity, vuln.Type)
	return *entity
}

func (u *UseCase) ParseAnalysisToVulnerabilitiesByLanguage(
	analysisEntity *analysis.Analysis) (entitiesByLanguage []dashboard.VulnerabilitiesByLanguage) {
	for indexNN := range analysisEntity.AnalysisVulnerabilities {
		manyToMany := analysisEntity.AnalysisVulnerabilities[indexNN]
		entitiesByLanguage = u.appendEntitiesByLanguage(&manyToMany,
			analysisEntity.WorkspaceID, analysisEntity.RepositoryID, entitiesByLanguage)
	}
	return entitiesByLanguage
}

// nolint:dupl // method is not necessary join with others
func (u *UseCase) appendEntitiesByLanguage(manyToMany *analysis.AnalysisVulnerabilities,
	workspaceID, repositoryID uuid.UUID,
	entitiesByLanguage []dashboard.VulnerabilitiesByLanguage) []dashboard.VulnerabilitiesByLanguage {
	index, exists := u.existsLanguageInList(entitiesByLanguage, manyToMany.Vulnerability.Language)
	if !exists {
		toAppend := u.newVulnerabilitiesByLanguage(manyToMany.Vulnerability.Language,
			workspaceID, repositoryID, &manyToMany.Vulnerability)
		entitiesByLanguage = append(entitiesByLanguage, toAppend)
	} else if len(entitiesByLanguage) > 0 {
		entityToAppend := &entitiesByLanguage[index]
		entityToAppend.AddCountVulnerabilityBySeverity(
			1, manyToMany.Vulnerability.Severity, manyToMany.Vulnerability.Type)
		entitiesByLanguage[index] = *entityToAppend
	}
	return entitiesByLanguage
}

func (u *UseCase) existsLanguageInList(
	listVulns []dashboard.VulnerabilitiesByLanguage, language languages.Language) (int, bool) {
	for index := range listVulns {
		if listVulns[index].Language == language {
			return index, true
		}
	}
	return 0, false
}

func (u *UseCase) newVulnerabilitiesByLanguage(language languages.Language,
	workspaceID, repositoryID uuid.UUID, vuln *vulnerability.Vulnerability) dashboard.VulnerabilitiesByLanguage {
	entity := &dashboard.VulnerabilitiesByLanguage{
		Language: language,
		Vulnerability: dashboard.Vulnerability{
			VulnerabilityID: uuid.New(),
			CreatedAt:       time.Now(),
			Active:          true,
			WorkspaceID:     workspaceID,
			RepositoryID:    repositoryID,
		},
	}
	entity.AddCountVulnerabilityBySeverity(1, vuln.Severity, vuln.Type)
	return *entity
}

func (u *UseCase) ParseAnalysisToVulnerabilitiesByTime(
	analysisEntity *analysis.Analysis) (entititesByTime []dashboard.VulnerabilitiesByTime) {
	entityToAppend := &dashboard.VulnerabilitiesByTime{
		Vulnerability: dashboard.Vulnerability{
			VulnerabilityID: uuid.New(),
			CreatedAt:       time.Now(),
			Active:          true,
			WorkspaceID:     analysisEntity.WorkspaceID,
			RepositoryID:    analysisEntity.RepositoryID,
		},
	}
	for indexNN := range analysisEntity.AnalysisVulnerabilities {
		manyToMany := analysisEntity.AnalysisVulnerabilities[indexNN]
		entityToAppend.AddCountVulnerabilityBySeverity(1, manyToMany.Vulnerability.Severity, manyToMany.Vulnerability.Type)
	}
	return []dashboard.VulnerabilitiesByTime{*entityToAppend}
}
