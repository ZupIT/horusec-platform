package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
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

	summary := make(map[string][]string)
	var logError string

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
		if err != nil {
			summary[analysis[i].ID.String()] = append(summary[analysis[i].ID.String()], err.Error())
		}

		err = dashboardController.AddVulnerabilitiesByLanguage(&analysis[i])
		if err != nil {
			summary[analysis[i].ID.String()] = append(summary[analysis[i].ID.String()], err.Error())
		}

		err = dashboardController.AddVulnerabilitiesByRepository(&analysis[i])
		if err != nil {
			summary[analysis[i].ID.String()] = append(summary[analysis[i].ID.String()], err.Error())
		}

		err = dashboardController.AddVulnerabilitiesByTime(&analysis[i])
		if err != nil {
			summary[analysis[i].ID.String()] = append(summary[analysis[i].ID.String()], err.Error())
		}

		if err != nil {
			tx.RollbackTransaction()

			summary["failed"] = append(summary["failed"], analysis[i].ID.String())

			logError += fmt.Sprintf("%s\n\n%s\n\n\n\n", msg, strings.Join(summary[analysis[i].ID.String()], "\n"))

			s.FinalMSG = fmt.Sprintf(ErrorColor, msg)
			s.Stop()
			fmt.Println()
			continue
		}

		summary["successfuly"] = append(summary["successfuly"], analysis[i].ID.String())
		tx.CommitTransaction()
		s.FinalMSG = fmt.Sprintf(NoticeColor, msg)

		s.Stop()
		fmt.Println()
	}

	if len(summary["failed"]) > 0 {
		f, _ := os.Create("/tmp/v1-2v2-horusec-analytic-log-error")
		w := bufio.NewWriter(f)
		w.WriteString(logError)
		w.Flush()
	}

	fmt.Println()
	fmt.Print("the analytic data migration is finished:")
	fmt.Println()
	fmt.Printf(NoticeColor, fmt.Sprintf(`successfuly: %d`, len(summary["successfuly"])))
	fmt.Println()
	fmt.Printf(ErrorColor, fmt.Sprintf(`failed: %d`, len(summary["failed"])))
}
