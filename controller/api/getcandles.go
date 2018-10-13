package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCandles -
func GetCandles() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "get candles")
	}
}
