package main

import (
	"fmt"
	"os"

	"github.com/ashwanthkumar/marathonctl/cmd"
)

var APP_VERSION = "dev-build"

func main() {
	setupSignalHandlers()

	// since we overrie the versions in Makefile
	cmd.VERSION = APP_VERSION
	if err := cmd.MarathonCtl.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
