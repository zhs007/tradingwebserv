package trading

import (
	"github.com/zhs007/ankadb/client"
	"github.com/zhs007/tradingwebserv/base"
)

var singleClient ankadbclient.AnkaClient

func checkInitMod() {
	if singleClient != nil {
		return
	}

	cfg := base.GetConfig()

	singleClient = ankadbclient.NewClient()
	singleClient.Start(cfg.TradingDB)
}
