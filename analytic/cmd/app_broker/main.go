package main

import (
	"github.com/ZupIT/horusec-platform/analytic/config/providers"
	"github.com/ZupIT/horusec-platform/analytic/internal/enums"
)

func main() {
	router, err := providers.Initialize(enums.DefaultPortBroker)
	if err != nil {
		panic(err)
	}

	router.ListenAndServe()
}
