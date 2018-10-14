package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zhs007/tradingwebserv/model/trading"
)

// GetCandles -
func GetCandles() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Request.FormValue("code")
		name := c.Request.FormValue("name")
		starttime := c.Request.FormValue("starttime")
		endtime := c.Request.FormValue("endtime")
		timezone := c.Request.FormValue("timezone")

		ret, err := trading.GetCandles(c.Request.Context(), code, name, starttime, endtime, timezone)
		if err != nil {
			c.String(http.StatusOK, err.Error())
		} else {
			c.String(http.StatusOK, ret)
		}
	}
}
