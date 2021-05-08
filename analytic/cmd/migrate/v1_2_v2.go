package main

import (
	"flag"
	"fmt"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	databaseURI := flag.String("v1-database-uri", "", "")
	flag.Parse()

	conn, _ := gorm.Open(postgres.Open(string(*databaseURI)), &gorm.Config{})

	analysis := []analysis.Analysis{}
	conn.Table("analysis").Preload("AnalysisVulnerabilities").Find(&analysis)

	for _, analyse := range analysis {
		fmt.Println(analyse)
	}

}
