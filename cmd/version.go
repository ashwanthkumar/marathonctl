package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var VERSION = "dev-build"

// Main command for Cobra.
var Version = &cobra.Command{
	Use:   "version",
	Short: "Version of the Marathon CLI",
	Long:  "Version of the Marathon CLI",
	Run:   AttachHandler(showVersion),
}

func showVersion(args []string) error {
	fmt.Println("Client: marathonctl/" + VERSION)
	version, err := marathon.ServerVersion()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Server: " + version.Name + "/" + version.Version)
	return err
}

func init() {
	MarathonCtl.AddCommand(Version)
}
