package enums

const (
	DefaultPortRest          = "8003"
	DefaultPortBroker        = "8004"
	BaseRouter               = "/analytic"
	HealthRouter             = BaseRouter + "/health"
	DashboardWorkspaceRouter = BaseRouter + "/dashboard/{workspaceID}"
)
