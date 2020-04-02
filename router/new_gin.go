package router

import (
	"github.com/gin-gonic/gin"
	"jaden/we-life/entity"
	"net/http"
)

type Context = gin.Context

type MiddleFunc func(c *Context)

type HandlerFunc func(c *Context) *entity.Result

var engine *gin.Engine

func NewServer() {
	gin.SetMode(gin.DebugMode)
	engine = gin.New()
	addMidware(middleware_recover())
	addMidware(middleware_session_check())
}

func RunServer(port string) {
	engine.Run("0.0.0.0:" + port)
}

func handleFuncTransfer(handlers []HandlerFunc) []gin.HandlerFunc {
	ginHandlers := make([]gin.HandlerFunc, len(handlers))
	for index := range handlers {
		hasJsonResp := false
		ginHandlers[index] = func(c *gin.Context) {
			result := handlers[index](c)
			if result != nil {
				c.JSON(http.StatusOK, result)
				hasJsonResp = true
			}
		}
		if hasJsonResp {
			break
		}
	}
	return ginHandlers
}

func addMidware(f MiddleFunc) {
	engine.Use(gin.HandlerFunc(f))
}

func Get(path string, handlers ...HandlerFunc) {
	engine.GET(path, handleFuncTransfer(handlers)...)
}
func POST(relatePath string, handlers ...HandlerFunc) {
	engine.POST(relatePath, handleFuncTransfer(handlers)...)
}

func DELETE(relatePath string, handlers ...HandlerFunc) {
	engine.DELETE(relatePath, handleFuncTransfer(handlers)...)
}

func PUT(relatePath string, handlers ...HandlerFunc) {
	engine.PUT(relatePath, handleFuncTransfer(handlers)...)
}
