package main

import (
	"github.com/galaco/squiddleb.it/src/routes"
	"github.com/galaco/yamigo"
	"os"
)

func registerRoutes(router *yamigo.Router) {
	router.Register("GET", "/", routes.Index())
}

func main() {
	app := yamigo.New(os.Getenv("YAMIGO_ENVIRONMENT"))
	registerRoutes(app.Router())

	app.Run()
}
