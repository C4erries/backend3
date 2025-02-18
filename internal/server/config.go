package server

type Config struct {
	binaddr     string `toml:"binaddr"`
	frontendUrl string `toml:"frontendUrl"`
}

func ConfigInit() *Config {
	return &Config{}
}
