package main

import (
	// "demoapp/database"
	// "demoapp/lead"
	// "fmt"

	"github.com/gofiber/fiber/v2"
)




func main() {
	app := fiber.New()
	
    app.Static("/", "./public")

	app.Listen(":3000")
}