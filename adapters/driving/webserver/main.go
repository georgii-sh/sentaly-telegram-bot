package webserver

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"sentaly.com/telegram-bot/adapters/driving/webserver/routes/healthcheck"
	"sentaly.com/telegram-bot/adapters/driving/webserver/routes/telegram"
	"sentaly.com/telegram-bot/ports"
)

// Run webserver
func Run(s ports.BotService, host string, port int) {
	r := gin.Default()

	api := r.Group("/api")

	healthcheck.SetRoutes(api)
	telegram.SetRoutes(api, s)

	addr := fmt.Sprintf("%s:%d", host, port)
	r.Run(addr)
}