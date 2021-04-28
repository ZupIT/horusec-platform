package dashboard

import (
	"time"

	"github.com/google/uuid"
)

type FilterDashboard struct {
	RepositoryID uuid.UUID
	WorkspaceID  uuid.UUID
	StartTime    time.Time
	EndTime      time.Time
	Page         int
	Size         int
}
