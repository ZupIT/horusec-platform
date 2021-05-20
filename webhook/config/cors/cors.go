package cors

import "github.com/go-chi/cors"

func NewCorsConfig() *cors.Options {
	return &cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "headers", "X-Horusec-Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}
}
