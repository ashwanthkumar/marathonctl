package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/ashwanthkumar/marathonctl/appconfig"
	"github.com/ashwanthkumar/marathonctl/util"

	"github.com/gosuri/uiprogress"
	"github.com/spf13/cobra"
)

var Deploy = &cobra.Command{
	Use:   "deploy <app.json>",
	Short: "Deploy an app using Marathon's app definition",
	Long:  "Deploy an app using Marathon's app definition",
	Run:   AttachHandler(performDeploy),
}

var environment string
var dryRun bool
var force bool
var timeoutInSeconds int

func prepareFlags() {
	Deploy.PersistentFlags().StringVarP(
		&environment, "environment", "e", "test", "Environment to deploy")
	Deploy.PersistentFlags().BoolVarP(
		&dryRun, "dry-run", "d", false, "Print the final application configuration but don't deploy")
	Deploy.PersistentFlags().BoolVarP(
		&force, "force", "f", false, "Force deploy the app")
	Deploy.PersistentFlags().IntVarP(
		&timeoutInSeconds, "timeout", "t", 60*15, "timeout in seconds for deployment to complete, else we'll fail")
}

func init() {
	prepareFlags()
	MarathonCtl.AddCommand(Deploy)
}

func performDeploy(args []string) error {
	if len(args) != 1 {
		return errors.New("deploy takes only 1 argument - Marathon's application definition file")
	}
	path := args[0]
	renderedConfig, err := appconfig.Render(environment, path)
	if dryRun {
		fmt.Println("Rendered app specification")
		fmt.Println(renderedConfig)
		return err
	}

	var configMap map[string]interface{}
	util.JsonDecode(renderedConfig, &configMap)
	appName := configMap["id"].(string)

	fmt.Println("Deploying " + appName + " from " + path)
	deployment, err := marathon.Deploy(appName, renderedConfig, force)
	if err != nil {
		return err
	}

	// Do fancy UI updates on the screen
	uiprogress.Start()
	defer uiprogress.Stop()
	bar := uiprogress.AddBar(timeoutInSeconds + 1)
	bar.AppendCompleted()
	bar.PrependElapsed()
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return appName + " deploying"
	})

	for bar.Incr() {
		time.Sleep(time.Millisecond * 1000)
		if bar.Current()%15 == 0 { // check the deployment stauts only every 15 seconds
			stillDeploying, err := marathon.IsStillDeploying(deployment)
			if err != nil {
				return err
			}
			if !stillDeploying {
				bar.Set(timeoutInSeconds)
				time.Sleep(time.Millisecond * 20) // wait for a few ms to refresh the output
				fmt.Println("Deployment completed")
				return nil
			}
		}
	}

	return errors.New("[ERROR] App deployment timed out. Check Marathon UI for more details.")
}
