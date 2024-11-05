package infrastructure

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMonogodb(databaseName string, collectionName string, URI string) *mongo.Collection {
	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(databaseName).Collection(collectionName)
	return collection
}
