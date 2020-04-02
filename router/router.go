package router

import (
	"jaden/we-life/controller"
)

func New() {
	NewServer()
}

func Run(port string) {
	RunServer(port)
}

func SetApi() {
	controller := new(controller.Handler)

	//通过id或者姓名查询用户
	Get("/api/getUserInfoByIdOrName", controller.GetUserByIdOrName)

	//涉及到首页相册列表
	Get("/api/directoryList", controller.GetDirectoryList)

	//相册中的动态列表
	Get("/api/news", controller.GetNews)

	//新建相册
	POST("/api/newDirectory", controller.AddDirectory)

	//发布动态
	POST("/api/publishNew", controller.PublishNew)

	//删除动态
	POST("/api/deleteNew", controller.DeleteNew)

	//修改相册名称
	POST("/api/updateDirectoryName", controller.UpdateDirectoryName)

	//用户登录
	POST("/api/userLogin", controller.Login)

	//获取用户详细信息
	Get("/api/getUserInfo", controller.GetUserInfoById)

	//更新用户信息
	POST("/api/updateUserInfo", controller.UpdateUserInfoById)

	//用户请求绑定另一个用户
	POST("/api/requestBind", controller.BindUserRelaRequest)

	//GET 获取请求绑定的用户列表
	Get("/api/user-list/request-bind", controller.UserBindListRequest)

	//同意建议与指定用户的关系(收到邀请建立的前提下才会建立)
	POST("/api/agreeBind", controller.UserAgreeBind)
}
