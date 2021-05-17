package dashboard

type ByAuthor struct {
	Author string `json:"author"`
	*BySeverities
}
