package cmd

import (
	"errors"
	"fmt"

	"github.com/ashwanthkumar/marathonctl/packages"
	"github.com/spf13/cobra"
)

var packageFind = &cobra.Command{
	Use:   "find <name>",
	Short: "Find for a package across all the package repositories",
	Long:  "Find for a package across all the package repositories",
	Run:   AttachHandler(findPackages),
}

func findPackages(args []string) error {
	if len(args) != 1 {
		return errors.New("Need atleast 1 argument for find command")
	}
	searchName := args[0]
	packages, err := packages.Find(searchName)
	if len(packages) != 0 {
		fmt.Printf("%d packages found\n", len(packages))
		fmt.Println("=================")
		for i, packageName := range packages {
			fmt.Printf("%d. %s\n", i+1, packageName)
		}
	} else {
		fmt.Println("No packages found for " + searchName)
	}
	return err
}

func init() {
	packageCommand.AddCommand(packageFind)
}
