package main

import (
	"cmd/api-users-grpc/internal/client"
	"cmd/api-users-grpc/internal/endpoints"
	"cmd/api-users-grpc/internal/logging"
	"cmd/api-users-grpc/internal/services"
	"cmd/api-users-grpc/internal/transport"
	"cmd/api-users-grpc/pkg/proto"
	"fmt"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
	"net"
)

func main() {
	logger := logging.GetLogger()

	githubClient := client.CreateGithubClient()
	userService := services.NewUsersService(githubClient)
	userEndpoint := endpoints.MakeEndpoints(userService)
	grpcServer := transport.NewGRPCServer(userEndpoint)

	_ = level.Info(logger).Log("status", "listening", "port", "9000")

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		_ = level.Error(logger).Log("failed to listen: %v", err)
		panic(err)
	}

	baseServer := grpc.NewServer()
	proto.RegisterUserServiceServer(baseServer, grpcServer)
	_ = level.Info(logger).Log("msg", "Server started successfully 🚀")

	_ = level.Error(logger).Log(baseServer.Serve(listen))
}
