package config

import (
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func AddDefaults(c *viper.Viper) {
	c.SetDefault("marathon", kv("host", "http://localhost:8080"))
	c.SetDefault("mesos", kv("master", "localhost:5050"))
	c.SetDefault("zk", kv("host", "localhost:2181"))

	cachePath, _ := homedir.Expand("~/.marathonctl/packages")
	c.SetDefault("package-cache-path", cachePath)
	c.SetDefault("package-repo", "github.com/ashwanthkumar/marathonctl-universe")
}

func kv(key, value string) map[string]interface{} {
	tempMap := make(map[string]interface{})
	tempMap[key] = value
	return tempMap
}
