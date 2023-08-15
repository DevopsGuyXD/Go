package route

import (
	"github.com/gorilla/mux"

	controller "github.com/DevopsGuyXD/testserver/Controllers"
	util "github.com/DevopsGuyXD/testserver/Utils"
)

func RouteCollection() *mux.Router{

	router := mux.NewRouter()

	router.HandleFunc("/", controller.HealthController).Methods("GET")
	router.HandleFunc("/jwt", util.GetJWT).Methods("GET")

	router.HandleFunc("/testpost", controller.TestController).Methods("POST")
	router.Handle("/testupdate/{_id}", util.ValidateJWT(controller.UpdateController)).Methods("PUT")
	router.HandleFunc("/testpatch/{_id}", controller.PatchController).Methods("PATCH")
	router.HandleFunc("/testdelete/{_id}", controller.DeleteController).Methods("DELETE")

	return router
}