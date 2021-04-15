package keycloak

const (
	EnvHorusecKeycloakBasePath     = "HORUSEC_KEYCLOAK_BASE_PATH"
	EnvHorusecKeycloakClientID     = "HORUSEC_KEYCLOAK_CLIENT_ID"
	EnvHorusecKeycloakClientSecret = "HORUSEC_KEYCLOAK_CLIENT_SECRET" //nolint:gosec // false positive
	EnvHorusecKeycloakRealm        = "HORUSEC_KEYCLOAK_REALM"
	EnvHorusecKeycloakTOPT         = "HORUSEC_KEYCLOAK_TOTP"
)
