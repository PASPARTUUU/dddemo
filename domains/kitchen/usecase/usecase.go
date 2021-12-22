package usecase

import "dddemo/domains/kitchen"

// шеф-повар
type UseCase struct {
	ChefUC       *ChefUseCase
	IngredientUC *IngredientUseCase
}

func NewUseCase(ingredientRepo kitchen.IngredientRepo) *UseCase {
	return &UseCase{
		ChefUC:       NewChefUseCase(ingredientRepo),
		IngredientUC: NewIngredientUseCase(ingredientRepo),
	}
}
