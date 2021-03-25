package services

import (
	"cmd/api-users-grpc/internal/dto"
	"cmd/api-users-grpc/internal/logging"
	"context"
	"github.com/go-kit/kit/log/level"
	"github.com/google/go-github/v33/github"
)

type UsersService interface {
	FindGithubUsers(ctx context.Context, username string) (userDto dto.UserDto, err error)
}

type UsersServiceImpl struct {
	GithubClient *github.Client
}

func NewUsersService(client *github.Client) UsersService {
	return UsersServiceImpl{
		GithubClient: client,
	}
}

func (s UsersServiceImpl) FindGithubUsers(ctx context.Context, username string) (userDto dto.UserDto, err error) {
	logger := logging.GetLogger()
	get, _, err := s.GithubClient.Users.Get(ctx, username)
	if err != nil {
		_ = level.Error(logger).Log(err.Error())
		return dto.UserDto{}, err
	}
	user := dto.ToDto(get)
	level.Info(logger).Log(user.Username)
	return user, nil
}
