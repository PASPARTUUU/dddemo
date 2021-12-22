package aggregates

import "dddemo/models"

type Dish struct {
	models.Dish
	Price int
}
