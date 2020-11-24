package webserver

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"sentaly.com/telegram-bot/webserver/routes/healthcheck"
)

// Run webserver
func Run(host string, port int) {
	r := gin.Default()

	api := r.Group("/api")

	healthcheck.Routes(api)

	addr := fmt.Sprintf("%s:%d", host, port)
	r.Run(addr)
}