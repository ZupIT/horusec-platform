package authentication

import "time"

const (
	TokenDuration             = time.Hour * 2
	TokenCheckExpiredDuration = time.Minute * 10
	TableWorkspaces           = "workspaces"
	TableRepositories         = "repositories"
)
