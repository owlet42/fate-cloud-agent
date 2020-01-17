package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

type kubeFATEDatabase struct{
	db            *mongo.Database
	mongoUrl      string
	mongoUsername string
	mongoPassword string
}

var DB kubeFATEDatabase = kubeFATEDatabase{mongoUrl:"localhost:27017", mongoUsername:"root", mongoPassword:"root"}

func ConnectDb() error {

	// 连接数据库
	// ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	opts := options.Client().ApplyURI("mongodb://"+DB.mongoUsername+":"+DB.mongoPassword+"@"+DB.mongoUrl)  // opts
	client, err := mongo.Connect(ctx, opts)                             // client
	if err != nil {
		log.Println(err)
		return err
	}

	// 使用
	DB.db = client.Database("KubeFate") // database
	return nil
}

func Disconnect() error {
	return nil
}

func Ping() error {
	return nil
}
