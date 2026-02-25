package config

import "github.com/spf13/viper"

type Config struct {
	LogLevel string
	Server   struct {
		Host string
		Mode string
	}
	MySQL struct {
		Dsn string
	}
	Ump struct {
		Url             string
		AdminToken      string
		CookieName      string
		APIAppID        int
		MenuAppID       int
		ResourceAppID   int
		FunctionalAppID int
	}
}

var cfg *Config

func GetConf() *Config {
	return cfg
}

func init() {
	viper.SetConfigFile("./config_file/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}
}
