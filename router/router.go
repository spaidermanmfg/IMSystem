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

	//用户注册
	e.POST("/register", service.Register)

	//Routing group
	auth := e.Group("/u", middlewares.AuthCheck())

	//Certification through, continue
	auth.GET("/user/detail", service.UserDetail)

	auth.GET("/user/query", service.UserQuery)

	//发送，接受消息
	auth.GET("/talk/message", service.WebSocketMessage)

	//获取聊天记录
	auth.GET("/chat/list", service.ChatList)

	//添加用户
	auth.POST("/user/add", service.UserAdd)

	//删除好友
	auth.DELETE("/user/delete", service.UserDelete)
	return e
}
