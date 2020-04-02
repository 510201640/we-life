package initialize

import (
	"jaden/we-life/common"
	"jaden/we-life/router"
)

func SetupServer() {
	router.New()
	router.SetApi()
	router.Run(common.Cfg.GetString("server.port"))
}
