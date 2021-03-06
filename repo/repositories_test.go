package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoriesExist(t *testing.T) {
	repos := newRepositories([]Repository{
		repo("r1", "loc1"), repo("r2", "loc2"),
	})
	assert.Equal(t, repos.Exists("r1"), true)
	assert.Equal(t, repos.Exists("r3"), false)
}

func TestRepositoriesAdd(t *testing.T) {
	repos := newRepositories([]Repository{
		repo("r1", "loc1"), repo("r2", "loc2"),
	})
	assert.Equal(t, repos.Exists("r3"), false)
	repos.
		Add(repo("r3", "loc3")).
		Add(repo("r4", "loc4"))
	assert.Equal(t, repos.Exists("r3"), true)
	assert.Equal(t, repos.Exists("r4"), true)
}

func TestRepositoriesRemove(t *testing.T) {
	repos := newRepositories([]Repository{
		repo("r1", "loc1"), repo("r2", "loc2"),
	})
	assert.Equal(t, true, repos.Exists("r2"))
	repos.Remove("r2")
	assert.Equal(t, false, repos.Exists("r2"))
}

func TestRepositoriesGet(t *testing.T) {
	repos := newRepositories([]Repository{
		repo("r1", "loc1"), repo("r2", "loc2"),
	})
	actual := repos.Get("r1")
	expected := repo("r1", "loc1")
	assert.Equal(t, &expected, actual)
}

func TestRepositoriesSerialize(t *testing.T) {
	repos := newRepositories([]Repository{
		repo("universe", "github.com/ashwanthkumar/marathonctl-universe"),
	})
	expectedJSON := `[{"name":"universe","loc":"github.com/ashwanthkumar/marathonctl-universe"}]`
	result, err := repos.Serialize()
	assert.NoError(t, err)
	assert.Equal(t, []byte(expectedJSON), result)
}

func TestDeserialize(t *testing.T) {
	reposInJSON := `[{
    "name": "universe",
    "loc": "github.com/ashwanthkumar/marathonctl-universe"
  }]`
	repos, err := Deserialize([]byte(reposInJSON))
	assert.NoError(t, err)
	assert.Equal(t, repos.Exists("universe"), true)
}

func TestDefaultRepositories(t *testing.T) {
	repos := DefaultRepositories()
	assert.Equal(t, true, repos.Exists("universe"))
}

func newRepositories(repos []Repository) Repositories {
	return repos
}

func repo(name, loc string) Repository {
	return Repository{
		Name: name,
		Loc:  loc,
	}
}
