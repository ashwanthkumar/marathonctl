package config

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var c = viper.New()

func init() {
	c.SetConfigName("config")
	c.SetConfigType("json")
	c.AddConfigPath("$HOME/.marathonctl")
	AddDefaults(c)

	err := c.ReadInConfig()
	// ignore the file not found error and catch everything else
	if err != nil && !os.IsNotExist(err) {
		fmt.Printf("Err %v\n", err)
		os.Exit(1)
	}
}

func BindFlags(flags *pflag.FlagSet) {
	c.BindPFlags(flags)
}

func GetString(key string) string {
	return c.GetString(key)
}

func GetPackageCachePath() string {
	return GetString("package-cache-path")
}

func AllSettings() map[string]interface{} {
	return c.AllSettings()
}
