package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func LoadPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! docker compose test")
	})

	port := LoadPort()
	log.Printf("Starting server on port: %s", port)
	app.Listen(":" + port)
}
