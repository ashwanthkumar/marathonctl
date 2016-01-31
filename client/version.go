package client

import (
  "time"
  "encoding/json"
)

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
