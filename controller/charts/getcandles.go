package charts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCandles -
func GetCandles() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "getcandles.html", gin.H{})
	}
}
