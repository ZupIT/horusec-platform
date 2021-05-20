package workspace

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	mockUtils "github.com/ZupIT/horusec-devkit/pkg/utils/mock"

	roleEntities "github.com/ZupIT/horusec-platform/core/internal/entities/role"
	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) ListWorkspacesAuthTypeHorusec(_ uuid.UUID) (*[]workspaceEntities.Response, error) {
	args := m.MethodCalled("ListWorkspacesAuthTypeHorusec")
	return args.Get(0).(*[]workspaceEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) ListWorkspacesAuthTypeLdap(_ []string) (*[]workspaceEntities.Response, error) {
	args := m.MethodCalled("ListWorkspacesAuthTypeLdap")
	return args.Get(0).(*[]workspaceEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetWorkspace(_ uuid.UUID) (*workspaceEntities.Workspace, error) {
	args := m.MethodCalled("GetWorkspace")
	return args.Get(0).(*workspaceEntities.Workspace), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetAccountWorkspace(_, _ uuid.UUID) (*workspaceEntities.AccountWorkspace, error) {
	args := m.MethodCalled("GetAccountWorkspace")
	return args.Get(0).(*workspaceEntities.AccountWorkspace), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) ListAllWorkspaceUsers(_ uuid.UUID) (*[]roleEntities.Response, error) {
	args := m.MethodCalled("ListAllWorkspaceUsers")
	return args.Get(0).(*[]roleEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) ListWorkspacesApplicationAdmin() (*[]workspaceEntities.Response, error) {
	args := m.MethodCalled("ListWorkspacesApplicationAdmin")
	return args.Get(0).(*[]workspaceEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}
