package db

import(
	"go.mongodb.org/mongo-driver/bson"
	"encoding/json"
	"context"
	"log"
	"time"
)

type Repository interface {
	getCollection() string
	FromBson(m *bson.M) interface{}
	GetUuid() string
}

func Save(repository Repository) (string, error){
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, _ := ConnectDb()
	collection := db.Collection(repository.getCollection())
	_, err := collection.InsertOne(ctx, repository)
	if err != nil {
		log.Println(err)
		return "",err
	}
	return repository.GetUuid(),nil
}

func Find(repository Repository) ([]interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, _ := ConnectDb()
	collection := db.Collection(repository.getCollection())
	cur, err := collection.Find(ctx, bson.M{}) // find
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cur.Close(ctx)
	var persistents []interface{}
	for cur.Next(ctx) {
		// Decode to bson map
		var result bson.M
		err := cur.Decode(&result)
		// Convert bson.M to struct
		r := repository.FromBson(&result)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		persistents = append(persistents, r)
	}
	return persistents, nil
}

func FindByUUID(repository Repository, uuid string) (*Repository, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, _ := ConnectDb()
	collection := db.Collection(repository.getCollection())
	filter := bson.M{"uuid": uuid}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		// Decode to bson map
		var result bson.M
		err := cur.Decode(&result)
		// Convert bson.M to struct
		repository.FromBson(&result)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return &repository, nil
}

func UpdateByUUID(repository Repository, uuid string) (*Repository, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, _ := ConnectDb()
	collection := db.Collection(repository.getCollection())
	doc, err := ToDoc(repository)
	update := bson.D{
		{"$set", doc},
	}
	filter := bson.D{{"uuid", uuid}}
	collection.FindOneAndUpdate(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &repository, nil
}

func ToDoc(v interface{}) (doc *bson.D, err error) {
    data, err := bson.Marshal(v)
    if err != nil {
        return
    }
    err = bson.Unmarshal(data, &doc)
    return
}

func ToJson(r interface{}) string {
	b, err := json.Marshal(r)
	if err != nil {
		log.Printf("Error: %s", err)
		return ""
	}
	return string(b)
}

func DeleteByUUID(repository Repository, uuid string) (int64, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, _ := ConnectDb()
	collection := db.Collection(repository.getCollection())
	filter := bson.D{{"uuid", uuid}}
	deleteResult, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return deleteResult.DeletedCount, err
}
