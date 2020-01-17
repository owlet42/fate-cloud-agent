package db

type BaseObject interface{
	getCollection() string
}

func Save(baseObject BaseObject) (string, error){
	collection := getCollection()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := DB.db.Collection(collection)      // collection

	_, err := collection.InsertOne(ctx, fate)
	if err != nil {
		log.Println(err)
		return "",err
	}
	return "",nil
}

func FindByUUID(uuid string) (*BaseObject, error)

func Find() (*[]BaseObject, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := DB.db.Collection("fate")
	cur, err := collection.Find(ctx, bson.M{}) // find
	if err != nil {
		log.Println(err)
		return nil,err
	}
	defer cur.Close(ctx)
	baseObjects := []BaseObject{}
	for cur.Next(ctx) {

		// 可以decode到bson.M  也就是一个map[string]interface{}中
		// 也可以直接decode到一个对象中
		s := BaseObject{}
		var result bson.M
		err := cur.Decode(&result) // decode 到map
		err = cur.Decode(&s)        // decode 到对象
		if err != nil {
			log.Println(err)
			return nil,err
		}

		// do something with result....
		// 可以将map 或对象序列化为json
		//js ,_:=json.Marshal(result)
		//json.Unmarshal(js,s) //反学序列化回来
		// fmt.Println(s)
		baseObjects = append(baseObjects, s)
	}
	fmt.Println(baseObjects)
	return *baseObjects,nil
}

func UpdateByUUID(uuid string, curd BaseObject) (*BaseObject, error)

func Delete(uuid string) error