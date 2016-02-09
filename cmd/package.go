package cmd

import (
	"github.com/spf13/cobra"
)

var Package = &cobra.Command{
	Use:   "package",
	Short: "Manage packages which needs to be installed on Marathon",
	Long:  "Manage packages which needs to be installed on Marathon",
}

func init() {
	MarathonCtl.AddCommand(Package)
}
