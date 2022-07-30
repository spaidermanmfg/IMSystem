package collection

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageBasic struct {
	//Identity      string `bson:identity`
	User_Identity string `bson:user_identity`
	Room_Identity string `bson:room_identity`
	Data          string `bson:data`
	Created_At    int32  `bson:created_at`
	Updated_At    int32  `bson:updated_at`
}

func (MessageBasic) CollectionName() string {
	return "message_basic"
}

//保存消息
func InsertOneMessageBasic(mb MessageBasic) error {
	_, err := Mongo.Collection(MessageBasic{}.CollectionName()).
		InsertOne(context.Background(), mb)
	return err
}

//聊天记录
func GetChatListByRoomIdentity(roomIdentity string, limit, skip *int64) ([]*MessageBasic, error) {
	data := make([]*MessageBasic, 0)

	cursor, err := Mongo.Collection(MessageBasic{}.CollectionName()).
		Find(context.Background(), bson.D{{"room_identity", roomIdentity}}, &options.FindOptions{
			Limit: limit,
			Skip:  skip,
			Sort:  bson.D{{"created_at", -1}},
		})

	if err != nil {
		log.Println("[ERROR]: ", err)
		return nil, err
	}

	for cursor.Next(context.Background()) {
		mb := new(MessageBasic)
		err := cursor.Decode(mb)
		if err != nil {
			log.Println("[ERROR]: ", err)
			return nil, err
		}
		data = append(data, mb)
	}

	return data, err
}
