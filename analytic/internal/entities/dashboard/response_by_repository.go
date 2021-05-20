package dashboard

type ByRepository struct {
	RepositoryName string `json:"repositoryName"`
	*BySeverities
}
