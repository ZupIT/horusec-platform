package workspace

import (
	"github.com/stretchr/testify/mock"

	mockUtils "github.com/ZupIT/horusec-devkit/pkg/utils/mock"

	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) Create(_ *workspaceEntities.CreateWorkspaceData) (*workspaceEntities.Workspace, error) {
	args := m.MethodCalled("Create")
	return args.Get(0).(*workspaceEntities.Workspace), mockUtils.ReturnNilOrError(args, 1)
}
