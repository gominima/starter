package main

import (
	"github.com/gominima/cors"
	"github.com/gominima/minima"
	"github.com/gominima/middlewares"
	"github.com/gominima/starter/routes"
)

func main() {
	app := minima.New()
	crs := cors.New()
	app.Use(crs.NewCors(cors.Options{
		AllowedOrigins:   []string{"https://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
	}))
	app.UseRaw(middleware.Logger)
	app.Get("/", func(res *minima.Response, req *minima.Request) {
		res.OK().Send("Hello World")
	})
	app.UseRouter(routes.Router())
	app.Listen(":3000")
}
