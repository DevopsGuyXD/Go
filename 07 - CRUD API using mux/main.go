package main

import (
	"Encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Database simulation
type Movie struct{
	Id int `json:"id,omitempty"`
	Movie string `json:"movie,omitempty"`
	Part int `json:"part,omitempty"`
	Director Director `json:"director,omitempty"`
}

type Director struct{
	Firstname string `json:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty"`
}

var movies []Movie


func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	fmt.Fprintln(w, "Welcome to the movies application")
}


func getAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	json.NewEncoder(w).Encode(movies)
}


func getMovieById(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	for _, value := range movies{
		if strconv.Itoa(value.Id) == params["id"]{
			json.NewEncoder(w).Encode(value)
			return
		}
	}
}


func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	var createMovie Movie
	json.NewDecoder(r.Body).Decode(&createMovie)
	rand.Seed(time.Now().UnixNano())
	createMovie.Id = rand.Intn(100000)
	movies = append(movies, createMovie)

	json.NewEncoder(w).Encode(movies)
}


func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	for index, value :=range movies{
		if strconv.Itoa(value.Id) == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
		}

		var updateMovie Movie
		json.NewDecoder(r.Body).Decode(&updateMovie)
		updateMovie.Id = value.Id
		movies = append(movies,updateMovie)

		fmt.Fprintln(w,"Record updated successfully - ID:",params["id"])
		json.NewEncoder(w).Encode(updateMovie)
		return
	}
}


func deleteMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	for index, value := range movies{
		if strconv.Itoa(value.Id) == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)

			fmt.Fprintln(w,"Record deleted successfully - ID:",params["id"])
			break
		}
	}
}


func main() {

	server := mux.NewRouter()

	server.HandleFunc("/",homeHandler).Methods("GET")
	server.HandleFunc("/movies",getAllMovies).Methods("GET")
	server.HandleFunc("/movies/{id}",getMovieById).Methods("GET")
	server.HandleFunc("/movies",createMovie).Methods("POST")
	server.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	server.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Println("Server listening on port 8000")

	err := http.ListenAndServe(":8000",server); if err != nil{
		log.Fatal(err)
	}
}