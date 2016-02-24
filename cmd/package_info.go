package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/ashwanthkumar/marathonctl/packages"
	"github.com/spf13/cobra"
)

var packageInfo = &cobra.Command{
	Use:   "info <name>",
	Short: "Display information about a package",
	Long:  "Display information about a package",
	Run:   AttachHandler(packageInfoHandler),
}

func packageInfoHandler(args []string) error {
	if len(args) != 1 {
		return errors.New("Need atleast 1 argument for find command")
	}
	searchName := args[0]
	latestPackageInfo, err := packages.Info(searchName)
	if err == nil {
		b, err := json.Marshal(latestPackageInfo)
		if err != nil {
			return err
		}
		var out bytes.Buffer
		json.Indent(&out, b, "", "  ")
		out.WriteTo(os.Stdout)
		fmt.Printf("\n")
	}
	return err
}

func init() {
	packageCommand.AddCommand(packageInfo)
}
