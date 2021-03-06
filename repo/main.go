package repo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ashwanthkumar/marathonctl/config"
	fetcher "github.com/hashicorp/go-getter"
	"github.com/joeshaw/multierror"
)

const repoIndixFile = "repos.json"

var packageCacheLocation string
var allRepositories *Repositories

func init() {
	packageCacheLocation = config.GetPackageCachePath() + "/" + repoIndixFile
	data, err := ioutil.ReadFile(packageCacheLocation)
	if os.IsNotExist(err) {
		allRepositories = DefaultRepositories()
	} else {
		handleErrorIfAny(err)
	}
	if allRepositories == nil {
		allRepositories, err = Deserialize(data)
		handleErrorIfAny(err)
	}
}

// Add the new remote repository
func Add(name, location string) error {
	if allRepositories.Exists(name) {
		return errors.New(name + " package repository already exist")
	}
	fmt.Printf("Adding package repository %s from %s\n", name, location)
	allRepositories.Add(Repository{
		Name: name,
		Loc:  location,
	})
	err := Update(name)
	if err != nil {
		return err
	}

	return writePackageMetadata()
}

// Remove a known remote repository
func Remove(name string) (err error) {
	repository := allRepositories.Get(name)
	if repository == nil {
		return fmt.Errorf("%s package repository not found\n", name)
	}
	allRepositories.Remove(repository.Name)
	fmt.Printf("Removing %s from %s\n", name, repository.LocationOnDisk())
	err = os.RemoveAll(repository.LocationOnDisk())
	if err != nil {
		return err
	}
	return writePackageMetadata()
}

// Update the given remote repository
func Update(name string) (err error) {
	repository := allRepositories.Get(name)
	if repository == nil {
		return fmt.Errorf("%s package repository not found\n", name)
	}
	fmt.Println("Updating " + repository.Name + " package repository from " + repository.Loc)
	return fetcher.Get(repository.LocationOnDisk(), repository.Loc)
}

// UpdateAll remote repositories
func UpdateAll() error {
	var Errors multierror.Errors
	for _, repository := range *allRepositories {
		err := Update(repository.Name)
		if err != nil {
			Errors = append(Errors, err)
		}
	}
	return Errors.Err()
}

// List all remote repositories
func List() *Repositories {
	return allRepositories
}

func writePackageMetadata() error {
	data, err := allRepositories.Serialize()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(packageCacheLocation, data, 0644)
}

func handleErrorIfAny(err error) {
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
