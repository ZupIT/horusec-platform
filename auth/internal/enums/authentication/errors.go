package authentication

import "errors"

var ErrorWrongEmailOrPassword = errors.New("{AUTHENTICATION} invalid username or password")
var ErrorAccountEmailNotConfirmed = errors.New("{AUTHENTICATION} account email not confirmed")
