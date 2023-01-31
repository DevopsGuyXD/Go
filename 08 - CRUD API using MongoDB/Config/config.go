package config

import (
	"context"
	"fmt"
	"os"

	utils "github.com/DevopsguyXD/test/Utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func MongoConnection(){

	utils.InitEnvFile()

	clientOptions := options.Client().ApplyURI(os.Getenv("CONNECTION_STRING"))

	client, err := mongo.Connect(context.TODO(), clientOptions)
	utils.CheckForNil(err)

	fmt.Println("MongoDB connecteed successfully...")
	Collection = client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COL_NAME"))
	fmt.Println("Collection ready to serve...")
}