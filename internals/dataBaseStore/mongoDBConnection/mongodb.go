package mongoDBConnection

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//var uri string = "mongodb://127.0.0.1:54327"

var uri string = "mongodb://127.0.0.1:27017/54327?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.6.0"

var Client *mongo.Client = MongoDBConnection()

func MongoDBConnection() *mongo.Client {
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	//defer func() {
	//	if err := client.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()
	//Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	listDB, _ := client.ListDatabases(ctx, bson.M{})
	fmt.Println(listDB)

	return client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("apartmentRentingApi").Collection(collectionName)
	return collection
}
