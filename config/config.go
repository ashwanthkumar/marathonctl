package config

import (
  "log"
  "fmt"
  "os"

  "github.com/spf13/viper"
)

var c = viper.New()
var initialized bool = false

func init() {
  c.SetConfigName(".marathon")
  c.SetConfigType("json")
  c.AddConfigPath("$HOME/")
}

func GetUrl() string {
  notNullConfiguration("url")
  return c.GetString("url")
}

// TODO - Add support for auth tokens

// For lazy init of the configuration
func Initialize() {
  if !initialized {
    err := c.ReadInConfig()
    if os.IsNotExist(err) {
      fmt.Println("Config file cannotbe found at your ~/.marathon.json")
      fmt.Println(err)
      os.Exit(1)
    }

    if err != nil { // Handle errors reading the config file
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }

    initialized = true
  }
}

func notNullConfiguration(key string) {
  if c.GetString(key) == "" {
    log.Fatal(key + " configuration is not found in " + c.ConfigFileUsed())
  }
}
