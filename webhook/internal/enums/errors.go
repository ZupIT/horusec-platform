package enums

import "errors"

var (
	ErrorWrongWorkspaceID = errors.New("{HORUSEC} workspaceID is not valid uuid")
	ErrorWrongWebhookID   = errors.New("{HORUSEC} webhookID is not valid uuid")
)
