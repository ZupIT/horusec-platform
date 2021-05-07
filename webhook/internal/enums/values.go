package enums

const (
	DefaultPort   = "8004"
	BaseRouter    = "/webhook"
	HealthRouter  = BaseRouter + "/health"
	WebhookRouter = BaseRouter + "/webhook/{workspaceID}"
)
