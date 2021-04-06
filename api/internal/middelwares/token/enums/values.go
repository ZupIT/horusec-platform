package enums

type CtxKey string

const (
	RepositoryName CtxKey = "repositoryName"
	WorkspaceName  CtxKey = "workspaceName"
	WorkspaceID    CtxKey = "workspaceID"
	RepositoryID   CtxKey = "repositoryID"
)
