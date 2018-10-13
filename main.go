package main

import "github.com/zhs007/tradingwebserv/base"

func main() {
	base.LoadConfig("./cfg/config.yaml")

	cfg := base.GetConfig()

	StartServ(cfg.WebServAddr)
}
