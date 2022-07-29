package test

import (
	"IMSystem/collection"
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//test mongodb driver
func TestFindOne(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:123456@114.55.237.245:27017"))

	if err != nil {
		log.Println("连接到数据库失败", err)
	}

	db := client.Database("im")
	ur := new(collection.UserRoom)
	err = db.Collection(ur.CollectionName()).FindOne(context.Background(), bson.D{}).Decode(ur)
	if err != nil {
		log.Println("解析错误", err)
	}
	fmt.Println("ub===>", ur)
}
