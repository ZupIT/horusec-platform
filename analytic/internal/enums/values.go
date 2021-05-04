package enums

const (
	DefaultPort              = "8005"
	BaseRouter               = "/analytic"
	HealthRouter             = BaseRouter + "/health"
	DashboardWorkspaceRouter = BaseRouter + "/dashboard/{workspaceID}"
)
