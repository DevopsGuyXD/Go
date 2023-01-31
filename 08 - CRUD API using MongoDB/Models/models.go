package models

import (
	"context"
	"fmt"

	config "github.com/DevopsguyXD/test/Config"
	utils "github.com/DevopsguyXD/test/Utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	Id      primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string	`json:"movie,omitempty"`
	Watched bool `json:"watched,omitempty"`
}

// -------------------- DB Helpers
// --- Get all movies
func MgetAllMovies() []primitive.M{

	filter := bson.D{{}}

	res, err := config.Collection.Find(context.Background(), filter)
	utils.CheckForNil(err)

	var allMovies []primitive.M
	for res.Next(context.Background()){
		var movies bson.M
		res.Decode(&movies)

		allMovies = append(allMovies,movies)
	}

	return allMovies
}

// --- Create movie
func McreateMovie(movie Netflix){
	res, err := config.Collection.InsertOne(context.Background(),movie)
	utils.CheckForNil(err)

	fmt.Println("Movie created successfully")
	fmt.Printf("Created at: %q\n",utils.FormatedTime())
	fmt.Printf("%v\n",res.InsertedID)
}

// --- Update movie
func MupdateMovie(movieId string) int64{
	id, err := primitive.ObjectIDFromHex(movieId)
	utils.CheckForNil(err)

	filter := bson.M{"_id":id}
	update := bson.M{"$set":bson.M{"watched":true}}

	res, err := config.Collection.UpdateOne(context.Background(), filter, update)
	utils.CheckForNil(err)

	if res.ModifiedCount == 1{
		fmt.Println("Movie updated successfully")
		fmt.Printf("Updated at: %q\n",utils.FormatedTime())
		fmt.Printf("Number of records modified: %v\n",res.ModifiedCount)
	}else{
		fmt.Println("Movie already marked as watch")
	}

	return res.ModifiedCount
}

// --- Delete movie
func MdeleteMovie(movieId string) int64{
	id, err := primitive.ObjectIDFromHex(movieId)
	utils.CheckForNil(err)

	filter := bson.M{"_id":id}

	res, err := config.Collection.DeleteOne(context.Background(), filter)
	utils.CheckForNil(err)

	if res.DeletedCount == 1{
		fmt.Println("Movie deleted successfully")
		fmt.Printf("Deleted at: %q\n",utils.FormatedTime())
		fmt.Printf("Number of records deleted: %v\n",res.DeletedCount)
	}else{
		fmt.Println("Movie doesn't exist")
	}

	return res.DeletedCount
}

// --- Delete all movies
func MdeleteAllMovies(){

	filter := bson.D{{}}

	res, err := config.Collection.DeleteMany(context.Background(),filter)
	utils.CheckForNil(err)

	fmt.Println("All movies delete successfully",res.DeletedCount )	
}
// ---------- End of DB Helpers