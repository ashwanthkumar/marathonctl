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
