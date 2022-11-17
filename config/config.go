package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server string `mapstructure:"SERVER"`
	SearchUrl string `mapstructure:"SEARCH_URL"`
	PlaylistUrl string `mapstructure:"PLAYLIST_URL"`
	SongUrl string `mapstructure:"SONG_URL"`
}

var Conf *Config

func LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
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
