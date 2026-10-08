package main

import (
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/hellogo", func(c fiber.Ctx) error {
		return c.SendString("Hello, Go!")
	})

	app.Get("/higo", func(c fiber.Ctx) error {
		return c.SendString("Hi, Go!")
	})

	if err := app.Listen(":5588"); err != nil {
		panic(err)
	}
}
