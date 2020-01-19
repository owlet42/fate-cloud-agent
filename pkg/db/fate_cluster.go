package db
import (
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"fmt"
	"log"
	"time"
)

type FateCluster struct {
	Uuid       string `json:"uuid"` 
	Name       string `json:"name"`
	NameSpaces string `json:"namespaces"`
	Version    string `json:"version"`
	PartyId    string `json:"party_id"`
	Chart      Helm   `json:"chart"`
}

type Helm struct {
	Name     string `json:"name"` 
	Value    string `json:"value"` 
	Template string `json:"template"` 
}

func NewFateCluster() *FateCluster {
	return new(FateCluster)
}

func (fate *FateCluster) GetCollection() string {
	return "fate"
}


func  SaveFateCluster(fateCluster *FateCluster) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := DB.db.Collection("fate")      // collection

	_, err := collection.InsertOne(ctx, fateCluster)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return fateCluster.Uuid, nil
}


func FindFateCluster() ([]*FateCluster, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := DB.db.Collection("fate")
	cur, err := collection.Find(ctx, bson.M{}) // find
	if err != nil {
		log.Println(err)
		return nil,err
	}
	defer cur.Close(ctx)
	fcs := []*FateCluster{}
	for cur.Next(ctx) {
		s := new(FateCluster)
 		var result bson.M
		err := cur.Decode(&result)
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, s)
		if err != nil {
			log.Println(err)
			return nil,err
		}
		fmt.Println(s)
		fcs = append(fcs, s)
	}
	return fcs,nil
}

func FindFateClusterByUUID(uuid string) (*FateCluster, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := DB.db.Collection("fate")
	result := new(FateCluster)
	filter := bson.M{"uuid": uuid}
	var err error
	err = collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
			log.Fatal(err)
			return nil, err
	}
	return result, nil
}

func UpdateFateClusterByUUID(uuid string, fateCluster *FateCluster) (*FateCluster, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := DB.db.Collection("fate")
	doc, err := toDoc(fateCluster)
	update := bson.D{
		{"$set", doc},
	}
	filter := bson.D{{"uuid", uuid}}
	// updateResult := collection.FindOneAndUpdate(ctx, filter, update)
	collection.FindOneAndUpdate(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return fateCluster, nil
}
func DeleteFateCluster(uuid string) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := DB.db.Collection("fate")
	filter := bson.D{{"uuid", uuid}}
	deleteResult, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	return err
}

func toDoc(v interface{}) (doc *bson.D, err error) {
    data, err := bson.Marshal(v)
    if err != nil {
        return
    }
    err = bson.Unmarshal(data, &doc)
    return
}
