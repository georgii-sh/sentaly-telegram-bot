package main

import (
	"log"
	"os"

	"sentaly.com/telegram-bot/configloader"
	"sentaly.com/telegram-bot/webserver"
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
	logger := log.New(os.Stderr, "", 0)

	var config AppConfig
	configloader.Load("config.yml", &config)
	
	logger.Println("loaded config: ", config)

	webserver.Run(config.Server.Host, config.Server.Port)
}
