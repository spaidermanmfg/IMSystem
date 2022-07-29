package router

import (
	"IMSystem/middlewares"
	"IMSystem/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	e := gin.Default()
	e.POST("/login", service.Login)

	e.POST("/send/code", service.SendCode)

	//Routing group
	auth := e.Group("/u", middlewares.AuthCheck())

	//Certification through, continue
	auth.GET("/user/detail", service.UserDetail)
	return e
}
