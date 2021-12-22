package shop

import (
	"dddemo/domains/shop/aggregates"
)

type DishUseCase interface {
	UCreateDish(name, recipe string, price int) error
	UGetDishes() ([]aggregates.Dish, error)
	UDeleteDish(dishName string) error
}
