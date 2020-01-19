package db

import(
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"fmt"
	"log"
	"time"
)

type BaseObject interface {
	getCollection() string
}

type BaseStruct struct {
	Uuid string
}

type TestBaseStruct struct {
	BaseStruct
	Name string
}

func (baseStruct BaseStruct) Save(col string,i interface{}) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := DB.db.Collection(col)
	_, err := collection.InsertOne(ctx, i)
	if err != nil {
		log.Println(err)
		return "",err
	}
	return baseStruct.Uuid, nil
}
func (baseStruct BaseStruct) Find(col string) ([]*BaseStruct, error) {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		collection := DB.db.Collection(col)
		cur, err := collection.Find(ctx, bson.M{}) // find
		if err != nil {
			log.Println(err)
			return nil,err
		}
		defer cur.Close(ctx)
		baseStructs := []*BaseStruct{}
		for cur.Next(ctx) {
			s := new(BaseStruct)
			var result bson.M
			err := cur.Decode(&result) // decode 到map
			err = cur.Decode(s)        // decode 到对象
			if err != nil {
				log.Println(err)
				return nil,err
			}
			fmt.Println(s)
			baseStructs = append(baseStructs, s)
		}
		return baseStructs,nil
	}




// func Save(baseObject BaseObject) (string, error){
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	collection := DB.db.Collection(baseObject.getCollection())      // collection

// 	_, err := collection.InsertOne(ctx, baseObject)
// 	if err != nil {
// 		log.Println(err)
// 		return "",err
// 	}
// 	return "",nil
// }

// func Find(baseObject BaseObject) ([]*BaseObject, error) {
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	collection := DB.db.Collection(baseObject.getCollection())
// 	cur, err := collection.Find(ctx, bson.M{}) // find
// 	if err != nil {
// 		log.Println(err)
// 		return nil,err
// 	}
// 	defer cur.Close(ctx)
// 	baseObjects := []*BaseObject{}
// 	for cur.Next(ctx) {

// 		// 可以decode到bson.M  也就是一个map[string]interface{}中
// 		// 也可以直接decode到一个对象中
// 		s := new(BaseObject)
// 		var result bson.M
// 		err := cur.Decode(&result) // decode 到map
// 		err = cur.Decode(s)        // decode 到对象
// 		if err != nil {
// 			log.Println(err)
// 			return nil,err
// 		}

// 		// do something with result....
// 		// 可以将map 或对象序列化为json
// 		//js ,_:=json.Marshal(result)
// 		//json.Unmarshal(js,s) //反学序列化回来
// 		// fmt.Println(s)
// 		baseObjects = append(baseObjects, s)
// 	}
// 	fmt.Println(baseObjects)
// 	return baseObjects,nil
// }

// func FindByUUID(uuid string) (*BaseObject, error)

// func UpdateByUUID(uuid string, curd BaseObject) (*BaseObject, error)

// func Delete(uuid string) error