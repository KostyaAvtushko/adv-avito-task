package soloAdv

import (
	"adv-backend-trainee-assignment/pkg/database"
	"adv-backend-trainee-assignment/pkg/entities"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func AddAdv(ctx *fiber.Ctx) error {
	var data map[string]any

	if err := ctx.BodyParser(&data); err != nil {
		log.Panic(err)
	}

	if len(data["description"].(string)) > 1000 {
		return ctx.JSON(fiber.Map{
			"message": "description is too long",
		})
	} else if len(data["name"].(string)) > 200 {
		return ctx.JSON(fiber.Map{
			"message": "name is too long",
		})
	} else if len(data["photoURL"].([]any)) > 3 {
		return ctx.JSON(fiber.Map{
			"message": "too much photos",
		})
	}

	photoArr := make([]string, len(data["photoURL"].([]any)))
	for i, v := range data["photoURL"].([]any) {
		photoArr[i] = fmt.Sprint(v)
	}

	adv := entities.Advertisement{
		Name:        data["name"].(string),
		Description: data["description"].(string),
		PhotoURL:    photoArr,
		Price:       data["price"].(float64),
	}

	database.DB.Save(&adv)

	return ctx.JSON(fiber.Map{
		"Id":      adv.ID,
		"message": "success",
	})
}
