package response

import (
	"time"
)

type ByTime struct {
	Time time.Time `json:"time"`
	*BySeverities
}
