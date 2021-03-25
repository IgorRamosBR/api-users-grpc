package endpoints

import (
	"cmd/api-users-grpc/internal/dto"
	"cmd/api-users-grpc/pkg/proto"
)

type FindGithubUsersResponse struct {
	Username   string
	Name       string
	Email      string
	Company    string
	AvatarUrl  string
	ProfileUrl string
	ReposUrl   string
	Followers  int
	Following  int
}

func (r FindGithubUsersResponse) ToGrpcResponse() *proto.UserResponse {
	return &proto.UserResponse{
		Username:   r.Username,
		Name:       r.Name,
		Email:      r.Email,
		Company:    r.Company,
		AvatarUrl:  r.AvatarUrl,
		ProfileUrl: r.ProfileUrl,
		ReposUrl:   r.ReposUrl,
		Followers:  int64(r.Followers),
		Following:  int64(r.Following),
	}
}

func ToResponse(u dto.UserDto) FindGithubUsersResponse {
	return FindGithubUsersResponse{
		Username:   u.Username,
		Name:       u.Name,
		Email:      u.Name,
		Company:    u.Company,
		AvatarUrl:  u.AvatarUrl,
		ProfileUrl: u.HtmlUrl,
		ReposUrl:   u.ReposUrl,
		Followers:  u.Followers,
		Following:  u.Following,
	}
}
