package dashboard

import (
	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
)

type ByLanguage struct {
	Language languages.Language `json:"language"`
	*BySeverities
}
