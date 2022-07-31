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
	code := tools.GenerateRandCode()
	err = tools.SendCode(email, code)
	if err != nil {
		log.Println("[ERROR]: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "验证码发送失败",
		})
		return
	}

	if err = collection.RDB.Set(define.RegisterPrefix+email, code, define.ExpireTime).Err(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "验证码保存失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "验证码发送成功",
	})

}

func UserQuery(c *gin.Context) {
	account := c.Query("account")
	if account == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户名不能为空",
		})
		return
	}

	userBasic, err := collection.GetUserBasicByAccounter(account)
	if err != nil {
		log.Printf("[DB ERROR]: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "系统错误",
		})
		return
	}

	uc := c.MustGet("userClaims").(*tools.UserClaims)
	flag, err := collection.JudgeUserIsFriend(uc.Identity, userBasic.Identity)
	if err != nil {
		log.Printf("[DB ERROR]: %v\n", err)
		return
	}
	data := collection.UserQueryResult{
		Nickname: userBasic.Nickname,
		Sex:      userBasic.Sex,
		Email:    userBasic.Email,
		IsFriend: flag,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "select success",
		"data":    data,
	})

}

func UserAdd(c *gin.Context) {
	account := c.PostForm("account")
	if account == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户名不能为空",
		})
		return
	}
	ub, err := collection.GetUserBasicByAccounter(account)
	if err != nil {
		log.Printf("[DB ERROR]: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "系统错误",
		})
		return
	}

	uc := c.MustGet("userClaims").(*tools.UserClaims) //获取当前用户信息
	//判断当前好友关系
	if ok, _ := collection.JudgeUserIsFriend(ub.Identity, uc.Identity); ok {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "已为好友，不可重复添加",
		})
		return
	}

	//保存房间记录
	rb := &collection.RoomBasic{
		Identity:      tools.GenerateUUID(),
		User_Identity: uc.Identity,
		Created_At:    int32(time.Now().Unix()),
		Updated_At:    int32(time.Now().Unix()),
	}

	err = collection.InsertOneRoomBasic(rb)
	if err != nil {
		log.Printf("[DB ERROR]: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "系统错误",
		})
		return
	}

	//保存用户房间关联记录
	ur1 := &collection.UserRoom{
		User_Identity: uc.Identity,
		Room_Identity: rb.Identity,
		Room_Type:     1,
		Created_At:    int32(time.Now().Unix()),
		Updated_At:    int32(time.Now().Unix()),
	}

	err = collection.InsertOneUserRoom(ur1)
	if err != nil {
		log.Printf("[DB ERROR]: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "系统错误",
		})
		return
	}

	ur2 := &collection.UserRoom{
		User_Identity: ub.Identity,
		Room_Identity: rb.Identity,
		Room_Type:     1,
		Created_At:    int32(time.Now().Unix()),
		Updated_At:    int32(time.Now().Unix()),
	}

	err = collection.InsertOneUserRoom(ur2)
	if err != nil {
		log.Printf("[DB ERROR]: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "系统错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "好友添加成功",
	})

}

//删除好友
func UserDelete(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户名不能为空",
		})
		return
	}

	uc := c.MustGet("userClaims").(*tools.UserClaims)
	roomIdentity := collection.GetUserRoomIdentity(identity, uc.Identity)
	if roomIdentity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "非好友关系, 无需删除",
		})
		return
	}

	//删除user_room关联关系
	err := collection.DeleteUserRoomByRoomIdentity(roomIdentity)
	if err != nil {
		log.Printf("[DB ERROR]: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "系统错误",
		})
		return
	}

	//删除room_basic
	err = collection.DeleteRoomBasicByRoomIdentity(roomIdentity)
	if err != nil {
		log.Printf("[DB ERROR]: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "系统错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除好友成功",
	})

}
