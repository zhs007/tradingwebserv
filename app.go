package main

import (
	"github.com/zhs007/tradingwebserv/router"
)

// StartServ -
func StartServ(servaddr string) {
	r := router.Router
	router.SetRouter()
	r.Run(servaddr)
}
