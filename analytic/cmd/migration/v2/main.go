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

package main

import (
	"fmt"
	"os"

	analysisentities "github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseconfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/utils/env"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-platform/analytic/cmd/migration/v2/enums"
	dashboardcontroller "github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
	dashboardenums "github.com/ZupIT/horusec-platform/analytic/internal/enums/dashboard"
	dashboardrepository "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
	dashboardusecases "github.com/ZupIT/horusec-platform/analytic/internal/usecases/dashboard"
)

type AnalyticMigration struct {
	dbConnectionAnalytic *database.Connection
	dbConnectionHorusec  *database.Connection
	dashboardController  dashboardcontroller.IController
	summary              map[string][]string
}

func NewAnalyticMigrationV2() *AnalyticMigration {
	analyticMigration := &AnalyticMigration{
		dbConnectionAnalytic: getAnalyticDatabaseConnection(),
		dbConnectionHorusec:  getHorusecDatabaseConnection(),
		summary:              setSummary(),
	}

	analyticMigration.dashboardController = dashboardcontroller.NewDashboardController(
		dashboardrepository.NewRepoDashboard(analyticMigration.dbConnectionAnalytic),
		dashboardrepository.NewWorkspaceDashboard(analyticMigration.dbConnectionAnalytic),
		analyticMigration.dbConnectionAnalytic, dashboardusecases.NewUseCaseDashboard())

	return analyticMigration
}

func setSummary() map[string][]string {
	return map[string][]string{
		enums.SummarySuccess: {},
		enums.SummaryFailed:  {},
	}
}

func getAnalyticDatabaseConnection() *database.Connection {
	analyticDatabaseConfig := databaseconfig.NewDatabaseConfig()
	analyticDatabaseConfig.SetURI(
		env.GetEnvOrDefault(enums.EnvAnalyticDatabaseURI, enums.DefaultAnalyticDatabaseURIValue))

	dbConnectionAnalytic, err := database.NewDatabaseReadAndWrite(analyticDatabaseConfig)
	if err != nil {
		logger.LogPanic(enums.MessageFailedToConnectToAnalyticDatabase, err)
	}

	return dbConnectionAnalytic
}

func getHorusecDatabaseConnection() *database.Connection {
	horusecDatabaseConfig := databaseconfig.NewDatabaseConfig()
	horusecDatabaseConfig.SetURI(env.GetEnvOrDefault(enums.EnvHorusecDatabaseURI, enums.DefaultHorusecDatabaseURIValue))

	dbConnectionHorusec, err := database.NewDatabaseReadAndWrite(horusecDatabaseConfig)
	if err != nil {
		logger.LogPanic(enums.MessageFailedToConnectToHorusecDatabase, err)
	}

	return dbConnectionHorusec
}

// nolint:funlen // recursive method is not necessary broken
func (a *AnalyticMigration) getAllAnalysisWithVulnerabilities(limit, page int,
	analysis []*analysisentities.Analysis) []*analysisentities.Analysis {
	logger.LogInfo(fmt.Sprintf("Getting vulnerabilities with limit: %v and page %v", limit, page))

	currentAnalysis := []*analysisentities.Analysis{}
	preloads := map[string][]interface{}{
		"AnalysisVulnerabilities":               {},
		"AnalysisVulnerabilities.Vulnerability": {},
	}

	if err := a.dbConnectionHorusec.Read.FindPreloadWitLimitAndPage(&currentAnalysis, map[string]interface{}{},
		preloads, "analysis", limit, page).GetErrorExceptNotFound(); err != nil {
		logger.LogPanic(enums.MessageFailedToGetAllAnalysis, err)
	}

	if len(currentAnalysis) > 0 {
		analysis = append(analysis, currentAnalysis...)

		return a.getAllAnalysisWithVulnerabilities(limit, page+1, analysis)
	}

	logger.LogInfo(fmt.Sprintf("Total analysis found: %v", len(analysis)))

	return analysis
}

func (a *AnalyticMigration) loggingRegisterBeingMigrated(analysis *analysisentities.Analysis) {
	message := fmt.Sprintf(enums.MessageRegisterBeingMigrated, analysis.WorkspaceID, analysis.WorkspaceName,
		analysis.RepositoryID, analysis.RepositoryName, analysis.CreatedAt, len(analysis.AnalysisVulnerabilities))

	logger.LogInfo(message)
}

func (a *AnalyticMigration) setMigrationInSummary(analysisID uuid.UUID, err error, table string) {
	if err == nil {
		message := fmt.Sprintf("analysis with id '%s' migrated with success on table '%s'",
			analysisID.String(), table)
		a.summary[enums.SummarySuccess] = append(a.summary[enums.SummarySuccess], message)

		return
	}

	message := fmt.Sprintf("failed to migrate analsysis with id '%s' on table '%s' with error -> '%v'",
		analysisID.String(), table, err)
	a.summary[enums.SummaryFailed] = append(a.summary[enums.SummaryFailed], message)
}

func (a *AnalyticMigration) migrateAnalysis(analysis *analysisentities.Analysis) {
	a.setMigrationInSummary(analysis.ID, a.dashboardController.AddVulnerabilitiesByAuthor(analysis),
		dashboardenums.TableVulnerabilitiesByAuthor)

	a.setMigrationInSummary(analysis.ID, a.dashboardController.AddVulnerabilitiesByLanguage(analysis),
		dashboardenums.TableVulnerabilitiesByLanguage)

	a.setMigrationInSummary(analysis.ID, a.dashboardController.AddVulnerabilitiesByRepository(analysis),
		dashboardenums.TableVulnerabilitiesByRepository)

	a.setMigrationInSummary(analysis.ID, a.dashboardController.AddVulnerabilitiesByTime(analysis),
		dashboardenums.TableVulnerabilitiesByTime)
}

func (a *AnalyticMigration) printResults(total int) {
	a.createResultLog()

	// nolint:forbidigo // only usage for the best experience on logs
	fmt.Println()
	logger.LogWarn("MIGRATION FINISHED! CHECK THE RESULTS -->")
	logger.LogWarn(fmt.Sprintf("TOTAL RECORDS TO MIGRATE: %d", total))
	logger.LogWarn(fmt.Sprintf("TOTAL RECORDS SUCCESSFULLY MIGRATED: %d", len(a.summary[enums.SummarySuccess])))
	logger.LogWarn(fmt.Sprintf("TOTAL RECORDS THAT FAILED TO MIGRATE: %d", len(a.summary[enums.SummaryFailed])))
	logger.LogWarn("YOU CAN SEE THE COMPLETE RESULT IN '/tmp/v1-to-v2-horusec-analytic-result'")
}

func (a *AnalyticMigration) createResultLog() {
	result := "RESULT HORUSEC ANALYTIC MIGRATION V1 TO V2\n\nANALYSIS ID MIGRATED WITHOUT ERRORS -->\n"
	for _, value := range a.summary[enums.SummarySuccess] {
		result += fmt.Sprintf("SUCCESS: %s\n", value)
	}

	result += "\nANALYSIS ID AND TABLE THAT FAILED TO MIGRATE -->\n"
	for _, value := range a.summary[enums.SummaryFailed] {
		result += fmt.Sprintf("FAILED: %s\n", value)
	}

	file, _ := os.Create("./tmp/v1-to-v2-horusec-analytic-result")
	_, _ = file.WriteString(result)
}

func (a *AnalyticMigration) existsAnalyticMigrations() (bool, error) {
	var migrations []string

	res := a.dbConnectionAnalytic.Read.Find(&migrations, map[string]interface{}{}, enums.MigrationTable)
	if res.GetErrorExceptNotFound() != nil {
		return false, res.GetErrorExceptNotFound()
	}

	for _, v := range migrations {
		if v == enums.MigrationV1ToV2Name {
			return true, nil
		}
	}

	return a.validateDatabasesIfWasEmpty()
}

func (a *AnalyticMigration) validateDatabasesIfWasEmpty() (bool, error) {
	existsHorusec, err := a.existsVulnerabilitiesInHorusec()
	if err != nil {
		return false, err
	}

	existsAnalytic, err := a.existsVulnerabilitiesInAnalytic()
	if err != nil {
		return false, err
	}

	if existsHorusec && existsAnalytic {
		return true, a.saveMigration()
	}

	return false, nil
}

func (a *AnalyticMigration) existsVulnerabilitiesInHorusec() (bool, error) {
	var vulnerabilitiesHorusec []vulnerability.Vulnerability

	resHorusec := a.dbConnectionHorusec.Read.Find(&vulnerabilitiesHorusec, map[string]interface{}{}, enums.MigrationTable)
	if resHorusec.GetErrorExceptNotFound() != nil {
		return false, resHorusec.GetErrorExceptNotFound()
	}

	return len(vulnerabilitiesHorusec) > 0, nil
}

func (a *AnalyticMigration) existsVulnerabilitiesInAnalytic() (bool, error) {
	var vulnerabilitiesAnalytic []dashboard.VulnerabilitiesByLanguage

	resHorusec := a.dbConnectionAnalytic.Read.Find(
		&vulnerabilitiesAnalytic, map[string]interface{}{}, enums.MigrationTable)
	if resHorusec.GetErrorExceptNotFound() != nil {
		return false, resHorusec.GetErrorExceptNotFound()
	}

	return len(vulnerabilitiesAnalytic) > 0, nil
}

func (a *AnalyticMigration) saveMigration() error {
	res := a.dbConnectionAnalytic.Write.Create(
		map[string]interface{}{"name": enums.MigrationV1ToV2Name}, enums.MigrationTable)

	return res.GetError()
}

// nolint:funlen // main is not necessary check
func main() {
	analyticMigration := NewAnalyticMigrationV2()

	exists, err := analyticMigration.existsAnalyticMigrations()
	if err != nil {
		logger.LogPanic("{Horusec} Error on get migrations from database", err)
	}

	if exists {
		logger.LogInfo("Migration has been applied with success!")

		return
	}

	// nolint:gomnd // is not necessary create magic number
	allPastAnalysis := analyticMigration.getAllAnalysisWithVulnerabilities(100, 0, []*analysisentities.Analysis{})
	for key, analysis := range allPastAnalysis {
		logger.LogInfo(fmt.Sprintf("Migrating analysis: [%v/%v]", key, len(allPastAnalysis)-1))
		analyticMigration.loggingRegisterBeingMigrated(analysis)
		analyticMigration.migrateAnalysis(analysis)
	}

	analyticMigration.printResults(enums.TotalOfTables * len(allPastAnalysis))

	if err := analyticMigration.saveMigration(); err != nil {
		logger.LogPanic("{Horusec} Error on save migrations in the database", err)
	}

	logger.LogInfo("Migration has been applied with success!")
}
