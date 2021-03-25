package transport

import (
	"cmd/api-users-grpc/internal/endpoints"
	"cmd/api-users-grpc/pkg/proto"
	"context"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	find gt.Handler
}

func NewGRPCServer(endpoints endpoints.Endpoints) proto.UserServiceServer {
	return &gRPCServer{
		find: gt.NewServer(
			endpoints.FindGithubUsers,
			decodeFindGithubUsersRequest,
			encodeFindGithubUsersResponse,
		),
	}
}

func (s *gRPCServer) FindGithubUser(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	_, resp, err := s.find.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.UserResponse), nil
}

func decodeFindGithubUsersRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*proto.UserRequest)
	return endpoints.FindGithubUsersRequest{Username: req.Username}, nil
}

func encodeFindGithubUsersResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.FindGithubUsersResponse)
	return resp.ToGrpcResponse(), nil
}