package client

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestServerVersion(t *testing.T) {
  marathon := newFakeMarathonServer(t)
  version, err := marathon.ServerVersion()
  assert.NoError(t, err)

  actualVersion := ServerVersion {
    Version: "0.8.1",
    Name: "marathon",
  }
  assert.Equal(t, version, &actualVersion)
}

func TestDeploy(t *testing.T) {
  marathon := newFakeMarathonServer(t)
  deployment, err := marathon.Deploy("my-app", "appspec", false)
  assert.NoError(t, err)

  actualDeployment := Deployment {
    DeploymentID: "83b215a6-4e26-4e44-9333-5c385eda6438",
    Version: "2014-08-26T07:37:50.462Z",
  }

  assert.Equal(t, deployment, &actualDeployment)
}
