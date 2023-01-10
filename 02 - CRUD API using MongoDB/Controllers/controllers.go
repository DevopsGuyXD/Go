package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	models "github.com/DevopsguyXD/test/Models"
	utils "github.com/DevopsguyXD/test/Utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ---------- Database connection
var collection *mongo.Collection

func MongoConnection() {

	utils.InitEnvFile()

	clientOptions := options.Client().ApplyURI(os.Getenv("CONNECTION_STRING"))

	client, err := mongo.Connect(context.TODO(), clientOptions)
	utils.CheckForNil(err)
	
	fmt.Println("MongoDB connecteed successfully...")
	collection = client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COL_NAME"))
	fmt.Println("Collection ready to serve...")
}
// ---------- End of Database connection












// ---------- DB Helpers
// --- Get all movies
func getAllMovies() []primitive.M{

	filter := bson.D{{}}

	res, err := collection.Find(context.Background(), filter)
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
func createMovie(movie models.Netflix){
	res, err := collection.InsertOne(context.Background(),movie)
	utils.CheckForNil(err)

	fmt.Println("Movie created successfully")
	fmt.Printf("Created at: %q\n",utils.FormatedTime())
	fmt.Printf("%v\n",res.InsertedID)
}

// --- Update movie
func updateMovie(movieId string) int64{
	id, err := primitive.ObjectIDFromHex(movieId)
	utils.CheckForNil(err)

	filter := bson.M{"_id":id}
	update := bson.M{"$set":bson.M{"watched":true}}

	res, err := collection.UpdateOne(context.Background(), filter, update)
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
func deleteMovie(movieId string) int64{
	id, err := primitive.ObjectIDFromHex(movieId)
	utils.CheckForNil(err)

	filter := bson.M{"_id":id}

	res, err := collection.DeleteOne(context.Background(), filter)
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
func deleteAllMovies(){

	filter := bson.D{{}}

	res, err := collection.DeleteMany(context.Background(),filter)
	utils.CheckForNil(err)

	fmt.Println("All movies delete successfully",res.DeletedCount )	
}
// ---------- End of DB Helpers








// ---------- Controllers
// --- Home
func HomeHandler(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	home_message := "Welcome to the Mongo API application"

	json.NewEncoder(w).Encode(home_message)
}

// ---- Get all movies
func GetAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	getallmovies := getAllMovies()
	json.NewEncoder(w).Encode(getallmovies)
}

// --- Create movie
func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	var addMovie models.Netflix
	json.NewDecoder(r.Body).Decode(&addMovie)

	createMovie(addMovie)

	json.NewEncoder(w).Encode(addMovie)
}

// --- Update movie
func UpdateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	updatemovie := updateMovie(params["id"])

	if updatemovie == 1{
		json.NewEncoder(w).Encode(params["id"])
	}else{
		json.NewEncoder(w).Encode("Movie already marked as watch")
	}
}

// --- Delete movie
func DeleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	deletemovie := deleteMovie(params["id"])

	if deletemovie == 1{
		json.NewEncoder(w).Encode(params["id"])
	}else{
		json.NewEncoder(w).Encode("Movie does not exist")
	}
}

// --- Delete all movies
func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	deleteAllMovies()

	json.NewEncoder(w).Encode("All movies deleted successfully")
} 
// ---------- End of Controllers