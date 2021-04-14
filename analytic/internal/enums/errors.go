package enums

import "errors"

var (
	ErrorWrongRepositoryID = errors.New("{HORUSEC} repositoryID is not valid uuid")
	ErrorWrongWorkspaceID  = errors.New("{HORUSEC} workspaceID is not valid uuid")
	ErrorWrongInitialDate  = errors.New("{HORUSEC} initialDate is not valid date")
	ErrorWrongFinalDate    = errors.New("{HORUSEC} finalDate is not valid date")
)
