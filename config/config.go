package config

import (
  "log"
  "fmt"
  "github.com/spf13/viper"
)

var c = viper.New()

func init() {
  c.SetConfigName(".marathon")
  c.SetConfigType("json")
  c.AddConfigPath("$HOME/")
  err := c.ReadInConfig() // Find and read the config file

  if err != nil { // Handle errors reading the config file
      panic(fmt.Errorf("Fatal error config file: %s \n", err))
  }
}

func GetUrl() string {
  notNullConfiguration("url")
  return c.GetString("url")
}

// TODO - Add support for auth tokens

func notNullConfiguration(key string) {
  if c.GetString(key) == "" {
    log.Fatal(key + " configuration is not found in " + c.ConfigFileUsed())
  }
}
