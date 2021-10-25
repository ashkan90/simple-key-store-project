package main

import (
	"os"
	"ys-project/adapters/rest"
)

func main() {
	// Rest adapter which serves the routes within given port
	var herokuPort = os.Getenv("PORT")
	var restAdapter = &rest.Adapter{
		Port: herokuPort,
		Host: "",
	}
	restAdapter.Serve()
}
