package controller

import (
	"github.com/gin-gonic/gin"
	"jaden/we-life/domain/photo"
	"jaden/we-life/entity"
	"jaden/we-life/errcode"
	"jaden/we-life/web"
	"strconv"
)

type NewsController struct {
	web.BaseController
}

func (c UserController) GetNews(ctx *gin.Context) *entity.Result {
	userId := ctx.Query("userId")
	dId := ctx.Query("directoryId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		return c.ReturnErrCode(errcode.PARAM_ERROR)
	}
	dirId, err := strconv.Atoi(dId)
	if err != nil {
		return c.ReturnErrCode(errcode.PARAM_ERROR)
	}
	if newsInfo, err := photo.NewService().GetFileList(id, dirId); err != nil {
		return c.ReturnErrCode(err)
	} else {
		return c.ReturnSuccess(newsInfo)
	}

}
