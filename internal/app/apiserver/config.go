package apiserver

import "go-rest/internal/app/store"

type Config struct {
	BindAddress string `toml:"bind_addr"`
	LogLevel    string `toml:"logging_level"`
	Store       *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddress: ":8080",
		LogLevel:    "debug",
		Store:       store.NewConfig(),
	}
}
