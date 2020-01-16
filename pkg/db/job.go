package db

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"log"
)

func Save(name string, i interface{}) string {
	ConnectDb()
	collection := db.Collection(name)

	result, err := collection.InsertOne(ctx, i)
	if err != nil {
		log.Println(err)
		return ""
	}
	// TODO result.InsertedID to string
	fmt.Println(result.InsertedID)
	return ""
}

// return one
func FindOne(name string, id string, i interface{}) {

	collection := db.Collection(name)

	objectId := bson.ObjectIdHex(id)
	_ = collection.FindOne(ctx, objectId).Decode(i)
}

// return list
func Find(name string, id string, i interface{}) []interface{} {

	//collection := db.Collection(name)

	return nil

}

// Update i to name collection    ObjectId=id
func Update(name string, id string, i interface{}) {

	//collection := db.Collection(name)
	//objectId := bson.ObjectIdHex(id)
	//_ = collection.FindOneAndUpdate(ctx, objectId).Decode(i)


}

func Delete(name string, id string) {

	//collection := db.Collection(name)
	//objectId := bson.ObjectIdHex(id)
	//_ = collection.FindOneAndUpdate(ctx, objectId).Decode(i)


}
