package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
var db *mongo.Database

func ConnectDb() {

	// 连接数据库
	// ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	opts := options.Client().ApplyURI("mongodb://root:root@localhost:27017")  // opts
	client, err := mongo.Connect(ctx, opts)                             // client
	if err != nil {
		log.Println(err)
		return
	}

	// 使用
	db = client.Database("KubeFate") // database
}
func InsertFate(fate Fate) {
	collection := db.Collection("fate")      // collection


	insertResult, err := collection.InsertOne(ctx, fate)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func FindFate(){
	// 查询数据  bson.D{}创建查询条件
	collection := db.Collection("fate")
	cur, err := collection.Find(ctx, bson.M{}) // find
	if err != nil {
		log.Println(err)
		return
	}

	// 延时关闭游标
	defer cur.Close(ctx)

	//查询结果
	for cur.Next(ctx) {

		// 可以decode到bson.M  也就是一个map[string]interface{}中
		// 也可以直接decode到一个对象中
		s := &Fate{}
		var result bson.M
		err := cur.Decode(&result) // decode 到map
		err = cur.Decode(s)        // decode 到对象
		if err != nil {
			log.Println(err)
			return
		}

		// do something with result....
		// 可以将map 或对象序列化为json
		//js ,_:=json.Marshal(result)
		//json.Unmarshal(js,s) //反学序列化回来
		fmt.Println(s)
	}

}

type FateCluster struct {
	Name       string `json:"name"`
	NameSpaces string `json:"namespaces"`
	Version    string `json:"version"`
	PartyID    string `json:"PartyID"`
	Chart      string `json:"chart"`
}

type Chart struct {
	version	    string `json:"version"`
	value       string `json:"value"`
	template    string `json:"template"`
}
