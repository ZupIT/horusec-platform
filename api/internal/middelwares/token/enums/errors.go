package enums

import "errors"

var ErrorTokenExpired = errors.New("this authorization token has expired, please renew it")
