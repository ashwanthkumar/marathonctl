package client

import (
  "encoding/json"
)

type Deployment struct {
  DeploymentID    string `json:"deploymentId"`
  Version         string `json:"version"`
}
// Deploy an application spec (JSON string) to Marathon
// appSpec - App specification in a valid JSON string
// force - Should we do a force deployment?
// returns deploymentID in string, error if any
func (m *Marathon) Deploy(app string, appSpec string, force bool) (*Deployment, error) {
  request := request.Put(m.Url + "/v2/apps/" + app)
  if force {
    request = request.Query("force=true")
  }
  body, err := handle(request.End())
  if err != nil {
    return nil, err
  }

  var deployment Deployment
  err = json.Unmarshal([]byte(body), &deployment)
  return &deployment, err
}
