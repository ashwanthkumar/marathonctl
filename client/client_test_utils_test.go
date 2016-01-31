package client

import (
  "testing"
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "sync"
)

var once sync.Once
var fakeServer *httptest.Server

func newFakeMarathonServer(t *testing.T) Marathon {
  once.Do(func() {
    fakeMarathonHandler := http.NewServeMux()
    // TODO - Add more handlers here as we implement more functionalities of the client
    fakeMarathonHandler.HandleFunc("/v2/info", serveFileAsJson(t, "../test-fixtures/info.json"))
    fakeMarathonHandler.HandleFunc("/v2/apps/my-app", serveFileAsJson(t, "../test-fixtures/deploy-app.json"))
    fakeMarathonHandler.HandleFunc("/v2/deployments", serveFileAsJson(t, "../test-fixtures/deployments.json"))
    fakeServer = httptest.NewServer(fakeMarathonHandler)    
  })

  return Marathon{
    Url: fakeServer.URL,
  }
}

func serveFileAsJson(t *testing.T,  filepath string) func(http.ResponseWriter, *http.Request) {
  return func(writer http.ResponseWriter, reader *http.Request) {
      contents, err := ioutil.ReadFile(filepath)
      if err != nil {
        t.Fatal(err)
      }
      writer.Header().Add("Content-Type", "application/json")
      writer.Write(contents)
    }
}
