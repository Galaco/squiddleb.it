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
	env := os.Getenv("YAMIGO_ENVIRONMENT")
	if len(env) < 3 {
		env = "development"
	}
	app := yamigo.New(env)
	registerRoutes(app.Router())

	app.Run()
}
