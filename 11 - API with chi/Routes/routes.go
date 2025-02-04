package router

import (
	controller "github.com/DevopsGuyXD/testing/Controller"
	"github.com/go-chi/chi/v5"
)

func RouteCollection() chi.Router {

	router := chi.NewRouter()

	router.Route("/api", func(r chi.Router) {
		r.Get("/test1", controller.Test1Controller)
		r.Get("/test2", controller.Test2Controller)
	})

	return router
}
