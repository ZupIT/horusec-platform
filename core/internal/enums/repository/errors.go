package repository

import "errors"

var ErrorRepositoryNameAlreadyInUse = errors.New("{CORE_REPOSITORY} repository name already in use")
var ErrorUserDoesNotBelongToWorkspace = errors.New("{CORE_REPOSITORY} this user does not belong to this workspace")
