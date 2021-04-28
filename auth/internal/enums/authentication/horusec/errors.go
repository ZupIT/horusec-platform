package horusec

import "errors"

var ErrorWrongEmailOrPassword = errors.New("{HORUSEC AUTH} invalid username or password")
var ErrorApplicationAdminNotEnabled = errors.New("{HORUSEC AUTH} application admin not enabled in service config")
var ErrorAccountEmailNotConfirmed = errors.New("{HORUSEC AUTH} account email not confirmed")
