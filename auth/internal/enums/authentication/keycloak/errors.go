package keycloak

import "errors"

var ErrorKeycloakMissingUsernameOrSub = errors.New("{KEYCLOAK AUTH} keycloak user info missing username or email")
