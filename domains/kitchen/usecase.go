package kitchen

import (
	"dddemo/domains/kitchen/aggregates"
	"dddemo/models"
)

type IngredientUseCase interface {
	UAddIngredient(ingr models.Ingredient) error
	UTakeIngredients(names ...string) ([]models.Ingredient, error)
	UShowIngredients() ([]aggregates.Ingredient, error)
}

type ChefUseCase interface {
	Cook(dish models.Dish) (string, error)
}
