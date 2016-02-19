package config

import (
  "fmt"
  "os"

  "github.com/spf13/viper"
  "github.com/spf13/pflag"
)

var c = viper.New()

func init() {
  c.SetConfigName(".marathon")
  c.SetConfigType("json")
  c.AddConfigPath("$HOME/")
  AddDefaults(c)

  err := c.ReadInConfig()
  // ignore the file not found error and catch everything else
  if err != nil && !os.IsNotExist(err) {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }
}

func BindUrl(flag *pflag.Flag) {
  c.BindPFlag("url", flag)
}

func GetUrl() string {
  return c.GetString("url")
}
