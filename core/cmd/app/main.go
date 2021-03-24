package main

import (
	"github.com/ZupIT/horusec-platform/core/config/providers"
)

func main() {
	router, err := providers.Initialize("8003")
	if err != nil {
		panic(err)
	}

	router.ListenAndServe()
}
