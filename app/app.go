package main

import (
	"adv-backend-trainee-assignment/pkg/database"
	"adv-backend-trainee-assignment/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	database.Connect()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
