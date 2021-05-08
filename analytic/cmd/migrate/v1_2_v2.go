package main

import (
	"flag"
	"fmt"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/enums/exchange"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	brokerconfig "github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	brokerConfig := brokerconfig.NewBrokerConfig()
	broker, _ := broker.NewBroker(brokerConfig)

	databaseURI := flag.String("v1-database-uri", "", "")
	flag.Parse()

	conn, _ := gorm.Open(postgres.Open(string(*databaseURI)), &gorm.Config{})

	analysis := &[]analysis.Analysis{}
	conn.Table("analysis").Preload("AnalysisVulnerabilities").Find(analysis)

	for _, analyse := range *analysis {
		broker.Publish("", exchange.NewAnalysis.ToString(), exchange.Fanout.ToString(), analyse.ToBytes())
		fmt.Printf(`%+v\n`, analyse)
	}

}
