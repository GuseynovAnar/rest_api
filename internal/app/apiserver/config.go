package apiserver

import "github.com/GuseynovAnar/rest_api.git/internal/app/store"

// Config
type Config struct {
	BindAdd  string `toml: "bind_address"`
	LogLevel string `toml: "log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAdd:  ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
