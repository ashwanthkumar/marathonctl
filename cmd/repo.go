package cmd

import (
	"github.com/spf13/cobra"
)

var repoCommand = &cobra.Command{
	Use:   "repo",
	Short: "Manage remote repositories where packages can be installed",
	Long:  "Manage remote repositories where packages can be installed",
}

func init() {
	MarathonCtl.AddCommand(repoCommand)
}
