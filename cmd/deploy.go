package cmd

import (
  "fmt"
  "errors"

  "github.com/spf13/cobra"
  "github.com/ashwanthkumar/marathonctl/appconfig"
)

// Main command for Cobra.
var Deploy = &cobra.Command{
  Use:   "deploy <app.json>",
  Short: "Deploy an app using Marathon's app definition",
  Long:  "Deploy an app using Marathon's app definition",
  Run: AttachHandler(performDeploy),
}

var environment string
var dryRun bool
var force bool
func prepareFlags() {
  Deploy.PersistentFlags().StringVarP(
    &environment, "environment", "e", "test", "Environment to deploy")
  Deploy.PersistentFlags().BoolVarP(
    &dryRun, "dry-run", "d", false, "Print the final application configuration but don't deploy")
  Deploy.PersistentFlags().BoolVarP(
    &force, "force", "f", false, "Force deploy the app")
}

func init() {
  prepareFlags()
  MarathonCtl.AddCommand(Deploy)
}

func performDeploy(args []string) error{
  if(len(args) != 1) {
    return errors.New("deploy takes only 1 argument - Marathon's application definition")
  }
  path := args[0]
  fmt.Println("Deploying application from", path)
  appConfig, err := appconfig.Render(environment, path)
  if dryRun {
    fmt.Println(appConfig)
    return err
  }

  // TODO - PUT to Marathon API and wait on it
  return err
}
