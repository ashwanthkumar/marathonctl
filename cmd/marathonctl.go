package cmd

import (
  "log"
  "os"

  "github.com/ashwanthkumar/marathonctl/config"
  "github.com/ashwanthkumar/marathonctl/client"
  "github.com/spf13/cobra"
)

var marathon client.Marathon

// Main command for Cobra.
var MarathonCtl = &cobra.Command{
  Use:   "marathonctl <args>",
  Short: "Command line client to Marathon",
  Long:  `Command line client to Marathon`,
}

type CommandHandler func(args []string) error

func AttachHandler(handler CommandHandler) func (*cobra.Command, []string) {
  return func (cmd *cobra.Command, args []string) {
    err := handler(args)
    if err != nil {
      log.Printf("[Error] %s", err.Error())
      os.Exit(1)
    }
  }
}

func init() {
  marathon = client.Marathon {
    Url: config.GetUrl(),
  }
  log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
