package soloAdv

import (
	"adv-backend-trainee-assignment/pkg/database"
	"adv-backend-trainee-assignment/pkg/entities"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
)

func GetAdv(ctx *fiber.Ctx) error {

	var adv entities.Advertisement

	var query entities.Query

	if err := ctx.QueryParser(&query); err != nil {
		log.Panic(err)
	}

	id := ctx.Params("id")
	err := database.DB.Where("id = ?", id).First(&adv).Error
	errors.Is(err, gorm.ErrRecordNotFound)

	fieldCount := len(query.Fields)
	if fieldCount == 2 && ((query.Fields[0] == "desc" && query.Fields[1] == "photos") ||
		(query.Fields[1] == "desc" && query.Fields[0] == "photos")) {
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
	if fieldCount > 2 {
		return ctx.JSON(fiber.Map{
			"message": "to much fields",
		})
	}

	return ctx.JSON(fiber.Map{
		"name":     adv.Name,
		"photoURL": adv.PhotoURL[0],
		"price":    adv.Price,
	})
}
