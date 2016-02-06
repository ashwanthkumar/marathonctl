package cmd

import (
  "github.com/ashwanthkumar/marathonctl/packages"
  "github.com/spf13/cobra"
)

var PackageUpdate = &cobra.Command{
  Use:   "update",
  Short: "Update the local package cache",
  Long:  "Update the local package cache",
  Run: AttachHandler(updatePackageCache),
}

func updatePackageCache(args []string) (err error) {
  return packages.UpdateAll()
}

func init() {
  Package.AddCommand(PackageUpdate)
}
