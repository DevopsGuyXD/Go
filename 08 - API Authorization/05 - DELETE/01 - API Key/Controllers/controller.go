package controller

import (
	"encoding/json"
	"net/http"

	model "github.com/DevopsGuyXD/testserver/Models"
	"github.com/gorilla/mux"
)

var SECRET = []byte("super-secret-auth-key")
var api_key = "1234"

func HealthController(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("Healthy")
}


func TestController(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)

	res := model.AddMovie(movie)

	json.NewEncoder(w).Encode(res)
}


func UpdateController(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)

	movieid := params["_id"]

	res := model.UpdateMovie(movieid, movie)

	json.NewEncoder(w).Encode(res)
}


func PatchController(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)

	movieid := params["_id"]

	res := model.PatchMovie(movieid, movie)

	json.NewEncoder(w).Encode(res)
}


func DeleteController(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	movieid := params["_id"]

	res := model.DeleteMovie(movieid)

	json.NewEncoder(w).Encode(res)
}