package routes

import "github.com/gominima/minima"

func Router() *minima.Router {
	rt := minima.NewRouter()
	//assign methods to routes 
	rt.Get("/user/:id", UserGet)
	rt.Delete("/user/:id", UserDelete)
	rt.Post("/user/:id", UserCreate)
	rt.Patch("/user/:id", UserUpdate)
	return rt
}
