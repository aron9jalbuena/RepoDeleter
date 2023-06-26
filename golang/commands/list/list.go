package list

import (
	"deleter/commands/client"

	"github.com/urfave/cli/v2"
)

func List() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "List all repo's you own into a repos.txt file",
		Action: func(ctx *cli.Context) error {
			return list(ctx)
		},
	}
}

func list(ctx *cli.Context) error {
	gClient := client.New(ctx.Context)
	return gClient.ListRepositories(ctx.Context)
}
