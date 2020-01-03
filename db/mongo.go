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

func Db() {

	// 连接数据库
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // ctx
	opts := options.Client().ApplyURI("mongodb://10.160.202.95:27017")  // opts
	client, err := mongo.Connect(ctx, opts)                             // client
	if err != nil {
		log.Println(err)
		return
	}

	// 使用
	db := client.Database("KubeFate")   // database
	stu := db.Collection("fate") // collection

	// 插入数据
	xm1 := Fate{Name: "fate-10000", NameSpaces: "fate-10000", Version: "1.2.0",}
	_, err = stu.InsertOne(ctx, xm1)
	if err != nil {
		log.Println(err)
		return
	}
	xm2 := Fate{Name: "fate-9999", NameSpaces: "fate-9999", Version: "1.2.0",}
	_, err = stu.InsertOne(ctx, xm2)
	if err != nil {
		log.Println(err)
		return
	}

	// 查询数据  bson.D{}创建查询条件
	cur, err := stu.Find(ctx, bson.M{}) // find
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

type Fate struct {
	Name string `json:"name"`
	NameSpaces  string    `json:"namespaces"`
	Version  string `json:"version"`
}
