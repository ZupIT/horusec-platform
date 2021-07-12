// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	healthGRPC "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/health"
	"github.com/ZupIT/horusec-devkit/pkg/utils/env"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"

	"github.com/ZupIT/horusec-platform/auth/config/grpc/enums"
	authHandler "github.com/ZupIT/horusec-platform/auth/internal/handlers/authentication"
)

const DefaultGRPCPort = 8007

type IAuthGRPCServer interface {
	ListenAndServeGRPCServer()
}

type AuthGRPCServer struct {
	Port        int
	GRPCServer  *grpc.Server
	authHandler *authHandler.Handler
}

func NewAuthGRPCServer(handlerAuth *authHandler.Handler) IAuthGRPCServer {
	server := &AuthGRPCServer{
		Port:        env.GetEnvOrDefaultInt(enums.EnvGRPCPort, DefaultGRPCPort),
		authHandler: handlerAuth,
	}

	return server.setup()
}

func (a *AuthGRPCServer) setup() *AuthGRPCServer {
	if a.useCredentials() {
		return a.serverWithCerts()
	}

	return a.setupWithoutCerts()
}

func (a *AuthGRPCServer) serverWithCerts() *AuthGRPCServer {
	return a.setGRPCServerAndRegisterServices(grpc.NewServer(grpc.Creds(a.getCredentials())))
}

func (a *AuthGRPCServer) setupWithoutCerts() *AuthGRPCServer {
	return a.setGRPCServerAndRegisterServices(grpc.NewServer())
}

func (a *AuthGRPCServer) getCredentials() credentials.TransportCredentials {
	cred, err := credentials.NewServerTLSFromFile(env.GetEnvOrDefault(enums.EnvGRPCCertPath, ""),
		env.GetEnvOrDefault(enums.EnvGRPCKeyPath, ""))
	if err != nil {
		logger.LogPanic(enums.MessageFailedToGetGRPCCredentials, err)
	}

	return cred
}

func (a *AuthGRPCServer) useCredentials() bool {
	return env.GetEnvOrDefaultBool(enums.EnvGrpcUseCerts, false)
}

func (a *AuthGRPCServer) setGRPCServerAndRegisterServices(grpcServer *grpc.Server) *AuthGRPCServer {
	a.GRPCServer = grpcServer
	a.registerServices()

	return a
}

func (a *AuthGRPCServer) getNetListener() net.Listener {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.Port))
	if err != nil {
		logger.LogPanic(enums.MessageFailedToGetTCPListener, err)
	}

	return listener
}

func (a *AuthGRPCServer) registerServices() {
	healthGRPC.RegisterHealthServer(a.GRPCServer, health.NewHealthCheckGrpcServer())
	proto.RegisterAuthServiceServer(a.GRPCServer, a.authHandler)
}

func (a *AuthGRPCServer) listeningMessage() {
	logger.LogInfo(fmt.Sprintf(enums.MessageGRPCServerListening, a.Port))
}

func (a *AuthGRPCServer) ListenAndServeGRPCServer() {
	a.listeningMessage()

	if err := a.GRPCServer.Serve(a.getNetListener()); err != nil {
		logger.LogPanic(enums.MessageGRPCServerFailedToStart, err)
	}
}
