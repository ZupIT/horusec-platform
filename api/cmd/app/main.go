package main

import (
	"github.com/ZupIT/horusec-platform/api/config/providers"
)

func main() {
	router, err := providers.Initialize("8001")
	if err != nil {
		panic(err)
	}

	router.ListenAndServe()
}
