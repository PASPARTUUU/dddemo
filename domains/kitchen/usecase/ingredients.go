package usecase

import (
	"dddemo/domains/kitchen"
	"dddemo/domains/kitchen/aggregates"
	"dddemo/models"
)

type IngredientUseCase struct {
	ingredientRepo kitchen.IngredientRepo
}

func NewIngredientUseCase(ingredientRepo kitchen.IngredientRepo) *IngredientUseCase {
	return &IngredientUseCase{
		ingredientRepo: ingredientRepo,
	}
}

func (uc IngredientUseCase) UAddIngredient(ingr models.Ingredient) error {
	return uc.ingredientRepo.RAddIngredient(ingr)
}
func (uc IngredientUseCase) UTakeIngredients(names ...string) ([]models.Ingredient, error) {
	return uc.ingredientRepo.RTakeIngredients(names...)
}
func (uc IngredientUseCase) UShowIngredients() ([]aggregates.Ingredient, error) {
	return uc.ingredientRepo.RShowIngredients()
}
