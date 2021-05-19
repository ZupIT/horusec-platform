package enums

const (
	MessageFailedToConnectToAnalyticDatabase = "failed to connect to analytic database, please check your uri informed" +
		" in HORUSEC_DATABASE_ANALYTIC_SQL_URI env variable"
	MessageFailedToConnectToHorusecDatabase = "failed to connect to horusec default database, please check your uri " +
		"informed in HORUSEC_DATABASE_HORUSEC_SQL_URI env variable"
	MessageFailedToGetAllAnalysis = "something went wrong while getting all your past analysis please check " +
		"your informed in HORUSEC_DATABASE_HORUSEC_SQL_URI env variable"
	MessageRegisterBeingMigrated = "Stating to migrate analysis with: WorkspaceID -> %s | Workspace Name ->" +
		" %s | Repository ID %s | Repository Name %s | Created At -> %v | Total Of Vulnerabilities -> %d"
)
