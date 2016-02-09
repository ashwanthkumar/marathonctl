package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeploy(t *testing.T) {
	marathon := newFakeMarathonServer(t)
	deployment, err := marathon.Deploy("my-app", "appspec", false)
	assert.NoError(t, err)

	actualDeployment := Deployment{
		DeploymentID: "573aff0d-c7bc-48f5-b453-8b450beeb241",
		Version:      "2016-01-31T18:21:29.964Z",
	}

	assert.Equal(t, deployment, &actualDeployment)
}
