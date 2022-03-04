package main

import (
	"github.com/gominima/minima"
	"github.com/gominima/starter/routes"
	"github.com/gominima/middlewares"
	"github.com/gominima/cors"
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
	app.UseRouter(routes.Router())
	app.Listen(":3000")
}
