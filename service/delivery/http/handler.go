package http

import (
	"dddemo/domains/kitchen"
	"dddemo/domains/shop"
)

type Handler struct {
	Cleaning *HandlerCleaning
}

func NewHTTPHandlerServices(ingredientRepo kitchen.IngredientRepo, dishRepo shop.DishRepo) *Handler {
	return &Handler{
		Cleaning: NewHandlerCleaning(ingredientRepo, dishRepo),
	}
}
