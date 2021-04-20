package account

import "errors"

var ErrorEmailAlreadyInUse = errors.New("{ACCOUNT} email already in use")
var ErrorUsernameAlreadyInUse = errors.New("{ACCOUNT} username already in use")
