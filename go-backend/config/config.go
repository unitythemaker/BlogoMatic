package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func ParseConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func GetRssUrls() []string {
	return viper.GetStringSlice("config.feeds")
}
