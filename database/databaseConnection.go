package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ConnectionString = "mongodb://localhost:27017"

func DBinstance() *mongo.Client {
	clientOption := options.Client().ApplyURI(ConnectionString)
	client, err := mongo.Connect(context.Background(), clientOption) // establish a connection to the MongoDB server

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDb Connection Sucess")

	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("resturant").Collection(collectionName)

	return collection

}
