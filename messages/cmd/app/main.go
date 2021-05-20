package main

import "github.com/ZupIT/horusec-platform/messages/config/providers"

// @title Horusec-Messages
// @description Service responsible for sending emails.
// @termsOfService http://swagger.io/terms/

// @contact.name Horusec
// @contact.url https://github.com/ZupIT/horusec-platform
// @contact.email horusec@zup.com.br

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Horusec-Authorization
func main() {
	router, err := providers.Initialize("8002")
	if err != nil {
		panic(err)
	}

	router.ListenAndServe()
}
