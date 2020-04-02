package web

import (
	"jaden/we-life/entity"
)

type BaseController struct {
}

func (b BaseController) ReturnSuccess(data interface{}) *entity.Result {
	return &entity.Result{Status: 0, Data: data, Msg: "success"}
}

func (b BaseController) ReturnErrCode(errCode *entity.ErrCode) *entity.Result {
	return errCode.ToResult()
}
