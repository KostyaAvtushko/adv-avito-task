package rateDate

import (
	"adv-backend-trainee-assignment/pkg/entities"
	"fmt"
	"sort"
)

func FindNewest(advs []entities.Advertisement) []entities.Advertisement {

	sort.SliceStable(advs, func(i, j int) bool {
		return advs[i].CreatedAt.After(advs[j].CreatedAt)
	})

	fmt.Println(advs)
	return advs
}
