package initialize

import (
	"github.com/Sirupsen/logrus"
	"jaden/we-life/common"
	"os"
)

func NewLogger() {
	common.Logger = &common.MyLogger{
		InnerLogger: logrus.New(),
	}
	//用日志实例的方式使用日志
	common.Logger.InnerLogger.Out = os.Stdout //日志标准输出
	file, err := os.OpenFile("logs/jaden.log", os.O_CREATE|os.O_WRONLY, 1)
	if err == nil {
		common.Logger.InnerLogger.Out = file
	} else {
		panic("failed to log to photo.......")
		return
	}
}
