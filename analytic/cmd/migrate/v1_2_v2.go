package main

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
)

func main() {
	brokerConfig := config.NewBrokerConfig()
	broker, err := broker.NewBroker(brokerConfig)

	if err != nil {
		return
	}
}
