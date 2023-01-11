package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	models "github.com/DevopsGuyXD/mongotest/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString = "mongodb+srv://Bharath_Dundi:Pa55word123@cluster0.acb6y7s.mongodb.net/?retryWrites=true&w=majority"
var dbName = "netflix"
var colName = "watchlist"

var collection *mongo.Collection

func init(){
	clientOption := options.Client().ApplyURI(`mongodb+srv://Bharath_Dundi:Pa55word123@cluster0.acb6y7s.mongodb.net/?retryWrites=true&w=majority`)

	client, err := mongo.Connect(context.TODO(),clientOption)
	checkForNil(err)

	fmt.Println("MongoDB connected successfully")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection ready to serve")
}

func getAllMovies() []primitive.M{
	filter := bson.D{{}}

	res, err := collection.Find(context.Background(),filter)
	checkForNil(err)

	var getallmovies []primitive.M
	for res.Next(context.Background()){
		var movie bson.M
		res.Decode(&movie)

		getallmovies = append(getallmovies, movie)
	}

	return getallmovies
}

func createMovie(movie models.Netflix){
	res, err := collection.InsertOne(context.Background(),movie)
	checkForNil(err)

	fmt.Println("Movie created successfully", res.InsertedID)
}

func updateMovie(movieId string){
	id, err := primitive.ObjectIDFromHex(movieId)
	checkForNil(err)

	filter := bson.M{"_id":id}
	updateMovie := bson.M{"$set":bson.M{"watched":true}}

	res, err := collection.UpdateOne(context.Background(),filter,updateMovie)
	checkForNil(err)

	fmt.Println("Record updated successfully", res.ModifiedCount)
}

func deleteMovie(movieId string){
	id, err := primitive.ObjectIDFromHex(movieId)
	checkForNil(err)

	filter := bson.M{"_id":id}

	res, err := collection.DeleteOne(context.Background(),filter)
	checkForNil(err)

	fmt.Println("Movie deleted", res.DeletedCount)
}

func deleteAllMovies() int64{
	filter := bson.D{{}}

	res, err := collection.DeleteMany(context.Background(),filter)
	checkForNil(err)

	fmt.Println("All movies have been deleted",res.DeletedCount)
	return res.DeletedCount
}

func GetAllMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")

	getallMovies := getAllMovies()

	json.NewEncoder(w).Encode(getallMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	if r.Body == nil{
		json.NewEncoder(w).Encode("Please send some data")
	}

	var addMovie models.Netflix
	err := json.NewDecoder(r.Body).Decode(&addMovie)
	checkForNil(err)

	createMovie(addMovie)

	json.NewEncoder(w).Encode(addMovie)
}

func UpdateMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	params := mux.Vars(r)
	deleteMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	res := deleteAllMovies()

	json.NewEncoder(w).Encode(res)
}

func checkForNil(err error){
	if err != nil{
		log.Fatal(err)
	}
}