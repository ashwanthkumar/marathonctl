package repo

import (
	"encoding/json"
)

// Repository represents the remote repo structure
type Repository struct {
	// Name of the source - eg. universe
	Name string `json:"name"`
	// Location of the source - eg. github.com/ashwanthkumar/marathonctl-universe
	// We support all the protocols supported by https://github.com/hashicorp/go-getter
	Loc string `json:"loc"`
}

// Repositories are the source of truth of all the package repositories in the system.
// General location of the file is ~/.marathonctl/repos.source
type Repositories []Repository

// Exists - Do we already know this repo?
func (r *Repositories) Exists(repo string) bool {
	for _, repository := range *r {
		if repository.Name == repo {
			return true
		}
	}
	return false
}

// Get the Repository represented by the name
func (r *Repositories) Get(repo string) *Repository {
	for _, repository := range *r {
		if repository.Name == repo {
			return &repository
		}
	}
	return nil
}

// Add new repository to the list
func (r *Repositories) Add(repo Repository) *Repositories {
	*r = append(*r, repo)
	return r
}

// Remove an existing repository, if it exists
func (r *Repositories) Remove(repo string) *Repositories {
	var newRepositories Repositories
	for _, repository := range *r {
		if repository.Name != repo {
			newRepositories = append(newRepositories, repository)
		}
	}

	*r = newRepositories
	return r
}

// Serialize the Repositories into a JSON []byte
func (r *Repositories) Serialize() ([]byte, error) {
	return json.Marshal(r)
}

// Deserialize from []byte to Repositories
func Deserialize(data []byte) (*Repositories, error) {
	var repos Repositories
	err := json.Unmarshal(data, &repos)
	return &repos, err
}

// DefaultRepositories - used the first time marathonctl is invoked
func DefaultRepositories() *Repositories {
	defaultRepo := Repository{
		Name: "universe",
		Loc:  "github.com/ashwanthkumar/marathonctl-universe",
	}
	var DefaultRepo Repositories
	DefaultRepo = []Repository{defaultRepo}
	return &DefaultRepo
}
