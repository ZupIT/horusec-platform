package enums

import "errors"

var (
	ErrorWebhookDuplicate = errors.New("{HORUSEC} webhook already exists to repository selected")
	ErrorWrongWorkspaceID = errors.New("{HORUSEC} workspaceID is not valid uuid")
	ErrorWrongWebhookID   = errors.New("{HORUSEC} webhookID is not valid uuid")
)
