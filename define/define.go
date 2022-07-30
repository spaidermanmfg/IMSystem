package define

import "time"

//var MailPassword = os.Getenv("MailPassword")
var MailPassword = "VWQLLBKPMVQHCIRE"

type Message struct {
	Message       string `json:message`
	Room_Identity string `json:room_identity`
}

var RegisterPrefix = "TOKEN_"     //前缀
var ExpireTime = time.Second * 60 //验证码过期时间
