package main

import (
	"log"
	"os"

	"deleter/commands/delete"
	"deleter/commands/list"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:  "github",
		Usage: "use to list and delete repo's",
		Commands: []*cli.Command{
			list.List(),
			delete.Delete(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		return
	}
}
