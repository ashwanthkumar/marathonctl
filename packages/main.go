package packages

import (
	"io/ioutil"
	"regexp"

	"github.com/ashwanthkumar/marathonctl/repo"
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
