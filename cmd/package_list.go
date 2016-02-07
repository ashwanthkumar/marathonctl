package cmd

import (
  "fmt"
  "github.com/ashwanthkumar/marathonctl/packages"
  "github.com/spf13/cobra"
)

var PackageList = &cobra.Command{
  Use:   "list",
  Short: "List all the package repositories",
  Long:  "List all the package repositories",
  Run: AttachHandler(listPackageRepositories),
}

func listPackageRepositories(args []string) (err error) {
  for _, repository := range *packages.AllRepos {
    fmt.Printf("%s\t%s\n", repository.Name, repository.Loc)
  }
  return err
}

func init() {
  Package.AddCommand(PackageList)
}
