package route

import (
	"github.com/gorilla/mux"

	controller "github.com/DevopsGuyXD/testserver/Controllers"
	util "github.com/DevopsGuyXD/testserver/Utils"
)

func RouteCollection() *mux.Router{

	router := mux.NewRouter()

	router.HandleFunc("/", controller.HealthController).Methods("HEAD")
	router.HandleFunc("/jwt", util.GetJWT).Methods("GET")

	return router
}