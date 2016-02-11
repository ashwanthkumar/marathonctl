package cmd

import (
	"github.com/ashwanthkumar/marathonctl/repo"
	"github.com/spf13/cobra"
)

var repoUpdate = &cobra.Command{
	Use:   "update [repository]",
	Short: "Update the local package cache",
	Long:  "Update the local package cache",
	Run:   AttachHandler(updateRepoCaches),
}

func updateRepoCaches(args []string) (err error) {
	if len(args) > 0 {
		repository := args[0]
		return repo.Update(repository)
	}

	return repo.UpdateAll()
}

func init() {
	repoCommand.AddCommand(repoUpdate)
}
