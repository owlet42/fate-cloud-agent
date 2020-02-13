package db

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

// Repository is the basic interface of the database CRUD
type Repository interface {
	getCollection() string
	FromBson(m *bson.M) interface{}
	GetUuid() string
}

// Save the object in the database
func Save(repository Repository) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, _ := ConnectDb()
	collection := db.Collection(repository.getCollection())
	_, err := collection.InsertOne(ctx, repository)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return repository.GetUuid(), nil
}

// Find find the objects from the database
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

// FindByUUID find the object from the database via uuid
func FindByUUID(repository Repository, uuid string) (interface{}, error) {
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
	var r interface{}
	for cur.Next(ctx) {
		// Decode to bson map
		var result bson.M
		err := cur.Decode(&result)
		// Convert bson.M to struct
		r = repository.FromBson(&result)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return r, nil
}

// UpdateByUUID Update the object in the database via uuid
func UpdateByUUID(repository Repository, uuid string) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, _ := ConnectDb()
	collection := db.Collection(repository.getCollection())
	doc, err := ToDoc(repository)
	if err != nil {
		log.Fatal(err)
		return err
	}

	update := bson.D{
		{"$set", doc},
	}
	filter := bson.M{"uuid": uuid}
	collection.FindOneAndUpdate(ctx, filter, update)

	return nil
}

// ToDoc convert object to bson document
func ToDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &doc)
	return
}

// ToJson convert object to json string
func ToJson(r interface{}) string {
	b, err := json.Marshal(r)
	if err != nil {
		log.Printf("Error: %s", err)
		return ""
	}
	return string(b)
}

// DeleteByUUID delete object from database via uuid
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

// FindByFilter find objects from database via custom filter, such as: findByName, findByStatus
func FindByFilter(repository Repository, filter bson.M) ([]interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, _ := ConnectDb()
	collection := db.Collection(repository.getCollection())
	cur, err := collection.Find(ctx, filter)
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
