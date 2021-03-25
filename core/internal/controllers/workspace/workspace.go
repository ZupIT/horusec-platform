package workspace

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
)

type IController interface {
	Test()
}

type Controller struct {
	broker        broker.IBroker
	databaseRead  database.IDatabaseRead
	databaseWrite database.IDatabaseWrite
}

// nolint
func NewWorkspaceController(broker broker.IBroker, databaseConnection *database.Connection) IController {
	return &Controller{
		broker:        broker,
		databaseRead:  databaseConnection.Read,
		databaseWrite: databaseConnection.Write,
	}
}

func (c *Controller) Test() {
	print("\n")
	print("test")
	print("\n")
}
