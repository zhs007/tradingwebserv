package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/zhs007/tradingwebserv/model/trading"
)

// GetCandles - [[time, open, close, low, high, volume], ...]
func GetCandles() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Request.FormValue("code")
		name := c.Request.FormValue("name")
		starttime := c.Request.FormValue("starttime")
		endtime := c.Request.FormValue("endtime")
		timezone := c.Request.FormValue("timezone")

		loc, err := time.LoadLocation(timezone)
		if err != nil {
			c.String(http.StatusOK, err.Error())

			return
		}

		ret, err := trading.GetCandles(c.Request.Context(), code, name, starttime, endtime, loc)
		if err != nil {
			c.String(http.StatusOK, err.Error())

			return
		}

		strret, err := trading.FormatCandles2Arr(loc, ret)
		if err != nil {
			c.String(http.StatusOK, err.Error())
		} else {
			c.String(http.StatusOK, strret)
		}
	}
}
