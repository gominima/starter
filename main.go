package main

import "github.com/gominima/minima"

func main() {
 app := minima.New()
 
 app.Listen(":3000")
}