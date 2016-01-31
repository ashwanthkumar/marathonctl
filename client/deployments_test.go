package client

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestDeployments(t *testing.T) {
  marathon := newFakeMarathonServer(t)
  deployments, err := marathon.Deployments()
  assert.NoError(t, err)

  deploy1 := ActiveDeployment {
    Id: "573aff0d-c7bc-48f5-b453-8b450beeb241",
    Version: "2016-01-31T18:21:29.964Z",
  }
  actualDeployments := Deployments {
    Deployments: []ActiveDeployment {deploy1},
  }

  assert.Equal(t, deployments, &actualDeployments)
}

func TestIsStillDeploying(t *testing.T) {
  marathon := newFakeMarathonServer(t)
  deployment, err := marathon.Deploy("my-app", "appspec", false)
  assert.NoError(t, err)

  stillHappening, err := marathon.IsStillDeploying(deployment)
  assert.NoError(t, err)

  assert.Equal(t, stillHappening, true)
}
