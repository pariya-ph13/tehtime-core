package app

import (
	"fmt"

	"github.com/TehranTime/tehtime-core/internal/config"
	"github.com/spf13/viper"
)

func initConfig(config *config.Config) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Error reading config file: %w", err)
	}

	if err := viper.Unmarshal(config); err != nil {
		return fmt.Errorf("Unable to decode into struct: %w", err)
	}

	return nil
}

// NewConfig constructs the application configuration instance and loads it from file/env.
func NewConfig() (*config.Config, error) {
	var cfg config.Config
	if err := initConfig(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
