package ratePrice

import (
	"adv-backend-trainee-assignment/pkg/entities"
	"fmt"
	"sort"
)

func FindExpensiveToCheapest(advs []entities.Advertisement) []entities.Advertisement {

	sort.SliceStable(advs, func(i, j int) bool {
		return advs[i].Price > advs[j].Price
	})

	fmt.Println(advs)
	return advs
}
