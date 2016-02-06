package cmd

import (
  "fmt"
  "github.com/ashwanthkumar/marathonctl/config"
  "github.com/spf13/cobra"
  fetcher "github.com/hashicorp/go-getter"
)

var PackageUpdate = &cobra.Command{
  Use:   "update",
  Short: "Update the local package cache",
  Long:  "Update the local package cache",
  Run: AttachHandler(updatePackageCache),
}

func updatePackageCache(args []string) (err error) {
  fmt.Printf("marathonctl is updating it's package cache in %s from %s\n", config.GetPackageCachePath(), config.GetPackageRepo())
  err = fetcher.Get(config.GetPackageCachePath(), config.GetPackageRepo())
  return err
}

func init() {
  Package.AddCommand(PackageUpdate)
}
