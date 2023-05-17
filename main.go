package main

import (
	// "demoapp/database"
	// "demoapp/lead"
	// "fmt"

	"github.com/gofiber/fiber/v2"
)




func main() {
	app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

	app.Listen(":3000")
}