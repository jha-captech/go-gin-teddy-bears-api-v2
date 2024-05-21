package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func HydrateConfigFromEnv() (Configuration, error) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		return Configuration{}, fmt.Errorf("unable to decode config file, %v", err)
	}

	config := Configuration{}
	if err := viper.Unmarshal(&config); err != nil {
		return Configuration{}, fmt.Errorf("unable to decode config file into configuration, %v", err)
	}

	return config, nil
}
