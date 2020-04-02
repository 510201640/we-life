package common

import (
	timeutil "common-lib/time-util"
	"fmt"
	"github.com/Sirupsen/logrus"
	"time"
)

type MyLogger struct {
	InnerLogger *logrus.Logger
}

var Logger *MyLogger

func (log MyLogger) Info(param, result, other string) {
	Logger.InnerLogger.WithFields(
		logrus.Fields{
			"content": result,
			"param":   param,
			"other":   other,
			"time":    timeutil.Parse("2006-01-02 15:04:05", time.Now().String()),
		},
	).Info()
}

func (log MyLogger) Error(param, result, other string) {
	Logger.InnerLogger.WithFields(
		logrus.Fields{
			"content": result,
			"param":   param,
			"other":   other,
			"time":    timeutil.Parse("2006-01-02 15:04:05", time.Now().String()),
		},
	).Error()
}

func (log MyLogger) SendDingTalkMsg(title string, describe string) {
	fmt.Println("发送钉钉消息")
}
