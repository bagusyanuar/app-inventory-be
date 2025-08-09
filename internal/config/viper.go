package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	cfg := viper.New()
	cfg.SetConfigFile(".env")
	err := cfg.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error environment file: %w", err))
	}
	return cfg
}
