package config

import (
  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"
)

func AddDefaults(c *viper.Viper) {
  c.SetDefault("url", "http://localhost:8080")
  cachePath, _ := homedir.Expand("~/.marathonctl/packages")
  c.SetDefault("package-cache-path", cachePath)
  c.SetDefault("package-repo", "github.com/ashwanthkumar/marathonctl-universe")
}
