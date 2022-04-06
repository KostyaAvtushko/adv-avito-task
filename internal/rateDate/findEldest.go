package rateDate

import (
	"adv-backend-trainee-assignment/pkg/entities"
	"fmt"
	"sort"
)

func FindEldest(advs []entities.Advertisement) []entities.Advertisement {

	sort.SliceStable(advs, func(i, j int) bool {
		return advs[i].CreatedAt.Before(advs[j].CreatedAt)
	})

	fmt.Println(advs)
	return advs
}
