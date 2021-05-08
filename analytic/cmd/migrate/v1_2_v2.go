package main

import (
	"flag"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseconfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	dashboardcontroller "github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
	dashboardrepository "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	databaseURI := flag.String("v1-database-uri", "", "")
	flag.Parse()

	coreConn, _ := gorm.Open(postgres.Open(string(*databaseURI)), &gorm.Config{})

	databaseConfig := databaseconfig.NewDatabaseConfig()
	conn, _ := database.NewDatabaseReadAndWrite(databaseConfig)
	dashboardRepository := dashboardrepository.NewRepoDashboard(conn)
	dashboardController := dashboardcontroller.NewControllerDashboardWrite(dashboardRepository)

	analysis := []analysis.Analysis{}
	coreConn.Table("analysis").Preload("AnalysisVulnerabilities").Preload("AnalysisVulnerabilities.Vulnerability").Find(&analysis)

	for _, analyse := range analysis {
		_ = dashboardController.AddVulnerabilitiesByAuthor(&analyse)
		_ = dashboardController.AddVulnerabilitiesByLanguage(&analyse)
		_ = dashboardController.AddVulnerabilitiesByRepository(&analyse)
		_ = dashboardController.AddVulnerabilitiesByTime(&analyse)
	}
}
