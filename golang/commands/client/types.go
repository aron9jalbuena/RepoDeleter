package client

import (
	"context"
	"deleter/vars"
	"fmt"

	"github.com/google/go-github/v53/github"
)

type Client struct {
	*github.Client
}

func (client *Client) ListRepositories(ctx context.Context) error {
	repos, _, err := client.Repositories.List(ctx, "", &github.RepositoryListOptions{Affiliation: "owner"})
	if err != nil {
		return fmt.Errorf("list repostiory failed with: %s", err)
	}

	err = writeRepositoriesToFile(vars.RepoFileName, repos)
	if err != nil {
		return fmt.Errorf("writing repos to %s failed with: %s", vars.RepoFileName, err)
	}

	return nil
}
