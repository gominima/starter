package middlewares

import ("github.com/gominima/minima"
        "github.com/fatih/color")

func Logger(res *minima.Response,  req *minima.Request) {
	color.Blue("Method %s in route :", req.Method())
	color.Magenta(req.GetPathURL())
}