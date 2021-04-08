package main

import (
	"github.com/ZupIT/horusec-platform/api/config/providers"
	"github.com/ZupIT/horusec-platform/api/internal/router/enums"
)

// @title Horusec-API
// @description Service responsible for workspace, repositories and token operations.
// @termsOfService http://swagger.io/terms/

// @contact.name Horusec
// @contact.url https://github.com/ZupIT/horusec-platform
// @contact.email horusec@zup.com.br

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Horusec-Authorization
func main() {
	router, err := providers.Initialize(enums.DefaultPort)
	if err != nil {
		panic(err)
	}

	router.ListenAndServe()
}
