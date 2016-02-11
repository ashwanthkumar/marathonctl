package cmd

import (
	"errors"

	"github.com/ashwanthkumar/marathonctl/repo"
	"github.com/spf13/cobra"
)

var repoAdd = &cobra.Command{
	Use:   "add <name> <location>",
	Short: "Add a package repository to local cache",
	Long:  "Add a package repository to local cache",
	Run:   AttachHandler(addRepoCache),
}

func addRepoCache(args []string) (err error) {
	if len(args) != 2 {
		return errors.New("We need exactly 2 arguments")
	}
	name := args[0]
	location := args[1]
	return repo.Add(name, location)
}

func init() {
	repoCommand.AddCommand(repoAdd)
}
