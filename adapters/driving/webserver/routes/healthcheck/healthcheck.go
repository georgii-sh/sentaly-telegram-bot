package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func respondOk(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

// SetRoutes for monitoring
func SetRoutes(r *gin.RouterGroup) {
	r.GET("/healthcheck", respondOk)
}