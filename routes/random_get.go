package routes

import (
    "github.com/gominima/minima"
    "math/rand"
	"strconv"
)

func RandomGet(res *minima.Response, req *minima.Request) {
    res.OK().Send(strconv.Itoa(rand.Intn(100)))
}