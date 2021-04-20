package main

import (
	"github.com/ZupIT/horusec-platform/analytic/config/providers_broker"
	"github.com/ZupIT/horusec-platform/analytic/internal/enums"
)

func main() {
	router, err := providersbroker.Initialize(enums.DefaultPort)
	if err != nil {
		panic(err)
	}

	router.ListenAndServe()
}
