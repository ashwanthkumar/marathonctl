package cmd

import (
	"errors"
	"github.com/ashwanthkumar/marathonctl/packages"
	"github.com/spf13/cobra"
)

var PackageAdd = &cobra.Command{
	Use:   "add <name> <location>",
	Short: "Add a package repository to local cache",
	Long:  "Add a package repository to local cache",
	Run:   AttachHandler(addPackageCache),
}

func addPackageCache(args []string) (err error) {
	if len(args) != 2 {
		return errors.New("We need exactly 2 arguments")
	}
	name := args[0]
	location := args[1]
	return packages.Add(name, location)
}

func init() {
	Package.AddCommand(PackageAdd)
}
