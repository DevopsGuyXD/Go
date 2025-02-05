package router

import (
	controller "github.com/DevopsGuyXD/testing/Controller"
	middleware "github.com/DevopsGuyXD/testing/Middleware"
	"github.com/go-chi/chi/v5"
)

func RouteCollection() chi.Router {

	router := chi.NewRouter()

	// Route Collection
	router.Route("/api", func(r chi.Router) {
		// r.Use(middleware.AuthMiddleware) // Adding middleware to route collection
		r.Get("/test1", controller.Test1Controller)
		r.Get("/test2", controller.Test2Controller)
	})

	// Query Parameters
	// http://localhost:8000/api/paramtest?firstname=bharath&lastname=dundi
	router.Get("/paramtest", controller.TestQueryParameter)

	// Protected route
	// curl -X GET http://localhost:8000/login
	// curl -X GET http://localhost:8000/login -H "Authorization: Bearer mysecrettoken"
	router.With(middleware.AuthMiddleware).Get("/login", controller.TestProtectedController)

	return router
}
