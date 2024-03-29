package config

import (
	"github.com/spf13/viper"
)

var cfg *Config

type Config struct {
	WebServerPort            string `mapstructure:"WEBSERVER_PORT"`
	RedisAddress             string `mapstructure:"REDIS_ADDRESS"`
	RedisPassword            string `mapstructure:"REDIS_PASSWORD"`
	MaxRequestPerMinuteIP    int    `mapstructure:"MAX_REQUESTS_PER_MINUTE_IP"`
	MaxRequestPerMinuteToken int    `mapstructure:"MAX_REQUESTS_PER_MINUTE_TOKEN"`
	JWTSecret                string `mapstructure:"JWT_SECRET"`
	JWTExpiresInMinutes      int    `mapstructure:"JWT_EXPIRES_IN_MINUTES"`
}

func LoadConfig(path string) (*Config, error) {
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
