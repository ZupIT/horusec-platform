package enums

import "errors"

var ErrorWorkspaceNotSelected = errors.New("{HORUSEC} workspace not found for token sent")
var ErrorRepositoryNotSelected = errors.New("{HORUSEC} repositoryName cannot be empty " +
	"in the body for a workspace token")
