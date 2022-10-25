package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server string `mapstructure:"server"`
	SearchURL string `mapstructure:"search_url"`
	TopItaliansURL string `mapstructure:"top_italians"`
	SongUrl string `mapstructure:"song_url"`
}

var Conf *Config

func LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		return err
	}
	return nil
}
