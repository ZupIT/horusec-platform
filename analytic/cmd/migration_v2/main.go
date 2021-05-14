package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	analysisEntities "github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseConfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/utils/env"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
)

const (
	EnvAnalyticDatabaseURI                   = "HORUSEC_DATABASE_ANALYTIC_SQL_URI"
	DefaultAnalyticDatabaseURIValue          = "postgresql://root:root@localhost:5432/horusec_analytic_db?sslmode=disable"
	MessageFailedToConnectToAnalyticDatabase = "failed to connect to analytic database, please check your uri informed in HORUSEC_DATABASE_ANALYTIC_SQL_URI env variable"
	EnvHorusecDatabaseURI                    = "HORUSEC_DATABASE_HORUSEC_SQL_URI"
	DefaultHorusecDatabaseURIValue           = "postgresql://root:root@localhost:5432/horusec_db?sslmode=disable"
	MessageFailedToConnectToHorusecDatabase  = "failed to connect to horusec default database, please check your uri informed in HORUSEC_DATABASE_HORUSEC_SQL_URI env variable"
	MessageFailedToGetAllAnalysis            = "something went wrong while getting all your past analysis please check your informed in HORUSEC_DATABASE_HORUSEC_SQL_URI env variable"
	MessageRegisterBeingMigrated             = "Stating to migrate analysis with: WorkspaceID -> %s | Workspace Name -> %s | Repository ID %s | Repository Name %s | Created At -> %v | Total Of Vulnerabilities -> %d"
	SummarySuccess                           = "success"
	SummaryFailed                            = "failed"
	TableAddVulnerabilitiesByAuthor          = "vulnerabilities_by_author"
	TableAddVulnerabilitiesByLanguage        = "vulnerabilities_by_language"
	TableAddVulnerabilitiesByRepository      = "vulnerabilities_by_repository"
	TableAddVulnerabilitiesByTime            = "vulnerabilities_by_time"
)

func main() {
	analyticMigration := NewAnalyticMigrationV2()

	allPastAnalysis := analyticMigration.getAllAnalysisWithVulnerabilities()
	for _, analysis := range allPastAnalysis {
		analyticMigration.loggingRegisterBeingMigrated(analysis.WorkspaceName, analysis.RepositoryName,
			analysis.WorkspaceID, analysis.RepositoryID, analysis.CreatedAt, len(analysis.AnalysisVulnerabilities))

		if err := analyticMigration.migrateAnalysis(&analysis); err == nil {
			analyticMigration.setSuccessMigrationInSummary(analysis.ID)
		}
	}

	analyticMigration.printResults(len(allPastAnalysis))
}

type AnalyticMigrationV2 struct {
	dbConnectionAnalytic *database.Connection
	dbConnectionHorusec  *database.Connection
	//dashboardUseCases    dashboardUseCases.IUseCase
	//dashboardController  dashboard.IWriteController
	summary map[string][]string
}

func NewAnalyticMigrationV2() *AnalyticMigrationV2 {
	test := &AnalyticMigrationV2{
		dbConnectionAnalytic: getAnalyticDatabaseConnection(),
		dbConnectionHorusec:  getHorusecDatabaseConnection(),
		summary:              setSummary(),
	}

	//repo := dashboard2.NewRepoDashboard(test.dbConnectionAnalytic)

	//test.dashboardController = dashboard.NewControllerDashboardWrite(repo, test.dbConnectionAnalytic.Write)

	return test
}

func setSummary() map[string][]string {
	return map[string][]string{
		SummarySuccess: {},
		SummaryFailed:  {},
	}
}

func getAnalyticDatabaseConnection() *database.Connection {
	analyticDatabaseConfig := databaseConfig.NewDatabaseConfig()
	analyticDatabaseConfig.SetURI(env.GetEnvOrDefault(EnvAnalyticDatabaseURI, DefaultAnalyticDatabaseURIValue))

	dbConnectionAnalytic, err := database.NewDatabaseReadAndWrite(analyticDatabaseConfig)
	if err != nil {
		logger.LogPanic(MessageFailedToConnectToAnalyticDatabase, err)
	}

	return dbConnectionAnalytic
}

func getHorusecDatabaseConnection() *database.Connection {
	horusecDatabaseConfig := databaseConfig.NewDatabaseConfig()
	horusecDatabaseConfig.SetURI(env.GetEnvOrDefault(EnvHorusecDatabaseURI, DefaultHorusecDatabaseURIValue))

	dbConnectionHorusec, err := database.NewDatabaseReadAndWrite(horusecDatabaseConfig)
	if err != nil {
		logger.LogPanic(MessageFailedToConnectToHorusecDatabase, err)
	}

	return dbConnectionHorusec
}

func (a *AnalyticMigrationV2) getAllAnalysisWithVulnerabilities() []analysisEntities.Analysis {
	analysis := &[]analysisEntities.Analysis{}

	preloads := map[string][]interface{}{
		"AnalysisVulnerabilities":               {},
		"AnalysisVulnerabilities.Vulnerability": {},
	}

	if err := a.dbConnectionHorusec.Read.FindPreload(analysis, map[string]interface{}{},
		preloads, "analysis").GetError(); err != nil {
		logger.LogPanic(MessageFailedToGetAllAnalysis, err)
	}

	return *analysis
}

func (a *AnalyticMigrationV2) loggingRegisterBeingMigrated(workspaceName, repositoryName string, workspaceID,
	repositoryID uuid.UUID, createdAt time.Time, totalOfVulnerabilities int) {
	message := fmt.Sprintf(MessageRegisterBeingMigrated, workspaceID, workspaceName, repositoryID,
		repositoryName, createdAt, totalOfVulnerabilities)

	logger.LogInfo(message)
}

func (a *AnalyticMigrationV2) setFailedMigrationInSummary(analysisID uuid.UUID, err error, table string) error {
	message := fmt.Sprintf("failed to insert analsysis with id %d in table %s with error -> %v",
		analysisID, table, err)

	a.summary[SummaryFailed] = append(a.summary[analysisID.String()], message)
	return err
}

func (a *AnalyticMigrationV2) setSuccessMigrationInSummary(analysisID uuid.UUID) {
	a.summary[SummarySuccess] = append(a.summary[SummarySuccess], analysisID.String())
}

func (a *AnalyticMigrationV2) migrateAnalysis(analysis *analysisEntities.Analysis) (err error) {
	//if err := a.dashboardController.AddVulnerabilitiesByAuthor(analysis); err != nil {
	//	return a.setFailedMigrationInSummary(analysis.ID, err, TableAddVulnerabilitiesByAuthor)
	//}
	//
	//if err := a.dashboardController.AddVulnerabilitiesByLanguage(analysis); err != nil {
	//	return a.setFailedMigrationInSummary(analysis.ID, err, TableAddVulnerabilitiesByLanguage)
	//}

	//if err := a.dashboardController.AddVulnerabilitiesByRepository(analysis); err != nil {
	//	return a.setFailedMigrationInSummary(analysis.ID, err, TableAddVulnerabilitiesByRepository)
	//}

	//if err := a.dashboardController.AddVulnerabilitiesByTime(analysis); err != nil {
	//	return a.setFailedMigrationInSummary(analysis.ID, err, TableAddVulnerabilitiesByTime)
	//}

	return nil
}

func (a *AnalyticMigrationV2) printResults(total int) {
	fmt.Println()
	logger.LogWarn("MIGRATION FINISHED! CHECK THE RESULTS -->")
	logger.LogWarn(fmt.Sprintf("TOTAL DE REGISTROS PARA MIGRAR: %d", total))
	logger.LogWarn(fmt.Sprintf("TOTAL RECORDS THAT SUCCESSFULLY MIGRATED: %d", len(a.summary[SummarySuccess])))
	logger.LogWarn(fmt.Sprintf("TOTAL RECORDS THAT FAILED TO MIGRATE: %d", len(a.summary[SummaryFailed])))
}

func (a *AnalyticMigrationV2) save(entity interface{}, table string) error {
	return a.dbConnectionAnalytic.Write.Create(entity, table).GetError()
}
