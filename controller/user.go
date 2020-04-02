package controller

import (
	"github.com/gin-gonic/gin"
	"io"
	"jaden/we-life/domain/photo"
	"jaden/we-life/domain/user"
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

type UserController struct {
	web.BaseController
}

//登录
func (c UserController) Login(ctx *gin.Context) *entity.Result {
	userInfo := new(UserLoginRequest)
	if err := ctx.BindJSON(userInfo); err != nil {
		return c.ReturnErrCode(errcode.PARAM_ERROR.AddMsg("缺少必要的参数信息"))
	}
	session, err := user.NewService().Login(userInfo.UserId, userInfo.Password)
	if err != nil {
		return c.ReturnErrCode(err)
	}
	return c.ReturnSuccess(session)
}

func (c UserController) UpdateUserInfoById(ctx *gin.Context) *entity.Result {
	forms, err := ctx.MultipartForm()
	if err != nil {
		return c.ReturnErrCode(errcode.PARAM_ERROR.AddMsg("更新用户信息失败:" + err.Error()))
	}
	ids, ok := forms.Value["userId"]
	if !ok {
		return c.ReturnErrCode(errcode.PARAM_ERROR.AddMsg("缺少必填参数userId"))
	}
	us := new(user.User)
	userId, err := strconv.Atoi(ids[0])
	if err != nil {
		return c.ReturnErrCode(errcode.PARAM_ERROR.AddMsg("用户id解析失败"))
	}
	us.Id = userId
	ns, ok := forms.Value["name"]
	if ok {
		us.Name = ns[0]
	}
	bts, ok := forms.Value["birthday"]

	if ok {
		birthday, err := strconv.Atoi(bts[0])
		if err != nil {
			return c.ReturnErrCode(errcode.PARAM_ERROR.AddMsg("生日解析失败"))
		}
		us.Birthday = birthday
	}

	fhs, ok := forms.File["avator"]
	if ok {
		v := fhs[0]
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
		us.Avatar = ufn.FilePath
	}
	ph, ok := forms.Value["phone"]
	if ok {
		us.Phone = ph[0]
	}
	ps, ok := forms.Value["password"]
	if ok {
		us.Password = ps[0]
	}
	if err := user.NewService().UpdateUserInfoById(us); err != nil {
		return c.ReturnErrCode(err)
	}
	return c.ReturnSuccess("")
}

func (c UserController) GetUserInfoById(ctx *gin.Context) *entity.Result {
	userId := ctx.Query("userId ")
	id, err := strconv.Atoi(userId)
	if err != nil {
		return c.ReturnErrCode(errcode.PARAM_ERROR)
	}
	if userInfo, err := user.NewService().GetUserDetailInfoById(id); err != nil {
		return c.ReturnErrCode(err)
	} else {
		return c.ReturnSuccess(userInfo)
	}
}

func (c UserController) BindUserRelaRequest(ctx *gin.Context) *entity.Result {
	userBind := new(UserBindRequest)
	if err := ctx.BindJSON(userBind); err != nil {
		return c.ReturnErrCode(errcode.PARAM_ERROR.AddMsg("参数错误"))
	}
	err := user.NewService().RequestBindUser(userBind.UserId, userBind.BindUserId)
	if err != nil {
		return c.ReturnErrCode(err)
	}
	return c.ReturnSuccess("")
}

func (c UserController) UserAgreeBind(ctx *gin.Context) *entity.Result {

	userBind := new(UserBindRequest)
	if err := ctx.BindJSON(userBind); err != nil {
		return c.ReturnErrCode(errcode.PARAM_ERROR.AddMsg("参数错误"))
	}
	err := user.NewService().AgreeBindUser(userBind.UserId, userBind.BindUserId)
	if err != nil {
		return c.ReturnErrCode(err)
	}
	return c.ReturnSuccess("")

}

func (c UserController) UserBindListRequest(ctx *gin.Context) *entity.Result {

	if userId, err := strconv.Atoi(ctx.Query("userId")); err != nil {
		return c.ReturnErrCode(errcode.PARAM_ERROR.AddMsg("参数错误"))
	} else {
		users, err := user.NewService().UserBindListRequest(userId)
		if err != nil {
			return c.ReturnErrCode(err)
		}
		return c.ReturnSuccess(users)
	}

}

func (c UserController) GetUserByIdOrName(ctx *gin.Context) *entity.Result {
	userId := ctx.Query("id")
	name := ctx.DefaultQuery("name", "")
	id, err := strconv.Atoi(userId)
	if err != nil {
		return c.ReturnErrCode(errcode.PARAM_ERROR)
	}
	if userInfo, err := user.NewService().GetUserByIdOrName(id, name); err != nil {
		return c.ReturnErrCode(err)
	} else {
		return c.ReturnSuccess(userInfo)
	}

}
