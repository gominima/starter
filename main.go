package main;

import (
	"github.com/gominima/minima"
	"github.com/gominima/middlewares"
	"github.com/gominima/starter/routes"
)

func main() {
	app := minima.New()
	app.UseRouter(routes.Router())
	app.UseRaw(middlewares.Logger())
	app.Listen(":3000")
}
