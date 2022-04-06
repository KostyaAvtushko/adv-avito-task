package routes

import (
	"adv-backend-trainee-assignment/pkg/controllers/manyAdvs"
	"adv-backend-trainee-assignment/pkg/controllers/soloAdv"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	get := app.Group("/get")
	get.Get("/groupAdv", manyAdvs.GetGroupAdv)
	get.Get("/adv/:id", soloAdv.GetAdv)

	add := app.Group("/add")
	add.Post("/adv", soloAdv.AddAdv)

}
