package health

import "errors"

var ErrorUnhealthyBroker = errors.New("broker is not healthy")
var ErrorUnhealthyMailer = errors.New("mailer is not healthy")
