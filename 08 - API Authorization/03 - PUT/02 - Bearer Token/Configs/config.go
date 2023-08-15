package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection() *mongo.Collection{
	clientOptions := options.Client().ApplyURI("mongodb+srv://Bharath_Dundi:Pa55word123@cluster0.acb6y7s.mongodb.net/?retryWrites=true&w=majority")

	client, err := mongo.Connect(context.TODO(), clientOptions); if err != nil{
		fmt.Println("Cannot connect")
	}

	Collection := client.Database("TestDB").Collection("TestCol")

	return Collection
}