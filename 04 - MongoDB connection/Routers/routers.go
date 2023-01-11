package routers

import (
	"github.com/DevopsGuyXD/mongotest/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/api/movies",controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movies",controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}",controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/api/movies/{id}",controllers.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/movies",controllers.DeleteAllMovies).Methods("DELETE")

	return router
}