package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection() *mongo.Collection{
	clientOptions := options.Client().ApplyURI("<CONNECTION_STRING>")

	client, err := mongo.Connect(context.TODO(), clientOptions); if err != nil{
		fmt.Println("Cannot connect")
	}

	Collection := client.Database("TestDB").Collection("TestCol")

	return Collection
}