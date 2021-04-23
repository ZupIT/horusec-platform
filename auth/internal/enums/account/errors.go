package account

import "errors"

var ErrorEmailAlreadyInUse = errors.New("{ACCOUNT} email already in use")
var ErrorUsernameAlreadyInUse = errors.New("{ACCOUNT} username already in use")
var ErrorIncorrectRetrievePasswordCode = errors.New("{ACCOUNT} wrong or invalid retrieve password")
var ErrorPasswordEqualPrevious = errors.New("{ACCOUNT} the new password cannot be the same as the previous one")
var ErrorInvalidOrExpiredToken = errors.New("{ACCOUNT} invalid or expired refresh token")
