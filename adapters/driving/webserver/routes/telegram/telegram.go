package telegram

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sentaly.com/telegram-bot/ports"
)

// SetRoutes for monitoring
func SetRoutes(r *gin.RouterGroup, s ports.BotService) {
	r.POST("/telegram", func (c *gin.Context) {
		update, err := s.ParseRequest(c.Request.Body)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		if err := s.ProcessRequest(update); err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
	
		c.Status(http.StatusOK)
	})
}