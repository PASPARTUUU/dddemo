package kitchen

import (
	"dddemo/domains/kitchen/aggregates"
	"dddemo/models"
)

type IngredientRepo interface {
	RAddIngredient(ingr models.Ingredient) error
	RTakeIngredients(names ...string) ([]models.Ingredient, error)
	RShowIngredients() ([]aggregates.Ingredient, error)
}
