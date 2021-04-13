package main

import "github.com/ZupIT/horusec-platform/auth/config/providers"

func main() {
	router, err := providers.Initialize("8006")
	if err != nil {
		panic(err)
	}

	router.ListenAndServe()
}
