package collection

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRoom struct {
	User_Identity    string `bson:user_identity`
	Room_Identity    string `bson:room_identity`
	Message_Identity string `bson:message_identity`
	Created_At       int32  `bson:created_at`
	Updated_At       int32  `bson:updated_at`
}

func (UserRoom) CollectionName() string {
	return "user_room"
}

func GetUserRoomByUIAndRI(userIdentity, roomIdentity string) (*UserRoom, error) {
	ur := new(UserRoom)
	err := Mongo.Collection(UserRoom{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"user_identity", userIdentity}, {"room_identity", roomIdentity}}).
		Decode(ur)
	if err != nil {
		return nil, err
	}
	return ur, nil
}

func GetUserRoomByIdentity(roomIdentity string) ([]*UserRoom, error) {
	urs := make([]*UserRoom, 0)
	cursor, err := Mongo.Collection(UserRoom{}.CollectionName()).Find(context.Background(), bson.D{{"room_identity", roomIdentity}})
	if err != nil {
		log.Println("[ERROR]: ", err)
		return nil, err
	}
	for cursor.Next(context.Background()) {
		ur := new(UserRoom)
		err := cursor.Decode(ur)
		if err != nil {
			log.Println("解析错误", err)
			return nil, err
		}
		urs = append(urs, ur)
	}
	return urs, nil
}
