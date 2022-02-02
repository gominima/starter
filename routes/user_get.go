package routes

import (
        "github.com/fatih/color"
	"github.com/gominima/minima"
)

func UserGet() *minima.Router {
	rt := minima.NewRouter()
        rt.Post("/:id", func(res *minima.Response, req *minima.Request) {
		//This is a super simple create method without any logic implemented!
		res.Status(200)
		param := req.GetParam("id")
		color.Cyan("User id is %v", param)
		res.Send("This is a Get method")
	})
	return rt
}