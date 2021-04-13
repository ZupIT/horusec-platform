package token

import "errors"

var ErrorInvalidTokenExpiresAt = errors.New("{TOKEN} expires at cannot be a past date")
