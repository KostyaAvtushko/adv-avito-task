package entities

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Advertisement struct {
	gorm.Model
	Name        string
	Description string
	PhotoURL    pq.StringArray `gorm:"type:text[]"`
	Price       float64
}
