package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"jaden/we-life/domain/photo"
	"jaden/we-life/entity"
	"jaden/we-life/errcode"
	"jaden/we-life/util"
	"jaden/we-life/web"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type PhotoController struct {
	web.BaseController
}

func (c PhotoController) GetDirectoryList(ctx *gin.Context) *entity.Result {
	userId := ctx.Query("userId")
	session := ctx.Query("session")
	id, err := strconv.Atoi(userId)
	if err != nil {
		return c.ReturnErrCode(errcode.PARAM_ERROR)
	}
	if session == "" {
		return c.ReturnErrCode(errcode.PARAM_ERROR)
	}
	if dirs, err := photo.NewService().GetDirList(id); err != nil {
		return c.ReturnErrCode(err)
	} else {
		return c.ReturnSuccess(dirs)
	}
}

func (c PhotoController) GetFileList(ctx *gin.Context) *entity.Result {
	nq := new(NewRequest)
	if err := ctx.BindQuery(nq); err != nil {
		return c.ReturnErrCode(errcode.PARAM_ERROR)
	}

	if files, err := photo.NewService().GetFileList(nq.UserId, nq.DirectoryId); err != nil {
		return c.ReturnErrCode(err)
	} else {
		return c.ReturnSuccess(files)
	}
}

func (c PhotoController) AddDirectory(ctx *gin.Context) *entity.Result {
	adq := new(AddDirRequest)
	if err := ctx.BindJSON(adq); err != nil {
		return c.ReturnErrCode(errcode.PARAM_ERROR)
	}

	if err := photo.NewService().AddDirectory(adq.UserID, adq.DirName, 0); err != nil {
		return c.ReturnErrCode(err)
	} else {
		return c.ReturnSuccess("")
	}
}

func (c PhotoController) DeleteNew(ctx *gin.Context) *entity.Result {
	obj := new(DeleteNewReqest)
	if err := ctx.BindJSON(obj); err != nil {
		return c.ReturnErrCode(errcode.PARAM_ERROR.AddMsg("参数解析失败"))
	}

	if err := photo.NewService().DeleteNew(obj.NewId); err != nil {
		return c.ReturnErrCode(err)
	} else {
		return c.ReturnSuccess("")
	}

}

func (c PhotoController) UpdateDirectoryName(ctx *gin.Context) *entity.Result {
	obj := new(UpdateDirNameRequest)
	err := ctx.BindJSON(obj)
	if err != nil {
		return c.ReturnErrCode(errcode.PARAM_ERROR.AddMsg("参数绑定失败"))
	}

	if err := photo.NewService().UpdateDirName(obj.DirId, obj.NewTitle); err != nil {
		return c.ReturnErrCode(err)
	} else {
		return c.ReturnSuccess("")
	}

}

func (c PhotoController) PublishNew(ctx *gin.Context) *entity.Result {

	mp, err := ctx.MultipartForm()
	if err != nil {
		return c.ReturnErrCode(errcode.PARAM_ERROR)
	}
	userId := mp.Value["userId"][0]
	address := mp.Value["address"][0]
	content := mp.Value["content"][0]
	directoryId := mp.Value["directoryId"][0]
	id, err := strconv.Atoi(userId)
	if err != nil {
		return c.ReturnErrCode(errcode.PARAM_ERROR.AddMsg("用户id解析失败"))
	}
	dirId, err := strconv.Atoi(directoryId)
	if err != nil {
		return c.ReturnErrCode(errcode.PARAM_ERROR.AddMsg("目录id解析失败"))
	}
	if id == 0 {
		return c.ReturnErrCode(errcode.PARAM_ERROR)
	}
	fileHeaders, findFile := mp.File["file"]
	fmt.Printf("userId=%s,address=%s,content=%s,fileHeaders=%v", userId, address, content, fileHeaders)
	var ufns []photo.UploadFileEntity
	if findFile {

		for _, v := range fileHeaders {
			ufn := new(photo.UploadFileEntity)
			fileName := v.Filename
			dotLen := strings.LastIndex(fileName, ".")
			suffix := fileName[dotLen:]
			ufn.FileName = fileName
			nowTimeUnix := strconv.Itoa(int(time.Now().UnixNano()))
			ufn.FilePath = util.GetRandString(16) + nowTimeUnix[len(nowTimeUnix)-8:] + suffix
			file, err := v.Open()
			if err != nil {
				return c.ReturnErrCode(errcode.SYSTEM_ERROR.AddMsg("文件开启失败"))
			}
			defer file.Close()
			sysType := runtime.GOOS
			var outputFilePath string
			if sysType == "linux" {
				// LINUX系统
				outputFilePath = "/data/www/images/" + ufn.FilePath
			}

			if sysType == "windows" {
				// windows系统
				outputFilePath = "O://" + ufn.FilePath
			}
			writer, err := os.OpenFile(outputFilePath, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				return c.ReturnErrCode(errcode.SYSTEM_ERROR.AddMsg("文件开始失败..." + err.Error()))
			}
			defer writer.Close()
			io.Copy(writer, file)
			ufns = append(ufns, *ufn)
		}
	}
	entity := &photo.PublishEntity{
		UserId:           id,
		DirId:            dirId,
		Content:          content,
		Address:          address,
		UploadFileEntity: ufns,
	}

	if err := photo.NewService().PublishNew(entity); err != nil {
		return c.ReturnErrCode(err)
	} else {
		return c.ReturnSuccess("")
	}
}
