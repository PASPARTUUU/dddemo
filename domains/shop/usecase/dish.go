package usecase

import (
	"dddemo/domains/shop"
	"dddemo/domains/shop/aggregates"
	"dddemo/models"
)

type DishUseCase struct {
	dishRepo shop.DishRepo
}

func NewDishUseCase(dishRepo shop.DishRepo) *DishUseCase {
	return &DishUseCase{
		dishRepo: dishRepo,
	}
}

func (d DishUseCase) UCreateDish(name, recipe string, price int) error {
	dish := aggregates.Dish{
		Dish: models.Dish{
			Name:   name,
			Recipe: recipe,
		},
		Price: price,
	}

	return d.dishRepo.RCreateDish(dish)
}

func (d DishUseCase) UGetDishes() ([]aggregates.Dish, error) {
	return d.dishRepo.RGetDishes()
}

func (d DishUseCase) UDeleteDish(dishName string) error {
	return d.dishRepo.RDeleteDish(dishName)
}
