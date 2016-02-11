package packages

// Package represents the package.json structure inside all repositories
type Package struct {
	Name             string    `json:"name"`
	Version          string    `json:"version"`
	Description      string    `json:"description"`
	Developers       []string  `json:"developers"`
	PreInstallNotes  string    `json:"preInstallNotes"`
	PostInstallNotes string    `json:"postInstallNotes"`
	Licenses         []License `json:"licenses"`
}

// License represents the License used in the package
type License struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
