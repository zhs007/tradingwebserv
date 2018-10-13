package router

import (
	"github.com/zhs007/tradingwebserv/controller/api"

	"github.com/gin-gonic/gin"
)

// Router -
var Router *gin.Engine

func init() {
	Router = gin.Default()
}

// SetRouter -
func SetRouter() {
	Router.GET("/api/getcandles", api.GetCandles())
	Router.POST("/api/getcandles", api.GetCandles())
}
