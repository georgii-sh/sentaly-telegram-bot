package main

import (
	"log"
	"os"

	"sentaly.com/telegram-bot/adapters/driven/messenger"
	"sentaly.com/telegram-bot/adapters/driving/webserver"
	"sentaly.com/telegram-bot/application"
	"sentaly.com/telegram-bot/configloader"
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
	Telegram struct {
		Token string `yaml:"token", envconfig:"TELEGRAM_TOKEN"`
	} `yaml:"telegram"`
}

func main() {
	logger := log.New(os.Stderr, "", 0)

	var config AppConfig
	configloader.Load("config.yml", &config)
	
	logger.Println("loaded config: ", config)

	messenger := messenger.NewTelegramMessenger(config.Telegram.Token)
	service := application.NewDelegatingBotService(messenger)

	webserver.Run(service, config.Server.Host, config.Server.Port)
}
