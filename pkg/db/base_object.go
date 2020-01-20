package db

import(
	"go.mongodb.org/mongo-driver/bson"
	"github.com/satori/go.uuid"
	"context"
	"log"
	"time"
)

type PersistentObject interface {
	Save(col string, po PersistentObject) (string, error)
	Find(col string, bo interface{}) (interface{}, error) 
	FindByUUID(col string, uuid string, bo interface{})  (interface{}, error)
	DeleteByUUID(col string, uuid string) error
}

type BaseObject struct {
	Uuid string `json:"UUID"`
}

func NewBaseObject() (*BaseObject){
	return &BaseObject{
		Uuid: uuid.NewV4().String(),
	}
}

func (baseObject BaseObject) Save(col string, po PersistentObject) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, _ := ConnectDb()
	collection := db.Collection(col)
	_, err := collection.InsertOne(ctx, po)
	if err != nil {
		log.Println(err)
		return "",err
	}
	return baseObject.Uuid, nil
}

func (baseObject BaseObject) Find(col string, bo interface{}) (interface{}, error) {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		db, _ := ConnectDb()
		collection := db.Collection(col)
		cur, err := collection.Find(ctx, bson.D{}) // find
		if err != nil {
			log.Println(err)
			return nil,err
		}
		defer cur.Close(ctx)
		var results []interface{}
		for cur.Next(ctx) {
			err = cur.Decode(&bo) 
			if err != nil {
				log.Println(err)
				return nil,err
			}
			results = append(results, bo)
		}
		return results,nil
	}

func (baseObject BaseObject) FindByUUID(col string, uuid string, bo interface{}) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, _ := ConnectDb()
	collection := db.Collection(col)
	filter := bson.M{"baseobject.uuid": uuid}
	var err error
	err = collection.FindOne(ctx, filter).Decode(&bo)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return bo, nil
}

func (baseObject BaseObject) DeleteByUUID(col string, uuid string) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, _ := ConnectDb()
	collection := db.Collection(col)
	filter := bson.M{"baseobject.uuid": uuid}
	deleteResult, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("Deleted %v documents in the %s collection\n", deleteResult.DeletedCount, col)
	return err
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