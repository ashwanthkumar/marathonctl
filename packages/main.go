package packages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"

	"github.com/ashwanthkumar/marathonctl/repo"
	"github.com/hashicorp/go-version"
)

// Find for a package by "name" through all the packages
// across all the repositories
func Find(name string) ([]string, error) {
	var allMatchingPackages []string
	for _, repository := range *repo.List() {
		packageDir := repository.LocationOnDisk() + "/" + "packages"
		packages, err := ioutil.ReadDir(packageDir)
		if err != nil {
			return nil, err
		}

		for _, file := range packages {
			if file.IsDir() {
				match, _ := regexp.MatchString(name, file.Name())
				if match {
					allMatchingPackages = append(allMatchingPackages, file.Name())
				}
			}
		}
	}

	return allMatchingPackages, nil
}

// Info - Returns a package information
func Info(name string) (*Package, error) {
	var packageInfo *Package
	for _, repository := range *repo.List() {
		packageDir := fmt.Sprintf("%s/packages/%s", repository.LocationOnDisk(), name)
		versions, err := ioutil.ReadDir(packageDir)
		if err != nil {
			return nil, err
		}

		numberOfVersions := len(versions)
		packageVersions := make([]*version.Version, numberOfVersions)
		for i, packageVersion := range versions {
			v, _ := version.NewVersion(packageVersion.Name())
			packageVersions[i] = v
		}
		sort.Sort(version.Collection(packageVersions))

		if len(packageVersions) > 0 {
			latestVersion := packageVersions[numberOfVersions-1]
			latestPackageMetadata := fmt.Sprintf("%s/%s/package.json", packageDir, latestVersion)
			file, err := os.Open(latestPackageMetadata)
			if err != nil {
				return nil, err
			}
			data, err := ioutil.ReadAll(file)
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal(data, &packageInfo)
			return packageInfo, err
		}
	}

	return packageInfo, nil
}
