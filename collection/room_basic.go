package collection

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type RoomBasic struct {
	Identity      string `bson:identity`
	Number        string `bson:number`
	Name          string `bson:name`
	Info          string `bson:info`
	User_Identity string `bson:info`
	Created_At    int32  `bson:created_at`
	Updated_At    int32  `bson:updated_at`
}

func (RoomBasic) CollectionName() string {
	return "room_basic"
}

func InsertOneRoomBasic(rb *RoomBasic) error {
	_, err := Mongo.Collection(RoomBasic{}.CollectionName()).
		InsertOne(context.Background(), rb)
	return err
}

func DeleteRoomBasicByRoomIdentity(roomIdentity string) error {
	_, err := Mongo.Collection(RoomBasic{}.CollectionName()).
		DeleteOne(context.Background(), bson.D{{"identity", roomIdentity}})
	if err != nil {
		log.Println("[DB ERROR]: ", err)
		return err
	}
	return nil
}
