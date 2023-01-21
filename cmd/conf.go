package main

import (
	"path/filepath"

	"github.com/TiregeRRR/image_service/internal/pkg/config"
	"github.com/spf13/viper"
)

func LoadConfig(path string) (*config.Config, error) {
	viper.SetConfigFile(filepath.Join(path, "dev.env"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var conf config.Config
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
