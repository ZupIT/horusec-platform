package dashboard

import "github.com/ZupIT/horusec-devkit/pkg/services/database"

type IRepoDashboard interface{}

type RepoDashboard struct {
	dbRead  database.IDatabaseRead
	dbWrite database.IDatabaseWrite
}

func NewRepoDashboard(connection *database.Connection) IRepoDashboard {
	return &RepoDashboard{
		dbRead:  connection.Read,
		dbWrite: connection.Write,
	}
}
