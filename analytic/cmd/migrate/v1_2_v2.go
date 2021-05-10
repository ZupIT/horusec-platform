package main

import (
	"flag"
	"fmt"

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

	coreConn, err := gorm.Open(postgres.Open(string(*databaseURI)), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
		return
	}

	databaseConfig := databaseconfig.NewDatabaseConfig()
	conn, _ := database.NewDatabaseReadAndWrite(databaseConfig)
	if err != nil {
		fmt.Print(err)
		return
	}

	dashboardRepository := dashboardrepository.NewRepoDashboard(conn)
	dashboardController := dashboardcontroller.NewControllerDashboardWrite(dashboardRepository)

	analysis := []analysis.Analysis{}
	coreConn.Table("analysis").Order("CreatedAt").Preload("AnalysisVulnerabilities").Preload("AnalysisVulnerabilities.Vulnerability").Find(&analysis)

	for _, analyse := range analysis {
		conn.Write.StartTransaction()
		err = dashboardController.AddVulnerabilitiesByAuthor(&analyse)
		err = dashboardController.AddVulnerabilitiesByLanguage(&analyse)
		err = dashboardController.AddVulnerabilitiesByRepository(&analyse)
		err = dashboardController.AddVulnerabilitiesByTime(&analyse)

		if err != nil {
			conn.Write.RollbackTransaction()
			return
		}

		conn.Write.CommitTransaction()
	}
}
