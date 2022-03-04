package routes

import "github.com/gominima/minima"

func Router() *minima.Router {
	router := minima.NewRouter()
    router.Get("/random", RandomGet)
	return router
}