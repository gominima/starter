package main;

import (
	"github.com/gominima/minima"
	"github.com/gominima/middlewares"
	"github.com/gominima/cors"
	"github.com/gominima/starter/routes"
)

func main() {
	app := minima.New()
	cors := cors.New()
	app.UseRouter(routes.Router())
	app.UseRaw(middlewares.Logger())
	app.Use(cors.NewCors(cors.Options{
        	AllowedOrigins: []string{"http://example.com"},
        	AllowCredentials: true,
        	// Enable Debugging for testing, consider disabling in production
        	Debug: true,
	})())
	app.Listen(":3000")
}
