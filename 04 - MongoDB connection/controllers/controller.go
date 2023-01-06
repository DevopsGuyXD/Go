// The Database parts of this file are mostly standard and need to be rememberd as muscle memory
package controller

import (
	"context"
	"fmt"
	"log"
	model "mongodb/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Bharath_Dundi :Pa55word123@cluster0.acb6y7s.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init(){
	clientOption := options.Client().ApplyURI(connectionString)
	client, err:= mongo.Connect(context.TODO(), clientOption); if err != nil{
		panic(err)
	}

	fmt.Println("MongoDB connection successful")

	collection = client.Database(dbName).Collection(colName)

	fmt.Print("Collection reference is ready")
}

//----------------- MongoDB Helpers - Should be in a different file --------------------

// Insert 1 record
func insertOneMovie(movie model.Netflix){
	inserted, err := collection.InsertOne(context.Background(), movie); if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Inserted one Movie into DB with id: ", inserted.InsertedID)
}

// Update 1 record

func updateOneMovie(movieId string){
	id, err := primitive.ObjectIDFromHex(movieId); if err != nil{
		log.Fatal(err)
	}
	
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update); if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Modified count: ", result.ModifiedCount)
}

// Delete 1 movie
func deleteOneMovie(movieId string){
	id, err := primitive.ObjectIDFromHex(movieId); if err != nil{
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter); if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with delete count: ", deleteCount)
}

// Get All movies
func getAllMovies() []primitive.M{
	filter := bson.D{{}}

	result,err := collection.Find(context.Background(),filter); if err != nil{
		log.Fatal(err)
	}

	var movies []primitive.M

	for result.Next(context.Background()){
		var movie bson.M
		err := result.Decode(&movie); if err != nil{
			log.Fatal()
		}

		movies = append(movies, movie)
	}

	return movies
}

// Delete all movies
func deleteAllMovies(){
	filter := bson.D{{}}

	result, err := collection.DeleteMany(context.Background(),filter); if err != nil{
		log.Fatal(err)
	}

	fmt.Println("All records deleted successfully", result.DeletedCount)
}
//---------------------------------- END --------------------------------------------


//------------------------------ Controllers ----------------------------------

// Get all movies controller
func GetAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")
	getallMovies := getAllMovies()

	json.NewEncoder(w).Encode(getallMovies)
}

// Get movie by ID controller
func CreateMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var createMovie model.Netflix

	err := json.NewDecoder(r.Body).Decode(&createMovie); if err != nil{
		log.Fatal(err)
	}

	insertOneMovie(createMovie)

	json.NewEncoder(w).Encode(createMovie)
}

func MarkAsWatched()