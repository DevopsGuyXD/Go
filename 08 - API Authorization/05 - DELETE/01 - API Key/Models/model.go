package model

import (
	"context"
	"fmt"

	config "github.com/DevopsGuyXD/testserver/Configs"
	util "github.com/DevopsGuyXD/testserver/Utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	Id    primitive.ObjectID  `json:"Id,omitempty" bson:"Id,omitempty"`
	Movie  string `json:"Movie,omitempty"`
	Part   string `json:"Part,omitempty"`
	Status string `json:"Status,omitempty"`
}

func AddMovie(movie Netflix) string{

	res, err := config.MongoConnection().InsertOne(context.Background(), movie); util.CheckForNil(err)
	fmt.Println(res.InsertedID)
	
	return "Added successfully"
}

func UpdateMovie(movieid string, movie Netflix) string{

	id, err := primitive.ObjectIDFromHex(movieid); util.CheckForNil(err)

	filter := bson.M{"_id":id}
	update := bson.M{"$set": movie}

	config.MongoConnection().UpdateOne(context.Background(), filter, update)

	return "Updated successfully(PUT)"
}

func PatchMovie(movieid string, movie Netflix) string{

	id, err := primitive.ObjectIDFromHex(movieid); util.CheckForNil(err)

	filter := bson.M{"_id":id}
	update := bson.M{"$set":movie}

	config.MongoConnection().UpdateOne(context.Background(), filter, update)

	return "Updated successfully(PATCH)"
}

func DeleteMovie(movieid string) string{

	id, err := primitive.ObjectIDFromHex(movieid); util.CheckForNil(err)

	filter := bson.M{"_id":id}

	config.MongoConnection().DeleteOne(context.Background(), filter)

	return "Deleted successfully"
}