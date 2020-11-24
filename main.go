package main

import (
	"fmt"

	"sentaly-telegram-bot/configloader"
)

// AppConfig data
type AppConfig struct {
	Server struct {
		Host string `yaml:"host", envconfig:"HOST"`
		Port int `yaml:"port", envconfig:"PORT"`
	} `yaml:"server"`
	Redis struct {
		Addr string `yaml:"address", envconfig:"REDIS_HOST"`
		Pass string `yaml:"password", envconfig:"REDIS_PASSWORD"`
	} `yaml:"redis"`
}

func main() {
	var config AppConfig
	configloader.Load("config.yml", &config)
	fmt.Printf("loaded config: %+v\n", config)
}
