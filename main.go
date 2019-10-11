package main

import (
	"github.com/theblueskies/aws-ses-go-emailer/handler"
)

func main() {
	r := handler.GetRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
