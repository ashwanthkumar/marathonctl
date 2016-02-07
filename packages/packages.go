package packages

import (
  "io/ioutil"
  "fmt"
  "os"
  "errors"
  "github.com/ashwanthkumar/marathonctl/config"
  "github.com/joeshaw/multierror"
  fetcher "github.com/hashicorp/go-getter"
)

const INDEX_FILE = "repos.json"

var PACKAGE_CACHE_LOCATION string
var AllRepos *Repositories

func init() {
  PACKAGE_CACHE_LOCATION = config.GetPackageCachePath() + "/" + INDEX_FILE
  data, err := ioutil.ReadFile(PACKAGE_CACHE_LOCATION)
  if os.IsNotExist(err) {
    AllRepos = DefaultRepositories()
  } else { 
    handleErrorIfAny(err)
  }
  if AllRepos == nil {
    AllRepos, err = Deserialize(data)
    handleErrorIfAny(err)
  }
}

func Add(name, location string) error {
  if AllRepos.Exists(name) {
    return errors.New(name + " package repository already exist")
  }
  fmt.Printf("Adding package repository %s from %s\n", name, location)
  AllRepos.Add(Repository {
    Name: name,
    Loc: location,
  })
  err := Update(name)
  if err != nil {
    return err
  }

  return WritePackageMetadata()
}

func Remove(name string) (err error) {
  AllRepos.Remove(name)
  packageRepoPath := config.GetPackageCachePath() + "/" + name
  fmt.Printf("Removing %s from %s\n", name, packageRepoPath)
  err = os.RemoveAll(packageRepoPath)
  if err != nil {
    return err
  }
  return WritePackageMetadata()
}

func Update(name string) (err error) {
  repository := AllRepos.Get(name)
  if repository == nil {
    fmt.Errorf("%s package repository not found\n", name)
    return errors.New(name + " package repository not found")
  }
  fmt.Println("Updating " + name + " package repository from " + repository.Loc)
  packageRepoPath := config.GetPackageCachePath() + "/" + name
  return fetcher.Get(packageRepoPath, repository.Loc)
}

func UpdateAll() error {
  var Errors multierror.Errors
  for _, repository := range *AllRepos {
    err := Update(repository.Name)
    if err != nil {
      Errors = append(Errors, err)
    }
  }
  return Errors.Err()
}

func WritePackageMetadata() error {
  data, err := AllRepos.Serialize()
  if err != nil {
    return err
  }
  return ioutil.WriteFile(PACKAGE_CACHE_LOCATION, data, 0644)
}

func handleErrorIfAny(err error) {
  if err != nil {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }
}
