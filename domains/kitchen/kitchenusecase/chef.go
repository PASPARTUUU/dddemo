package kitchenusecase

import (
	"fmt"
	"strings"

	"dddemo/domains/kitchen"
	"dddemo/models"
)

// шеф-повар
type ChefUseCase struct {
	ingredientRepo kitchen.IngredientRepo
}

func NewChefUseCase(ingredientRepo kitchen.IngredientRepo) *ChefUseCase {
	return &ChefUseCase{
		ingredientRepo: ingredientRepo,
	}
}

func (uc ChefUseCase) Cook(dish models.Dish) (string, error) {

	ingr, err := uc.ingredientRepo.RTakeIngredients(strings.Split(dish.Recipe, ";")...)

	res := fmt.Sprintf("making dish with [%v]", ingr)

	return res, err
}
