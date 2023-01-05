package apiserver

// Config ..

type Config struct {
	BindAdd     string `toml:"bind_address"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
}

func NewConfig() *Config {
	return &Config{}
}
