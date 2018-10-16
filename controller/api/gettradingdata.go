package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/zhs007/tradingwebserv/model/trading"
)

// GetTradingData - [[id, type, side, newtime, price, volume, tradetime, avgprice, donevolume], ...]
func GetTradingData() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Request.FormValue("name")
		timezone := c.Request.FormValue("timezone")

		loc, err := time.LoadLocation(timezone)
		if err != nil {
			c.String(http.StatusOK, err.Error())

			return
		}

		rtd, err := trading.GetTradingData(c.Request.Context(), name)
		if err != nil {
			c.String(http.StatusOK, err.Error())

			return
		}

		strret, err := trading.FormatTradingData2Arr(loc, rtd)
		if err != nil {
			c.String(http.StatusOK, err.Error())
		} else {
			c.String(http.StatusOK, strret)
		}
	}
}
