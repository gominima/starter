package main

import (
	"github.com/gominima/minima"
	"github.com/gominima/starter/routes"
	"github.com/gominima/middlewares"
	"github.com/gominima/cors"
)

func main() {
	app := minima.New()
	app.UseRouter(routes.Router())
	app.UseRaw(middleware.Logger)

	crs := cors.New()
	app.Use(crs.NewCors(cors.Options{
		AllowedOrigins:   []string{"https://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
	}))

	app.Listen(":3000")
}
