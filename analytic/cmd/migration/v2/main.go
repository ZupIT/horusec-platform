package main

import (
	"fmt"

	"github.com/google/uuid"

	analysisEntities "github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseConfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/utils/env"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"

	"github.com/ZupIT/horusec-platform/analytic/cmd/migration/v2/enums"
	dashboardController "github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
	dashboardEnums "github.com/ZupIT/horusec-platform/analytic/internal/enums/dashboard"
	dashboardRepository "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
	dashboardUseCases "github.com/ZupIT/horusec-platform/analytic/internal/usecases/dashboard"
)

type AnalyticMigration struct {
	dbConnectionAnalytic *database.Connection
	dbConnectionHorusec  *database.Connection
	dashboardController  dashboardController.IController
	summary              map[string][]string
}

func NewAnalyticMigrationV2() *AnalyticMigration {
	analyticMigration := &AnalyticMigration{
		dbConnectionAnalytic: getAnalyticDatabaseConnection(),
		dbConnectionHorusec:  getHorusecDatabaseConnection(),
		summary:              setSummary(),
	}

	analyticMigration.dashboardController = dashboardController.NewDashboardController(
		dashboardRepository.NewRepoDashboard(analyticMigration.dbConnectionAnalytic),
		analyticMigration.dbConnectionAnalytic, dashboardUseCases.NewUseCaseDashboard())

	return analyticMigration
}

func setSummary() map[string][]string {
	return map[string][]string{
		enums.SummarySuccess: {},
		enums.SummaryFailed:  {},
	}
}

func getAnalyticDatabaseConnection() *database.Connection {
	analyticDatabaseConfig := databaseConfig.NewDatabaseConfig()
	analyticDatabaseConfig.SetURI(
		env.GetEnvOrDefault(enums.EnvAnalyticDatabaseURI, enums.DefaultAnalyticDatabaseURIValue))

	dbConnectionAnalytic, err := database.NewDatabaseReadAndWrite(analyticDatabaseConfig)
	if err != nil {
		logger.LogPanic(enums.MessageFailedToConnectToAnalyticDatabase, err)
	}

	return dbConnectionAnalytic
}

func getHorusecDatabaseConnection() *database.Connection {
	horusecDatabaseConfig := databaseConfig.NewDatabaseConfig()
	horusecDatabaseConfig.SetURI(env.GetEnvOrDefault(enums.EnvHorusecDatabaseURI, enums.DefaultHorusecDatabaseURIValue))

	dbConnectionHorusec, err := database.NewDatabaseReadAndWrite(horusecDatabaseConfig)
	if err != nil {
		logger.LogPanic(enums.MessageFailedToConnectToHorusecDatabase, err)
	}

	return dbConnectionHorusec
}

func (a *AnalyticMigration) getAllAnalysisWithVulnerabilities() []*analysisEntities.Analysis {
	analysis := &[]*analysisEntities.Analysis{}

	preloads := map[string][]interface{}{
		"AnalysisVulnerabilities":               {},
		"AnalysisVulnerabilities.Vulnerability": {},
	}

	if err := a.dbConnectionHorusec.Read.FindPreload(analysis, map[string]interface{}{},
		preloads, "analysis").GetError(); err != nil {
		logger.LogPanic(enums.MessageFailedToGetAllAnalysis, err)
	}

	return *analysis
}

func (a *AnalyticMigration) loggingRegisterBeingMigrated(analysis *analysisEntities.Analysis) {
	message := fmt.Sprintf(enums.MessageRegisterBeingMigrated, analysis.WorkspaceID, analysis.WorkspaceName,
		analysis.RepositoryID, analysis.RepositoryName, analysis.CreatedAt, len(analysis.AnalysisVulnerabilities))

	logger.LogInfo(message)
}

func (a *AnalyticMigration) setFailedMigrationInSummary(analysisID uuid.UUID, err error, table string) error {
	message := fmt.Sprintf("failed to insert analsysis with id %d in table %s with error -> %v",
		analysisID, table, err)

	a.summary[enums.SummaryFailed] = append(a.summary[enums.SummaryFailed], message)
	return err
}

func (a *AnalyticMigration) setSuccessMigrationInSummary(analysisID uuid.UUID) {
	a.summary[enums.SummarySuccess] = append(a.summary[enums.SummarySuccess], analysisID.String())
}

func (a *AnalyticMigration) migrateAnalysis(analysis *analysisEntities.Analysis) (err error) {
	if err := a.dashboardController.AddVulnerabilitiesByAuthor(analysis); err != nil {
		return a.setFailedMigrationInSummary(analysis.ID, err, dashboardEnums.TableVulnerabilitiesByAuthor)
	}

	if err := a.dashboardController.AddVulnerabilitiesByLanguage(analysis); err != nil {
		return a.setFailedMigrationInSummary(analysis.ID, err, dashboardEnums.TableVulnerabilitiesByLanguage)
	}

	if err := a.dashboardController.AddVulnerabilitiesByRepository(analysis); err != nil {
		return a.setFailedMigrationInSummary(analysis.ID, err, dashboardEnums.TableVulnerabilitiesByRepository)
	}

	return a.setFailedMigrationInSummary(analysis.ID,
		a.dashboardController.AddVulnerabilitiesByTime(analysis), dashboardEnums.TableVulnerabilitiesByTime)
}

func (a *AnalyticMigration) printResults(total int) {
	fmt.Println()
	logger.LogWarn("MIGRATION FINISHED! CHECK THE RESULTS -->")
	logger.LogWarn(fmt.Sprintf("TOTAL DE REGISTROS PARA MIGRAR: %d", total))
	logger.LogWarn(fmt.Sprintf("TOTAL RECORDS THAT SUCCESSFULLY MIGRATED: %d", len(a.summary[enums.SummarySuccess])))
	logger.LogWarn(fmt.Sprintf("TOTAL RECORDS THAT FAILED TO MIGRATE: %d", len(a.summary[enums.SummaryFailed])))
}

func main() {
	analyticMigration := NewAnalyticMigrationV2()

	allPastAnalysis := analyticMigration.getAllAnalysisWithVulnerabilities()
	for _, analysis := range allPastAnalysis {
		analyticMigration.loggingRegisterBeingMigrated(analysis)
		if err := analyticMigration.migrateAnalysis(analysis); err == nil {
			analyticMigration.setSuccessMigrationInSummary(analysis.ID)
		}
	}

	analyticMigration.printResults(len(allPastAnalysis))
}
