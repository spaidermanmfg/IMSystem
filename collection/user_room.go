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
	Room_Type        int32  `bson:room_type`
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
		log.Println("[DB ERROR]", err)
		return nil, err
	}
	return ur, nil
}

func GetUserRoomByIdentity(roomIdentity string) ([]*UserRoom, error) {
	urs := make([]*UserRoom, 0)
	cursor, err := Mongo.Collection(UserRoom{}.CollectionName()).Find(context.Background(), bson.D{{"room_identity", roomIdentity}})
	if err != nil {
		log.Println("[DB ERROR]: ", err)
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

func JudgeUserIsFriend(userIdentity1, userIdentity2 string) (bool, error) {
	//TODO:查询user1的所有单聊房间
	cursor, err := Mongo.Collection(UserRoom{}.CollectionName()).
		Find(context.Background(), bson.D{{"user_identity", userIdentity1}, {"room_type", 1}})
	if err != nil {
		log.Println("[DB ERROR]: ", err)
		return false, err
	}

	userRooms1 := make([]string, 0)
	for cursor.Next(context.Background()) {
		ur := new(UserRoom)
		err := cursor.Decode(ur)
		if err != nil {
			log.Println("[ERROR]: ", err)
			return false, err
		}
		userRooms1 = append(userRooms1, ur.Room_Identity)
	}
	//TODO:判断user2的房间是否在user1的单聊房间列表中
	count, err := Mongo.Collection(UserRoom{}.CollectionName()).
		CountDocuments(context.Background(), bson.M{"user_identity": userIdentity2, "room_identity": bson.M{"$in": userRooms1}})
	if err != nil {
		log.Println("[DB ERROR]: ", err)
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, err
}

func InsertOneUserRoom(ur *UserRoom) error {
	_, err := Mongo.Collection(UserRoom{}.CollectionName()).
		InsertOne(context.Background(), ur)
	return err
}

func GetUserRoomIdentity(userIdentity1, userIdentity2 string) string {
	//TODO:查询user1的所有单聊房间
	cursor, err := Mongo.Collection(UserRoom{}.CollectionName()).
		Find(context.Background(), bson.D{{"user_identity", userIdentity1}, {"room_type", 1}})
	if err != nil {
		log.Println("[DB ERROR]: ", err)
		return ""
	}

	userRooms1 := make([]string, 0)
	for cursor.Next(context.Background()) {
		ur := new(UserRoom)
		err := cursor.Decode(ur)
		if err != nil {
			log.Println("[ERROR]: ", err)
			return ""
		}
		userRooms1 = append(userRooms1, ur.Room_Identity)
	}

	ur2 := new(UserRoom)
	//TODO:判断user2的房间是否在user1的单聊房间列表中
	err = Mongo.Collection(UserRoom{}.CollectionName()).
		FindOne(context.Background(), bson.M{"user_identity": userIdentity2, "room_identity": bson.M{"$in": userRooms1}}).
		Decode(ur2)

	if err != nil {
		log.Println("[DB ERROR]: ", err)
		return ""
	}
	return ur2.Room_Identity
}

func DeleteUserRoomByRoomIdentity(roomIdentity string) error {
	_, err := Mongo.Collection(UserRoom{}.CollectionName()).
		DeleteOne(context.Background(), bson.D{{"room_identity", roomIdentity}})
	if err != nil {
		log.Println("[DB ERROR]: ", err)
		return err
	}
	return nil
}
