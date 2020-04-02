package initialize

import (
	"common-lib/module"
	"fmt"
	"jaden/we-life/common"
)

func SetUpModule() {
	module.InitModule(&module.Values{
		Env:                 common.Cfg.GetString("env"),
		Service:             fmt.Sprintf("%s.%s", common.Cfg.GetString("name"), common.Cfg.GetString("nameSpace")),
		DingTalkAccessToken: common.Cfg.GetString("dingtalktoken"),
	})
}
