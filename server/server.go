package main

import (
	"log"
	"os"

	//"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

func serveStatic(app *fiber.App) {
	app.Static("/", "./build")
}

func main() {
	app := fiber.New()
	//app.Use(cors.New())

	serveStatic(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"

	}
	log.Printf("Listening on port %s\n\n", port)
	app.Listen(port)
}
