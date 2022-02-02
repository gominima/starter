package routes

import "github.com/gominima/minima"

func Router() *minima.Router{
	rt := minima.NewRouter()
	//Mount all routes

	//As these routers are already mounted on /user we just append it to the stack
	rt.UseRouter(UserCreate())
	rt.UseRouter(UserDelete())

	// But in the case of these routers these arent mounted we specifically mount them to the router /user
	rt.Mount("/user", UserGet()) 
	rt.Mount("/user",UserUpdate())
	return rt
}