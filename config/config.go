package config

import (
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	WebServerPort       string `mapstructure:"WEBSERVER_PORT"`
	JWTSecret           string `mapstructure:"JWT_SECRET"`
	JWTExpiresInMinutes int    `mapstructure:"JWT_EXPIRES_IN_MINUTES"`
}

func LoadConfig(path string) (*conf, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, err
}
