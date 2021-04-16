package authentication

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	accountEnums "github.com/ZupIT/horusec-devkit/pkg/enums/account"
	mockUtils "github.com/ZupIT/horusec-devkit/pkg/utils/mock"

	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) GetWorkspaceGroups(_ uuid.UUID) (*authEntities.AuthzGroups, error) {
	args := m.MethodCalled("GetWorkspaceGroups")
	return args.Get(0).(*authEntities.AuthzGroups), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetRepositoryGroups(_ uuid.UUID) (*authEntities.AuthzGroups, error) {
	args := m.MethodCalled("GetRepositoryGroups")
	return args.Get(0).(*authEntities.AuthzGroups), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetWorkspaceRole(_, _ uuid.UUID) (accountEnums.Role, error) {
	args := m.MethodCalled("GetWorkspaceRole")
	return args.Get(0).(accountEnums.Role), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetRepositoryRole(_, _ uuid.UUID) (accountEnums.Role, error) {
	args := m.MethodCalled("GetRepositoryRole")
	return args.Get(0).(accountEnums.Role), mockUtils.ReturnNilOrError(args, 1)
}
