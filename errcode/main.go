package errcode

import (
	"jaden/we-life/entity"
)

//
// 定义服务内部错误码
//
var (
	SUCCESS             = entity.NewErrCode(0, "ok")
	SYSTEM_ERROR        = entity.NewErrCode(-1, "系统异常")
	PARAM_ERROR         = entity.NewErrCode(1001, "参数错误")
	PARAM_PARSE_ERROR   = entity.NewErrCode(1002, "参数解析错误")
	HEADER_ERROR        = entity.NewErrCode(1003, "请求头信息错误")
	RATE_ERROR          = entity.NewErrCode(1014, "请求过于频繁")
	DATA_ERROR          = entity.NewErrCode(1016, "业务调用失败")
	DATA_NOT_EXIST      = entity.NewErrCode(1020, "数据记录不存在")
	CACHE_NOT_EXIST     = entity.NewErrCode(1020, "缓存记录不存在")
	CACHE_ERROR         = entity.NewErrCode(1021, "缓存异常")
	RES_SWITCH_ERROR    = entity.NewErrCode(3001, "响应结果转换异常")
	EMPTY_DATA          = entity.NewErrCode(6003, "数据为空")
	SESSION_CHECK_ERROR = entity.NewErrCode(-2, "身份验证失败")
)
