package soloAdv

import (
	"adv-backend-trainee-assignment/pkg/database"
	"adv-backend-trainee-assignment/pkg/entities"
	"github.com/gofiber/fiber/v2"
	"log"
)

func GetAdv(ctx *fiber.Ctx) error {
	var adv entities.Advertisement

	var query entities.Query

	if err := ctx.QueryParser(&query); err != nil {
		log.Panic(err)
	}

	id := ctx.Params("id")
	database.DB.Where("id = ?", id).First(&adv)

	fieldCount := len(query.Fields)
	if fieldCount == 2 {
		return ctx.JSON(fiber.Map{
			"name":     adv.Name,
			"desc":     adv.Description,
			"photoURL": adv.PhotoURL,
			"price":    adv.Price,
		})
	}
	if fieldCount == 1 && query.Fields[0] == "desc" {
		return ctx.JSON(fiber.Map{
			"name":     adv.Name,
			"desc":     adv.Description,
			"photoURL": adv.PhotoURL[0],
			"price":    adv.Price,
		})
	}
	if fieldCount == 1 && query.Fields[0] == "photos" {
		return ctx.JSON(fiber.Map{
			"name":     adv.Name,
			"photoURL": adv.PhotoURL,
			"price":    adv.Price,
		})
	}

	return ctx.JSON(fiber.Map{
		"name":     adv.Name,
		"photoURL": adv.PhotoURL[0],
		"price":    adv.Price,
	})
}
