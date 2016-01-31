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
    fakeMarathonHandler.HandleFunc("/v2/info", func(writer http.ResponseWriter, reader *http.Request) {
      contents, err := ioutil.ReadFile("../test-fixtures/info.json")
      if err != nil {
        t.Fatal(err)
      }
      writer.Header().Add("Content-Type", "application/json")
      writer.Write(contents)
    })
    fakeServer = httptest.NewServer(fakeMarathonHandler)    
  })

  return Marathon{
    Url: fakeServer.URL,
  }
}
