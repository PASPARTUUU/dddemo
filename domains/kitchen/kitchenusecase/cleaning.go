package kitchenusecase

import (
	"dddemo/domains/kitchen"
	"dddemo/models"
)

type CleaningUseCase struct {
	ingredientRepo kitchen.IngredientRepo
}

func NewCleaningUseCase(ingredientRepo kitchen.IngredientRepo) *CleaningUseCase {
	return &CleaningUseCase{
		ingredientRepo: ingredientRepo,
	}
}

func (uc CleaningUseCase) CleanIt() ([]models.Trash, error) {
	ingrs, _ := uc.ingredientRepo.RShowIngredients()

	ingrNames := make([]string, 0, len(ingrs))
	for _, el := range ingrs {
		ingrNames = append(ingrNames, el.Name)
	}

	trashIngrs, _ := uc.ingredientRepo.RTakeIngredients(ingrNames...)

	trash := make([]models.Trash, 0, len(trashIngrs))
	for _, el := range trashIngrs {
		trash = append(trash, models.Trash{Name: el.Name})
	}

	return trash, nil
}
