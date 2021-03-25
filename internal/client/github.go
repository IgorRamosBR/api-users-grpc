package client

import (
	"context"
	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
) // with go modules enabled (GO111MODULE=on or outside GOPATH)

func CreateGithubClient() *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "9de9630ca6223575cd1146e88481a4572a5b0fd9"},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}
