package initialize

import (
	"common-lib/config"
	"flag"
	"jaden/we-life/common"
)

func SetupConfig() {
	common.Cfg = config.SetupConfig(configPath())
}

func configPath() string {
	defaultConfig := "conf/dev.json"
	confPath := flag.String("conf", defaultConfig, "config photo path")
	flag.Parse()
	return *confPath
}
