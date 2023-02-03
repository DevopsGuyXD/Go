package routes

import (
	controllers "github.com/DevopsguyXD/test/Controllers"
	"github.com/gorilla/mux"
)

func RouteCollection() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	router.HandleFunc("/api/movies/all", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movies/create", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/update/{id}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/api/movies/delete/{id}", controllers.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/movies/deleteall",controllers.DeleteAllMovies).Methods("DELETE")

	return router
}