package main

import (
	"flag"
	"fmt"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// brokerConfig := brokerconfig.NewBrokerConfig()
	// broker, err := broker.NewBroker(brokerConfig)

	databaseURI := flag.String("v1-database-uri", "", "")
	flag.Parse()

	conn, _ := gorm.Open(postgres.Open(string(*databaseURI)), &gorm.Config{})

	analysis := &analysis.Analysis{}
	conn.Table("analysis").Preload("AnalysisVulnerabilities").Find(analysis)

	fmt.Printf(`%+v\n`, analysis)
}
