package main

import (
	"log"
	"os"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

func serveStatic(app *fiber.App) {
	app.Static("/", "./build")
}
func main() {
	//server
	app := fiber.New()
	//Handle Cors
	app.Use(cors.New())
	//Serve the build file
	serveStatic(app)
	//Setup Routes
	//setupRoutes(app)
	//Heroku automatically assigns a port our web server. If it   //fails we instruct it to use port 5000
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"

	}
	log.Printf("Listening on port %s\n\n", port)
	app.Listen(port)
}
