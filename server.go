package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func LoadPort() string {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	return os.Getenv("PORT")
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Printf("Starting server on port: %s", LoadPort())
	app.Listen(":" + LoadPort())
}
