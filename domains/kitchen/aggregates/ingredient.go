package aggregates

import "dddemo/models"

type Ingredient struct {
	models.Ingredient
	Count int // кол на складе
}
