package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/c4erries/backend3/internal/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := server.ConfigInit()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Start(config); err != nil {
		log.Fatal(err)
	}

}
