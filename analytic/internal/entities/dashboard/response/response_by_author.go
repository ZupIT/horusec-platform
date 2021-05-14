package response

type ByAuthor struct {
	Author string `json:"author"`
	BySeverities
}
