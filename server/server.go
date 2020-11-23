package main

import (
	"log"
	"os"

	//"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	//app.Use(cors.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello Heroku")
	})

	app.Static("/", "./build")

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"

	}
	log.Printf("Listening on port %s\n\n", port)
	app.Listen(port)
}
