package main

import (
	"log"
	"github.com/rijalfm/go-todo/config"
	"github.com/rijalfm/go-todo/routers"
)

func main() {
	// Connect to DB
	config.Connect()

	// Initialize Router
	app := routers.SetupRouter()

	// Running app on port 1212
	log.Fatal(app.Run("localhost:1212"))
	
}
