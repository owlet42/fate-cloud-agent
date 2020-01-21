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

var db kubeFATEDatabase = kubeFATEDatabase{mongoUrl:"127.0.0.1:27017", mongoUsername:"root", mongoPassword:"root"}

func ConnectDb() (*mongo.Database, error) {
	if db.db == nil {
		opts := options.Client().ApplyURI("mongodb://"+db.mongoUsername+":"+db.mongoPassword+"@"+db.mongoUrl)  // opts
		client, err := mongo.Connect(ctx, opts)   // client
		if err != nil {
			log.Println(err)
			return nil, err
		}
		db.db = client.Database("KubeFate")
	}
	
	return db.db, nil
}

func Disconnect() error {
	return nil
}

func Ping() error {
	return nil
}
