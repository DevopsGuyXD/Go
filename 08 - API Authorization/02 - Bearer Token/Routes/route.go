package route

import (
	"github.com/gorilla/mux"

	controller "github.com/DevopsGuyXD/testserver/Controllers"
	util "github.com/DevopsGuyXD/testserver/Utils"
)

func RouteCollection() *mux.Router{

	router := mux.NewRouter()

	router.Handle("/", util.ValidateJWT(controller.HealthController)).Methods("GET")
	router.HandleFunc("/bearer", util.GetBearerToken).Methods("GET")

	return router
}