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
    marathon = client.Marathon {
      Url: config.GetString("marathon.host"),
    }
    err := handler(args)
    if err != nil {
      log.Printf("[Error] %s", err.Error())
      os.Exit(1)
    }
  }
}

func init() {
  prepareFlagsForMarathonCtl()
  log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func prepareFlagsForMarathonCtl() {
  var str string // ignored since we directly pass the flag values as config overrides
  MarathonCtl.PersistentFlags().StringVarP(
    &str, "marathon.host", "", config.GetString("marathon.host"), "Marathon host in http://host:port form.")
  MarathonCtl.PersistentFlags().StringVarP(
    &str, "mesos.master", "", config.GetString("mesos.master"), "Mesos host in host:port form.")
  MarathonCtl.PersistentFlags().StringVarP(
    &str, "zk.host", "", config.GetString("zk.host"), "ZK host in host:port form.")
  config.BindFlags(MarathonCtl.PersistentFlags())
}
