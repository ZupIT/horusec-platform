package main

import "github.com/ZupIT/horusec-platform/auth/config/providers"

// @title Horusec-Auth
// @description Service responsible for authentication and account operations.
// @termsOfService http://swagger.io/terms/

// @contact.name Horusec
// @contact.url https://github.com/ZupIT/horusec-platform
// @contact.email horusec@zup.com.br

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Horusec-Authorization
func main() {
	router, err := providers.Initialize("8006")
	if err != nil {
		panic(err)
	}

	router.ListenAndServe()
}
