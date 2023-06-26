package delete

import (
	"bufio"
	"fmt"
	"os"

	"deleter/commands/client"
	"deleter/vars"

	"github.com/urfave/cli/v2"
)

func Delete() *cli.Command {
	return &cli.Command{
		Name:  "delete",
		Usage: "Deletes any repos that are in repos.txt",
		Action: func(ctx *cli.Context) error {
			return delete(ctx)
		},
	}
}

func delete(ctx *cli.Context) error {
	repos, err := readRepoFile()
	if err != nil {
		return err
	}

	gClient := client.New(ctx.Context)

	for _, name := range repos {
		_, err := gClient.Repositories.Delete(ctx.Context, vars.User, name)
		if err != nil {
			return fmt.Errorf("deleting repo %s failed with: %s", name, err)
		}

	}

	return nil
}

func readRepoFile() ([]string, error) {
	file, err := os.Open(vars.RepoFileName)
	if err != nil {
		return nil, fmt.Errorf("opening %s failed with: %s", vars.RepoFileName, err)
	}
	defer file.Close()

	var repoNames []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		repoNames = append(repoNames, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return repoNames, nil
}
