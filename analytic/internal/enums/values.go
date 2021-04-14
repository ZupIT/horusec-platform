package enums

const (
	DefaultPort               = "8003"
	BaseRouter                = "/analytic"
	HealthRouter              = BaseRouter + "/health"
	DashboardWorkspaceRouter  = BaseRouter + "/dashboard/{workspaceID}"
	DashboardRepositoryRouter = DashboardWorkspaceRouter + "/{repositoryID}"
)
