package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseconfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	dashboardcontroller "github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
	dashboardrepository "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
	"github.com/briandowns/spinner"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

func main() {
	databaseURI := flag.String("v1-database-uri", "", "")
	flag.Parse()

	coreConn, err := gorm.Open(postgres.Open(*databaseURI), &gorm.Config{})
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
	coreConn.Table("analysis").Order("created_at DESC").Preload("AnalysisVulnerabilities").Preload("AnalysisVulnerabilities.Vulnerability").Find(&analysis)

	migrationCounter := make(map[string][]string)

	for i := range analysis {
		tx := conn.Write.StartTransaction()
		s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
		msg := fmt.Sprintf(
			` repository %s | date %s | vulnerabilities %d`,
			analysis[i].RepositoryName,
			analysis[i].CreatedAt,
			len(analysis[i].AnalysisVulnerabilities),
		)
		s.Suffix = msg
		s.Color("cyan")

		s.Start()

		err = dashboardController.AddVulnerabilitiesByAuthor(&analysis[i])
		err = dashboardController.AddVulnerabilitiesByLanguage(&analysis[i])
		err = dashboardController.AddVulnerabilitiesByRepository(&analysis[i])
		err = dashboardController.AddVulnerabilitiesByTime(&analysis[i])

		if err != nil {
			tx.RollbackTransaction()
			migrationCounter["failed"] = append(migrationCounter["failed"], analysis[i].ID.String())
			s.FinalMSG = fmt.Sprintf(ErrorColor, msg)

			time.Sleep(2 * time.Second)
			s.Stop()
			fmt.Println()
			continue
		}

		migrationCounter["successfuly"] = append(migrationCounter["successfuly"], analysis[i].ID.String())
		tx.CommitTransaction()
		s.FinalMSG = fmt.Sprintf(NoticeColor, msg)

		time.Sleep(2 * time.Second)
		s.Stop()
		fmt.Println()
	}

	fmt.Println()
	fmt.Print("the analytic data migration is finished:")
	fmt.Println()
	fmt.Printf(NoticeColor, fmt.Sprintf(`successfuly: %d`, len(migrationCounter["successfuly"])))
	fmt.Println()
	fmt.Printf(ErrorColor, fmt.Sprintf(`failed: %d`, len(migrationCounter["failed"])))
}
