package workspace

import (
	mockUtils "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) Create(_ *workspaceEntities.Data) (*workspaceEntities.Workspace, error) {
	args := m.MethodCalled("Create")
	return args.Get(0).(*workspaceEntities.Workspace), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) Get(_ *workspaceEntities.Data) (*workspaceEntities.Response, error) {
	args := m.MethodCalled("Get")
	return args.Get(0).(*workspaceEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) Update(_ *workspaceEntities.Data) (*workspaceEntities.Workspace, error) {
	args := m.MethodCalled("Update")
	return args.Get(0).(*workspaceEntities.Workspace), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) Delete(_ uuid.UUID) error {
	args := m.MethodCalled("Delete")
	return mockUtils.ReturnNilOrError(args, 0)
}
