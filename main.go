package main

import (
	"github.com/rijalfm/go-todo/routers"
)

func main() {

	app := routers.SetupRouter()

	// Running app on port 1212
	app.Run("localhost:1212")
	
}
