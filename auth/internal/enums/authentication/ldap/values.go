package ldap

const (
	EnvLdapHost               = "HORUSEC_LDAP_HOST"
	EnvLdapPort               = "HORUSEC_LDAP_PORT"
	EnvLdapBase               = "HORUSEC_LDAP_BASE"
	EnvLdapBindDn             = "HORUSEC_LDAP_BINDDN"
	EnvLdapBindPassword       = "HORUSEC_LDAP_BINDPASSWORD"
	EnvLdapUseSSL             = "HORUSEC_LDAP_USESSL"
	EnvLdapSkipTLS            = "HORUSEC_LDAP_SKIP_TLS"
	EnvLdapInsecureSkipVerify = "HORUSEC_LDAP_INSECURE_SKIP_VERIFY"
	EnvLdapUserFilter         = "HORUSEC_LDAP_USERFILTER"
	EnvLdapAdminGroup         = "HORUSEC_LDAP_ADMIN_GROUP"
	DefaultLdapUserFilter     = "(sAMAccountName=%s)"
)
