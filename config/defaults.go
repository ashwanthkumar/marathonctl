package config

import "github.com/spf13/viper"

func AddDefaults(c *viper.Viper) {
  c.SetDefault("url", "http://localhost:8080")
}