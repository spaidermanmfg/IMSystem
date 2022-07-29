package service

import (
	"IMSystem/collection"
	"IMSystem/tools"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	account := c.PostForm("account")
	password := c.PostForm("password")
	if account == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户名或密码不能为空",
		})
		return
	}

	ub, err := collection.GetUserBasicByAP(account, tools.GetMd5(password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户名或密码错误",
		})
		return
	}

	//generate token
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

func UserDetail(c *gin.Context) {
	value, _ := c.Get("userClaims")
	uc := value.(*tools.UserClaims)
	ub, err := collection.GetUserBasicByIdentity(uc.Identity)
	if err != nil {
		log.Printf("[DB ERROR]:%\v", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户数据查询异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "用户数据查询成功",
		"data":    ub,
	})
}

func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "邮箱名不能为空",
		})
		return
	}

	count, err := collection.GetUserBasicByEmail(email)
	if err != nil {
		log.Printf("[DB ERROR]: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "系统错误",
		})
		return
	}

	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "当前邮箱已被占用",
		})
		return
	}

	err = tools.SendCode(email, "777777")
	if err != nil {
		log.Println("[ERROR]: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "验证码发送失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "验证码发送成功",
	})

}
