package dashboard

import (
	"time"
)

type ByTime struct {
	Time time.Time `json:"time"`
	*BySeverities
}
