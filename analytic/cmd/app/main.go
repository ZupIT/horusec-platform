package main

import (
	"github.com/ZupIT/horusec-platform/analytic/config/providers"
	"github.com/ZupIT/horusec-platform/analytic/internal/router/enums"
)

// @title Horusec-Analytic
// @description Service responsible for managing vulnerabilities.
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
