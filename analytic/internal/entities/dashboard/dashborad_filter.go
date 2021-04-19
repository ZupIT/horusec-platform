package dashboard

import (
	"time"

	"github.com/google/uuid"
)

type FilterDashboard struct {
	RepositoryID    uuid.UUID
	WorkspaceID     uuid.UUID
	InitialDate     time.Time
	FinalDate       time.Time
	Page            int
	Size            int
	AdditionalQuery FilterDashboardAdditionalQuery
}

type FilterDashboardAdditionalQuery struct {
	Query string
	Args  []interface{}
}
