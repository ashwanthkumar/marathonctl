package cmd

import (
	"fmt"

	"github.com/ashwanthkumar/marathonctl/repo"
	"github.com/spf13/cobra"
)

var repoLs = &cobra.Command{
	Use:   "ls",
	Short: "List all the package repositories",
	Long:  "List all the package repositories",
	Run:   AttachHandler(listPackageRepositories),
}

func listPackageRepositories(args []string) (err error) {
	for _, repository := range *repo.List() {
		fmt.Printf("%s\t%s\n", repository.Name, repository.Loc)
	}
	return err
}

func init() {
	repoCommand.AddCommand(repoLs)
}
