package collection

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type UserBasic struct {
	Identity   string `bson:identity`
	Account    string `bson:account`
	Password   string `bson:password`
	Nickname   string `bson:nickname`
	Sex        int32  `bson:sex`
	Email      string `bson:email`
	Avatar     string `bson:avatar`
	Created_At int32  `bson:created_at`
	Updated_At int32  `bson:updated_at`
}

func (UserBasic) CollectionName() string {
	return "user_basic"
}

//select user by account and password
func GetUserBasicByAP(account, password string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(
			context.Background(),
			bson.D{{"account", account}, {"password", password}}).
		Decode(ub)
	return ub, err
}

//select user details by identity
func GetUserBasicByIdentity(identity string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(
			context.Background(),
			bson.D{{"identity", identity}}).
		Decode(ub)
	return ub, err
}

//select email whether or not repeat
func GetUserBasicByEmail(email string) (int64, error) {
	num, err := Mongo.Collection(UserBasic{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"email", email}})
	return num, err
}

func GetUserBasicByAccount(account string) (int64, error) {
	num, err := Mongo.Collection(UserBasic{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"account", account}})
	return num, err
}

//保存注册信息
func InsertOneUserBasic(ub UserBasic) error {
	_, err := Mongo.Collection(UserBasic{}.CollectionName()).
		InsertOne(context.Background(), ub)
	return err
}
