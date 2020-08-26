package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

//Config model to database start
type Config struct {
	Server   string
	Database string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
