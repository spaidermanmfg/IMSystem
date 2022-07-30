package service

import (
	"IMSystem/collection"
	"IMSystem/tools"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ChatList(c *gin.Context) {
	roomIdentity := c.Query("room_identity")
	if roomIdentity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "房间号为空",
		})
		return
	}

	//判读当前用户是否属于该房间
	uc := c.MustGet("userClaims").(*tools.UserClaims)
	_, err := collection.GetUserRoomByUIAndRI(uc.Identity, roomIdentity)
	if err != nil {
		log.Println("[ERROR]: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "非法访问",
		})
		return
	}

	//获取聊天记录
	pageIndex, _ := strconv.ParseInt(c.Query("page_index"), 10, 32)
	pageSize, _ := strconv.ParseInt(c.Query("page_size"), 10, 32)
	skip := (pageIndex - 1) * pageSize
	mb, err := collection.GetChatListByRoomIdentity(roomIdentity, &pageSize, &skip)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "系统异常" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取聊天记录成功",
		"data": gin.H{
			"list": mb,
		},
	})

}
