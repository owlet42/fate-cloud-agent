package db

import(
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"log"
	"time"
)

type Repository interface {
	getCollection() string
	FromBson(m *bson.M)
	GetUuid() string
}

func Save(pepository Repository) (string, error){
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, _ := ConnectDb()
	collection := db.Collection(pepository.getCollection())
	_, err := collection.InsertOne(ctx, pepository)
	if err != nil {
		log.Println(err)
		return "",err
	}
	return pepository.GetUuid(),nil
}

func Find(pepository Repository) ([]*Repository, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, _ := ConnectDb()
	collection := db.Collection(pepository.getCollection())
	cur, err := collection.Find(ctx, bson.M{}) // find
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cur.Close(ctx)
	persistents := []*Repository{}
	for cur.Next(ctx) {
		// Decode to bson map
		var result bson.M
		err := cur.Decode(&result)
		// Convert bson.M to struct
		pepository.FromBson(&result)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		persistents = append(persistents, &pepository)
	}
	return persistents, nil
}

func FindByUUID(pepository Repository, uuid string) (*Repository, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, _ := ConnectDb()
	collection := db.Collection(pepository.getCollection())
	filter := bson.M{"baseobject.uuid": uuid}
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
		pepository.FromBson(&result)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return &pepository, nil
}

func UpdateByUUID(pepository Repository, uuid string) (*Repository, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, _ := ConnectDb()
	collection := db.Collection(pepository.getCollection())
	doc, err := ToDoc(pepository)
	update := bson.D{
		{"$set", doc},
	}
	filter := bson.D{{"baseobject.uuid", uuid}}
	collection.FindOneAndUpdate(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &pepository, nil
}

func ToDoc(v interface{}) (doc *bson.D, err error) {
    data, err := bson.Marshal(v)
    if err != nil {
        return
    }
    err = bson.Unmarshal(data, &doc)
    return
}

func DeleteByUUID(pepository Repository, uuid string) (int64, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, _ := ConnectDb()
	collection := db.Collection(pepository.getCollection())
	filter := bson.D{{"baseobject.uuid", uuid}}
	deleteResult, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	// fmt.Printf("Deleted %v documents in the %v collection\n", deleteResult.DeletedCount, pepository.getCollection())
	return deleteResult.DeletedCount, err
}
