package shop

import (
	"dddemo/domains/shop/aggregates"
	"dddemo/models"
)

type DishUseCase interface {
	UCreateDish(name, recipe string, price int) error
	UGetDishes() ([]aggregates.Dish, error)
	UDeleteDish(dishName string) error
}

type Cleaning interface {
	CleanIt() ([]models.Trash, error)
}
