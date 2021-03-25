package endpoints

import (
	"cmd/api-users-grpc/internal/logging"
	"cmd/api-users-grpc/internal/services"
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log/level"
)

type Endpoints struct {
	FindGithubUsers endpoint.Endpoint
}

func MakeEndpoints(service services.UsersService) Endpoints {
	return Endpoints{
		FindGithubUsers: makeFindGithubUsersEndpoint(service),
	}
}

func makeFindGithubUsersEndpoint(service services.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		logger := logging.GetLogger()
		req := request.(FindGithubUsersRequest)
		user, err := service.FindGithubUsers(ctx, req.Username)
		if err != nil {
			level.Error(logger).Log(err.Error())
		}
		return ToResponse(user), nil
	}
}
