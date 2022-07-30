package service

import (
	"IMSystem/collection"
	"IMSystem/define"
	"IMSystem/tools"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	code := c.PostForm("code")
	email := c.PostForm("email")
	account := c.PostForm("account")
	password := c.PostForm("password")
	if code == "" || email == "" || account == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数校验失败",
		})
		return
	}

	//TODO:判断验证码是否正确
	s, err := collection.RDB.Get(define.RegisterPrefix + email).Result()
	if err != nil {
		log.Println("[ERROR]: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "验证码不正确",
		})
		return
	}

	if s != code {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "验证码不正确",
		})
		return
	}

	//TODO:判断账号是否唯一
	act, err := collection.GetUserBasicByAccount(account)
	if err != nil {
		log.Println("[ERROR]: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "系统错误",
		})
		return
	}

	if act > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "账号已经注册",
		})
		return
	}

	//TODO:进行注册操作
	ub := &collection.UserBasic{
		Identity:   tools.GenerateUUID(),
		Account:    account,
		Password:   tools.GetMd5(password),
		Email:      email,
		Created_At: int32(time.Now().Unix()),
		Updated_At: int32(time.Now().Unix()),
	}

	err = collection.InsertOneUserBasic(*ub)
	if err != nil {
		log.Println("[ERROR]: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "注册失败",
		})
		return
	}

	//TODO:注册成功后自动登录
	token, err := tools.GenerateToken(ub.Identity, ub.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "系统错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Login success.",
		"data": gin.H{
			"token": token,
		},
	})
}
