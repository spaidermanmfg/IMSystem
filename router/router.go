package router

import (
	"IMSystem/middlewares"
	"IMSystem/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	e := gin.Default()
	//登录
	e.POST("/login", service.Login)

	//使用邮箱发送验证码
	e.POST("/send/code", service.SendCode)

	e.POST("/register", service.Register)

	//Routing group
	auth := e.Group("/u", middlewares.AuthCheck())

	//Certification through, continue
	auth.GET("/user/detail", service.UserDetail)

	//发送，接受消息
	auth.GET("/talk/message", service.WebSocketMessage)

	//获取聊天记录
	auth.GET("/chat/list", service.ChatList)
	return e
}
