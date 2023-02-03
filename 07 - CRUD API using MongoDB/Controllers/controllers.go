package controllers

import (
	"encoding/json"
	"net/http"

	models "github.com/DevopsguyXD/test/Models"
	"github.com/gorilla/mux"
)

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

	getallmovies := models.MgetAllMovies()
	json.NewEncoder(w).Encode(getallmovies)
}

// --- Create movie
func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	var addMovie models.Netflix
	json.NewDecoder(r.Body).Decode(&addMovie)

	models.McreateMovie(addMovie)

	json.NewEncoder(w).Encode(addMovie)
}

// --- Update movie
func UpdateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	updatemovie := models.MupdateMovie(params["id"])

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

	deletemovie := models.MdeleteMovie(params["id"])

	if deletemovie == 1{
		json.NewEncoder(w).Encode(params["id"])
	}else{
		json.NewEncoder(w).Encode("Movie does not exist")
	}
}

// --- Delete all movies
func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	models.MdeleteAllMovies()

	json.NewEncoder(w).Encode("All movies deleted successfully")
} 
// ---------- End of Controllers