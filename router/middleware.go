package router

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/go-errors/errors"
	"io/ioutil"
	"jaden/we-life/common"
	"jaden/we-life/errcode"
	"jaden/we-life/util"
	"math/rand"
	"net/http"
	"time"
)

//必须先检测session
func middleware_session_check() func(*Context) {
	return func(ctx *Context) {
		url := ctx.Request.URL.Path
		params := ctx.Params
		read, err := ctx.Request.GetBody()
		var body []byte
		if err == nil {
			body, _ = ioutil.ReadAll(read)
		}
		go func() {
			LogRequest(url, params, string(body))
		}()
		/*  session := ctx.Query("session")
		    if session == "" {
		            ctx.AbortWithStatusJSON(http.StatusOK, errcode.CACHE_NOT_EXIST.ToResult())
		            return
		    }*/
		//session是否存在校验
		/* _ ,err := common.Redis.Get("session").Result()
		   if err != nil{
		           ctx.AbortWithStatusJSON(http.StatusOK, errcode.SESSION_CHECK_ERROR.ToResult())
		           return
		   }*/
		ctx.Next()
	}
}

//记录普通的请求
func LogRequest(url string, params interface{}, bodyParams string) {
	common.Logger.InnerLogger.WithFields(logrus.Fields{
		"interface": url,
		"prams":     params,
		"body":      bodyParams,
		"time":      util.TimestampToDateTime(time.Now().Unix()),
	}).Info()
}

//程序崩溃中间件
func middleware_recover() func(*Context) {
	return func(context *Context) {
		defer func() {
			if err := recover(); err != nil {
				common.Logger.InnerLogger.WithFields(logrus.Fields{
					"interface": context.Request.URL.Path,
					"result":    fmt.Sprintf("panic : %v", errors.Wrap(err, 2).ErrorStack()),
					"alarmId":   rand.Int31(),
				}).Warn()
				context.JSON(http.StatusOK, errcode.SYSTEM_ERROR.ToResult())
			}
		}()
		context.Next()
	}
}
