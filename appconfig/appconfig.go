package appconfig

import (
  "strings"
  "bytes"
  "os"
  "text/template"

  "github.com/ashwanthkumar/marathonctl/config"
)

type Context map[string]interface{}

func (c *Context) Env() map[string]string {
  env := make(map[string]string)
  for _, i := range os.Environ() {
    sep := strings.Index(i, "=")
    env[i[0:sep]] = i[sep+1:]
  }
  return env
}

func Render(environment, path string) (string, error) {
  var configInBytes bytes.Buffer
  var context Context
  context = config.AllSettings()
  context["DEPLOY_ENV"] = environment
  err := template.Must(template.ParseFiles(path)).Execute(&configInBytes, &context)
  return configInBytes.String(), err
}
