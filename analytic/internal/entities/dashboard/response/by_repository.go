package response

type ByRepository struct {
	RepositoryName string `json:"repositoryName"`
	*BySeverities
}
