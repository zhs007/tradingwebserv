package router

import (
	"github.com/zhs007/tradingwebserv/controller/api"
	"github.com/zhs007/tradingwebserv/controller/charts"

	"github.com/gin-gonic/gin"
)

// Router -
var Router *gin.Engine

func init() {
	Router = gin.Default()
}

// SetRouter -
func SetRouter() {
	Router.LoadHTMLGlob("./www/views/*.html")

	Router.GET("/api/getcandles", api.GetCandles())
	Router.POST("/api/getcandles", api.GetCandles())
	Router.GET("/api/gettradingdata", api.GetTradingData())
	Router.POST("/api/gettradingdata", api.GetTradingData())

	Router.GET("/charts/getcandles", charts.GetCandles())
	Router.GET("/charts/getpnl", charts.GetPNL())
}
