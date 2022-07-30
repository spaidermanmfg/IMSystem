package service

import (
	"IMSystem/collection"
	"IMSystem/define"
	"IMSystem/tools"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Upgrade upgrades the HTTP server connection to the WebSocket protocol.
var upgrade = websocket.Upgrader{}
var wc = make(map[string]*websocket.Conn)

func WebSocketMessage(c *gin.Context) {
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("[ERROR]: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "HTTP服务器连接升级到WebSocket协议失败",
		})
		return
	}
	defer conn.Close()

	uc := c.MustGet("userClaims").(*tools.UserClaims)
	wc[uc.Identity] = conn
	for {
		msg := new(define.Message)
		err := conn.ReadJSON(msg)
		if err != nil {
			log.Println("[ERROR]: ", err)
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "消息接受失败",
			})
			return
		}
		// TODO: 判断用户是否属于同一房间
		ur, err := collection.GetUserRoomByUIAndRI(uc.Identity, msg.Room_Identity)
		if err != nil {
			log.Printf("UserIdentity: %v, RoomIdentity: %v Not Found.\n", uc.Identity, msg.Room_Identity)
			return
		}
		fmt.Println("userroom======>", ur)

		// TODO：保存消息
		mb := &collection.MessageBasic{
			//Identity      :,
			User_Identity: uc.Identity,
			Room_Identity: msg.Room_Identity,
			Data:          msg.Message,
			Created_At:    int32(time.Now().Unix()),
			Updated_At:    int32(time.Now().Unix()),
		}
		err = collection.InsertOneMessageBasic(*mb)
		if err != nil {
			log.Println("[ERROR]: 消息保存失败！", err)
			return
		}

		// TODO:获取在同一房间的所有用户
		userRooms, err := collection.GetUserRoomByIdentity(msg.Room_Identity)
		if err != nil {
			log.Printf("RoomIdentity: %v Not Exist.", msg.Room_Identity)
			return
		}

		for _, v := range userRooms {
			fmt.Println("userIdentity====>", v.User_Identity)
			if conn, ok := wc[v.User_Identity]; ok {
				err := conn.WriteMessage(websocket.TextMessage, []byte(msg.Message))
				if err != nil {
					log.Println("[ERROR]: ", err)
					return
				}
			}
		}

	}
}
