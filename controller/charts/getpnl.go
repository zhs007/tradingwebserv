package charts

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zhs007/tradingwebserv/model/trading"
)

// GetPNL -
func GetPNL() gin.HandlerFunc {
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

		lsttime, lstval := trading.FormatTradingData2PNLChart(loc, rtd)
		if err != nil {
			c.String(http.StatusOK, err.Error())
		} else {
			c.HTML(http.StatusOK, "getpnl.html", gin.H{"lsttime": lsttime, "lstval": lstval})
		}

	}
}
