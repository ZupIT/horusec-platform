package ldap

import "errors"

var ErrorEmptyBindDNOrBindPassword = errors.New("{LDAP} empty bind dn or bind password")
var ErrorUserDoesNotExist = errors.New("{LDAP} user does not exist")
var ErrorTooManyEntries = errors.New("{LDAP} too many entries returned")
var ErrorLdapUnauthorized = errors.New("{LDAP} it was not possible to authorize your login with these credentials")
var ErrorLdapApplicationAdminGroupNotSet = errors.New("{LDAP} horusec application admin group env not set")
