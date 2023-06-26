package client

import (
	"context"
	"deleter/vars"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

func New(ctx context.Context) *Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: vars.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	gclient := github.NewClient(tc)

	return &Client{
		Client: gclient,
	}
}
