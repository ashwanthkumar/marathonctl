package cmd

import (
	"errors"
	"github.com/ashwanthkumar/marathonctl/packages"
	"github.com/spf13/cobra"
)

var PackageRm = &cobra.Command{
	Use:   "rm <repository>",
	Short: "Remove a package repository from local cache",
	Long:  "Remove a package repository from local cache",
	Run:   AttachHandler(rmPackageCache),
}

func rmPackageCache(args []string) (err error) {
	if len(args) > 0 {
		repository := args[0]
		return packages.Remove(repository)
	}

	return errors.New("We need a repository name to remove")
}

func init() {
	Package.AddCommand(PackageRm)
}
