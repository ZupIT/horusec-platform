package enums

const (
	EnvAuthURL                  = "HORUSEC_AUTH_URL"
	EnvAuthType                 = "HORUSEC_AUTH_TYPE"
	EnvDisableBroker            = "HORUSEC_DISABLE_BROKER"
	EnvEnableApplicationAdmin   = "HORUSEC_ENABLE_APPLICATION_ADMIN"
	ApplicationAdminDefaultData = "{\"username\": \"horusec-admin\", \"email\":\"horusec-admin@example.com\", \"password\":\"Devpass0*\"}"
	EnvApplicationAdminData     = "HORUSEC_APPLICATION_ADMIN_DATA"
	EnvEnableDefaultUser        = "HORUSEC_ENABLE_DEFAULT_USER"
	EnvDefaultUserData          = "HORUSEC_DEFAULT_USER_DATA"
	DefaultUserData             = "{\"username\": \"dev\", \"email\":\"dev@example.com\", \"password\":\"Devpass0*\"}"
)
