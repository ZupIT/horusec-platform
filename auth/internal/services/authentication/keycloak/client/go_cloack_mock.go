package keycloak

import (
	"context"

	"github.com/Nerzal/gocloak/v7"
	"github.com/stretchr/testify/mock"

	mockUtils "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
)

type GoCloakMock struct {
	mock.Mock
	gocloak.GoCloak
}

func (m *GoCloakMock) LoginOtp(_ context.Context, _, _, _, _, _, _ string) (*gocloak.JWT, error) {
	args := m.MethodCalled("LoginOtp")
	return args.Get(0).(*gocloak.JWT), mockUtils.ReturnNilOrError(args, 1)
}

func (m *GoCloakMock) RetrospectToken(_ context.Context, _, _, _, _ string) (*gocloak.RetrospecTokenResult, error) {
	args := m.MethodCalled("RetrospectToken")
	return args.Get(0).(*gocloak.RetrospecTokenResult), mockUtils.ReturnNilOrError(args, 1)
}

func (m *GoCloakMock) GetUserInfo(_ context.Context, _, _ string) (*gocloak.UserInfo, error) {
	args := m.MethodCalled("GetUserInfo")
	return args.Get(0).(*gocloak.UserInfo), mockUtils.ReturnNilOrError(args, 1)
}
