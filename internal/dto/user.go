package dto

import (
	"github.com/google/go-github/v33/github"
)

type UserDto struct {
	Username        string
	AvatarUrl       string
	HtmlUrl         string
	ReposUrl        string
	Name            string
	Email           string
	Company         string
	Bio             string
	ApiUrl          string
	TwitterUsername string
	Followers       int
	Following       int
}

func ToDto(user *github.User) UserDto {
	return UserDto{
		Username:        extractString(user.Login),
		AvatarUrl:       extractString(user.AvatarURL),
		HtmlUrl:         extractString(user.HTMLURL),
		ReposUrl:        extractString(user.ReposURL),
		Name:            extractString(user.Name),
		Email:           extractString(user.Email),
		Company:         extractString(user.Company),
		Bio:             extractString(user.Bio),
		ApiUrl:          extractString(user.URL),
		TwitterUsername: extractString(user.TwitterUsername),
		Followers:       extractInt(user.Followers),
		Following:       extractInt(user.Following),
	}
}

func extractString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func extractInt(i *int) int {
	if i != nil {
		return *i
	}
	return 0
}