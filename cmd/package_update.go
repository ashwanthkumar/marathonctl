package cmd

import (
	"github.com/ashwanthkumar/marathonctl/packages"
	"github.com/spf13/cobra"
)

var packageUpdate = &cobra.Command{
	Use:   "update [repository]",
	Short: "Update the local package cache",
	Long:  "Update the local package cache",
	Run:   AttachHandler(updatePackageCache),
}

func updatePackageCache(args []string) (err error) {
	if len(args) > 0 {
		repository := args[0]
		return packages.Update(repository)
	}

	return packages.UpdateAll()
}

func init() {
	packageCommand.AddCommand(packageUpdate)
}
