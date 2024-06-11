package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Configuration struct {
	Env      string
	Database Database
	HTTP     HTTP
}

type Database struct {
	Name            string
	User            string
	Password        string
	Host            string
	Port            string
	ConnectionRetry int
}

type HTTP struct {
	Domain string
	Port   string
}

func MustNewConfig() Configuration {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("unable to decode config file, %v", err))
	}

	config := Configuration{}
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Sprintf("unable to decode config file into configuration, %v", err))
	}

	return config
}
