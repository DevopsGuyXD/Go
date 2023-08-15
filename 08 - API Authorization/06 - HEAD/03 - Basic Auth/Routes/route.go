package route

import (
	"github.com/gorilla/mux"

	controller "github.com/DevopsGuyXD/testserver/Controllers"
	util "github.com/DevopsGuyXD/testserver/Utils"
)

func RouteCollection() *mux.Router{

	router := mux.NewRouter()

	router.HandleFunc("/", util.BasicAuth(controller.HealthController)).Methods("HEAD")

	return router
}