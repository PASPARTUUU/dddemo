package shop

import "dddemo/domains/shop/aggregates"

type DishRepo interface {
	RCreateDish(dish aggregates.Dish) error
	RGetDishes() ([]aggregates.Dish, error)
	RDeleteDish(dishName string) error
}
