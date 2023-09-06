package main

import (
	"github.com/seetharamugn/car-microservices/Routes"
	"net/http"
)

func main() {
	Routes.SetupRotes()
	http.ListenAndServe(":8080", nil)
}
