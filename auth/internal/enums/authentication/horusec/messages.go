package horusec

const (
	ErrorFailedToGetAccountIDFromToken = "{HORUSEC AUTH} failed to get account id from jwt token" //nolint:gosec, lll // false positive
	ErrorFailedToGetWorkspaceRole      = "{HORUSEC AUTH} failed to get account role for workspace"
	ErrorFailedToGetAccountAppAdmin    = "{HORUSEC AUTH} failed to get account to check for application admin"
)
