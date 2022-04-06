package groupAdv

import (
	"adv-backend-trainee-assignment/internal/rateDate"
	"adv-backend-trainee-assignment/internal/ratePrice"
	"adv-backend-trainee-assignment/pkg/database"
	"adv-backend-trainee-assignment/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

func GetGroupAdv(ctx *fiber.Ctx) error {

	var advs []entities.Advertisement

	database.DB.Where("ID <> ?", 1).Find(&advs)

	var urlQuery entities.Query
	ctx.QueryParser(&urlQuery)

	if urlQuery.Type == "CheapestToExpensive" {
		advs = ratePrice.FindСheapestToExpensive(advs)
	} else if urlQuery.Type == "ExpensiveToCheapest" {
		advs = ratePrice.FindExpensiveToCheapest(advs)
	} else if urlQuery.Type == "findNewest" {
		advs = rateDate.FindNewest(advs)
	} else if urlQuery.Type == "findEldest" {
		advs = rateDate.FindEldest(advs)
	}

	var readyGroupAdv [10]entities.AdvertisementGroupItem

	for i := 0; i < 10; i++ {
		adv := entities.AdvertisementGroupItem{
			Name:     advs[i].Name,
			PhotoURL: advs[i].PhotoURL[0],
			Price:    advs[i].Price,
		}
		readyGroupAdv[i] = adv
	}

	return ctx.JSON(readyGroupAdv)
}
