package client

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v53/github"
)

func writeRepositoriesToFile(fileName string, repos []*github.Repository) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, repo := range repos {
		name := strings.Split(*repo.FullName, "/")[1]
		_, err := file.WriteString(name + "\n")
		if err != nil {
			return fmt.Errorf("writing repo %s into file failed with: %s", name, err)
		}
	}

	return nil
}
