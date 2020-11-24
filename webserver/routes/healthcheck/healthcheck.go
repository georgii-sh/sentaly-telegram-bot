package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func respondOk(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

// Routes for monitoring
func Routes(r *gin.RouterGroup) {
	r.GET("/healthcheck", respondOk)
}