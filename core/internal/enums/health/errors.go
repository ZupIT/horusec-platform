package health

import "errors"

var ErrorUnhealthyDatabase = errors.New("database connection is not healthy")
var ErrorUnhealthyBroker = errors.New("broker connection is not healthy")
