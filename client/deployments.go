package client

import (
  "encoding/json"
)

// Get active deployments in the cluster
func (m *Marathon) Deployments() (*Deployments, error) {
  body, err := handle(httpClient.Get(m.Url + "/v2/deployments").End())
  if err != nil {
    return nil, err
  }

  var deployments Deployments
  err = json.Unmarshal([]byte(body), &deployments.Deployments)
  return &deployments, err
}

// Check if the Deployment is still happening. Usually we call this after
// calling the Deploy() method to wait.
// deployment   Deployment returned by Deploy() method
func (m *Marathon) IsStillDeploying(deployment *Deployment) (bool, error) {
  deployments, err := m.Deployments()
  if err != nil {
    return false, err
  }

  isPresent := deployments.Contains(deployment)
  return isPresent, nil
}

type Deployments struct {
  Deployments []ActiveDeployment
}
func (d *Deployments) Contains(deployment *Deployment) bool {
    for _, activeDeployment := range d.Deployments {
      if activeDeployment.Id == deployment.DeploymentID && activeDeployment.Version == deployment.Version {
        return true
      }
    }
    return false
}

type ActiveDeployment struct {
  Version       string `json:"version"`
  Id            string `json:"id"`
}
