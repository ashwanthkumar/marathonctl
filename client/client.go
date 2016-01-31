package client

import (
  "errors"
  "encoding/json"
  "time"
  "github.com/parnurzeal/gorequest"
)

type Marathon struct{
  Url string
}

var request = gorequest.New()

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

type ServerVersion struct {
  Version         string `json:"version"`
  Name            string `json:"name"`
}
// Get the server version of the Marathon instance
func (m *Marathon) ServerVersion() (*ServerVersion, error) {
  body, err := handle(request.Timeout(5000 * time.Millisecond).Get(m.Url + "/v2/info").End())
  if err != nil {
    return nil, err
  }

  var version ServerVersion
  err = json.Unmarshal([]byte(body), &version)
  return &version, err
}

func handle(response gorequest.Response, body string, errs []error) (string, error) {
  return string(body), combineErrors(errs)
}

func combineErrors(errs []error) error {
  if(len(errs) == 1) {
    return errs[0]
  } else if(len(errs) > 1) {
    msg := "Error(s):"
    for _, err := range errs {
      msg += " " + err.Error()
    }
    return errors.New(msg)
  } else {
    return nil
  }
}
