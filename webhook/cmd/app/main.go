package main

import (
	"github.com/ZupIT/horusec-platform/webhook/config/providers"
	"github.com/ZupIT/horusec-platform/webhook/internal/enums"
)

// @title Horusec-Webhook
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
