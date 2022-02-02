package main

import (
	"github.com/gominima/minima"
	"github.com/gominima/starter/middlewares"
	"github.com/gominima/starter/routes"
	"os"
)

func main() {
	app := minima.New()
	app.Get("/", func(res *minima.Response, req *minima.Request) {
		res.Status(200).Send("Hello World")
	})
	app.UseRouter(routes.Router())
	app.Use(middlewares.Logger)
	app.Listen(os.Getenv("PORT"))
}
